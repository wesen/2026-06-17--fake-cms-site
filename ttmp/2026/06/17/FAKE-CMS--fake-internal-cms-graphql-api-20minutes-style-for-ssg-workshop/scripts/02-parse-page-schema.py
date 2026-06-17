#!/usr/bin/env python3
"""
02-parse-page-schema.py — Fetch representative pages and extract a COMPACT,
structured schema for each, mirroring what a CMS Content API would expose.

For each sample URL we extract (no raw HTML dumped to stdout/sources):
  - url, http_status, content_type_guess, slug
  - <title>, meta description, canonical, og:* tags
  - JSON-LD blocks (Yoast/Schema.org: Article, BreadcrumbList, Organization, ...)
  - WordPress body class -> post type + template hints
  - headline (h1), published/modified dates (from meta or JSON-LD)
  - author slug/name (from /author/<slug> links + JSON-LD author)
  - taxonomy: category slugs + tag ("rubrique") slugs
  - featured image url + alt
  - image inventory (count + a few srcs)
  - body block structure: counts of h2/h3/p/figure/blockquote/iframe/ul/li
    (this models a "blocks"-based content field without dumping text)

Output:
  sources/_raw/page-schemas.json         (one record per URL)
  sources/02-page-schemas.md             (compact human-readable table + notes)

Why parse this way:
  A legacy CMS GraphQL API almost always surfaces exactly these fields
  (Yoast SEO schema, WordPress post meta, term taxonomies). Capturing the
  real shape here drives the GraphQL type design in the design doc.
"""
from __future__ import annotations

import json
import re
import sys
import time
import urllib.request
from collections import Counter
from pathlib import Path

from bs4 import BeautifulSoup

# Representative samples across content types / post types / templates.
SAMPLES = [
    # pages (static CMS pages)
    ("page", "https://www.20minutes-media.com/nos-solutions/display"),
    ("page", "https://www.20minutes-media.com/nos-solutions/video"),
    ("page", "https://www.20minutes-media.com/presse"),
    ("page", "https://www.20minutes-media.com/a-propos"),
    ("page", "https://www.20minutes-media.com/marketing/panel"),
    # posts: best-cases (custom post type, the bulk of the site)
    ("post/best-cases", "https://www.20minutes-media.com/best-cases/suez-environnement-2"),
    ("post/best-cases", "https://www.20minutes-media.com/best-cases/mini-adventure-trip-2"),
    ("post/best-cases", "https://www.20minutes-media.com/best-cases/ratp-navigo-2"),
    # posts: actualites (news/announcement posts)
    ("post/actualites", "https://www.20minutes-media.com/actualites/la-21e-minute-prendre-le-temps"),
    ("post/actualites", "https://www.20minutes-media.com/actualites/reuters-institute-20-minutes-2e-marque-la-plus-utilisee-pour-sinformer-en-ligne"),
    # posts: etudes (studies)
    ("post/etudes", "https://www.20minutes-media.com/etudes/dela-de-laudience-lattachement-francais-aux-marques-dinfo"),
    # taxonomy archives
    ("category", "https://www.20minutes-media.com/archives/actualites"),
    ("category", "https://www.20minutes-media.com/archives/best-cases"),
    ("post_tag", "https://www.20minutes-media.com/rubrique/audience"),
    ("post_tag", "https://www.20minutes-media.com/rubrique/moijeune"),
    # author archive
    ("author", "https://www.20minutes-media.com/author/eking20minutes-fr"),
    # home (acts as a curated landing / collection)
    ("home", "https://www.20minutes-media.com/"),
]

HEADERS = {"User-Agent": "fake-cms-research/1.0"}


def fetch(url: str) -> tuple[int, str]:
    req = urllib.request.Request(url, headers=HEADERS)
    try:
        with urllib.request.urlopen(req, timeout=30) as resp:
            return resp.status, resp.read().decode("utf-8", errors="replace")
    except urllib.error.HTTPError as e:
        return e.code, e.read().decode("utf-8", errors="replace") if e.fp else ""


def slug_from(url: str) -> str:
    path = url.rstrip("/").split("/", 3)[-1] if "://" in url else url
    return path.rstrip("/").split("/")[-1] or "<root>"


def meta(soup: BeautifulSoup, name: str) -> str | None:
    for sel in (
        {"name": name},
        {"property": name},
        {"itemprop": name},
    ):
        tag = soup.find("meta", sel)
        if tag and tag.get("content"):
            return tag["content"].strip()
    return None


def all_meta_prefix(soup: BeautifulSoup, prefix: str) -> dict[str, str]:
    out: dict[str, str] = {}
    for m in soup.find_all("meta"):
        for k in ("property", "name"):
            v = m.get(k)
            if v and v.startswith(prefix) and m.get("content"):
                out[v] = m["content"].strip()
    return out


def json_ld(soup: BeautifulSoup) -> list[dict]:
    blocks: list[dict] = []
    for s in soup.find_all("script", type="application/ld+json"):
        raw = s.string or s.get_text() or ""
        # JSON-LD can be a list or an object, and may contain @graph.
        try:
            data = json.loads(raw)
        except json.JSONDecodeError:
            continue
        items = data if isinstance(data, list) else [data]
        for it in items:
            blocks.append(it)
            graph = it.get("@graph") if isinstance(it, dict) else None
            if isinstance(graph, list):
                blocks.extend(graph)
    return blocks


def body_class(soup: BeautifulSoup) -> str:
    body = soup.find("body")
    if body and body.get("class"):
        return " ".join(body["class"])
    return ""


def block_counts(soup: BeautifulSoup, main: BeautifulSoup | None) -> dict[str, int]:
    root = main or soup
    tags = ["h1", "h2", "h3", "h4", "p", "figure", "img", "blockquote",
            "iframe", "ul", "ol", "li", "table", "a", "div"]
    return {t: len(root.find_all(t)) for t in tags}


def find_main(soup: BeautifulSoup) -> BeautifulSoup | None:
    for sel in ("main", "article", "[role=main]"):
        node = soup.select_one(sel)
        if node:
            return node
    return None


def featured_image(soup: BeautifulSoup) -> dict | None:
    # WordPress often marks featured image via og:image and picture/img in article header.
    og = meta(soup, "og:image")
    return {"url": og} if og else None


def terms(soup: BeautifulSoup) -> dict[str, list[str]]:
    """Extract category + tag(term) slugs from archive links."""
    cats: set[str] = set()
    tags: set[str] = set()
    for a in soup.find_all("a", href=True):
        href = a["href"]
        m_cat = re.search(r"/archives/(?:actualites/)?([^/?#]+)/?$", href)
        m_tag = re.search(r"/rubrique/([^/?#]+)/?$", href)
        m_author = re.search(r"/author/([^/?#]+)/?$", href)
        if m_cat:
            cats.add(m_cat.group(1))
        if m_tag:
            tags.add(m_tag.group(1))
    return {
        "categories": sorted(cats),
        "tags": sorted(tags),
    }


def author_slug(soup: BeautifulSoup) -> str | None:
    for a in soup.find_all("a", href=True):
        m = re.search(r"/author/([^/?#]+)/?$", a["href"])
        if m:
            return m.group(1)
    return None


def parse_page(kind: str, url: str) -> dict:
    status, html = fetch(url)
    rec: dict = {"kind": kind, "url": url, "slug": slug_from(url),
                 "http_status": status}
    if not html:
        rec["error"] = "empty body"
        return rec
    soup = BeautifulSoup(html, "lxml")

    rec["title"] = (soup.title.get_text(strip=True) if soup.title else None)
    rec["meta_description"] = meta(soup, "description")
    rec["canonical"] = (soup.find("link", rel="canonical") or {}).get("href") \
        if soup.find("link", rel="canonical") else None
    rec["og"] = all_meta_prefix(soup, "og:")
    rec["twitter"] = all_meta_prefix(soup, "twitter:")
    rec["body_class"] = body_class(soup)
    rec["h1"] = [h.get_text(strip=True) for h in soup.find_all("h1")][:3]
    rec["author_slug"] = author_slug(soup)
    rec["terms"] = terms(soup)
    rec["featured_image"] = featured_image(soup)

    ld = json_ld(soup)
    rec["json_ld_types"] = [b.get("@type") for b in ld if isinstance(b, dict)]
    # Pull the Article/WebPage graph node if present.
    article_node = None
    for b in ld:
        if not isinstance(b, dict):
            continue
        t = b.get("@type")
        if t in ("Article", "NewsArticle", "BlogPosting", "WebPage") or \
           (isinstance(t, list) and any(x in ("Article", "NewsArticle", "BlogPosting") for x in t)):
            article_node = b
            break
    if article_node:
        rec["article_schema"] = {
            k: article_node.get(k)
            for k in ("headline", "datePublished", "dateModified",
                      "author", "image", "articleSection", "keywords")
        }

    main = find_main(soup)
    rec["main_selector"] = main.name if main else None
    rec["block_counts"] = block_counts(soup, main)

    # image sample (compact)
    imgs = [im.get("src") for im in soup.find_all("img") if im.get("src")]
    rec["image_count"] = len(imgs)
    rec["image_sample"] = imgs[:5]
    return rec


def main() -> int:
    here = Path(__file__).resolve().parent
    ticket = here.parent
    sources = ticket / "sources"
    raw = sources / "_raw"
    raw.mkdir(parents=True, exist_ok=True)

    records: list[dict] = []
    for kind, url in SAMPLES:
        print(f"fetch {kind}: {url}", file=sys.stderr)
        try:
            records.append(parse_page(kind, url))
        except Exception as e:  # noqa: BLE001
            records.append({"kind": kind, "url": url, "error": repr(e)})
        time.sleep(0.6)  # be polite

    (raw / "page-schemas.json").write_text(
        json.dumps(records, indent=2, ensure_ascii=False)
    )

    # Compact human-readable summary.
    lines: list[str] = []
    lines.append("# Page schema extraction (compact)\n")
    lines.append(
        "One record per sampled URL. Full structured data is in "
        "`sources/_raw/page-schemas.json`. This view only shows the fields "
        "that drive the GraphQL schema design.\n"
    )
    lines.append("## Per-page summary\n")
    lines.append("| kind | slug | http | post-type (body class) | h1 | author | cats | tags | datePublished | #img |")
    lines.append("|------|------|------|------------------------|----|--------|------|------|---------------|------|")
    for r in records:
        bc = r.get("body_class", "")
        # extract post-type and template tokens from body class
        pt = re.findall(r"(?:post-type|single|page-template|category|tag|author|home|archive)-([a-zA-Z0-9_-]+)", bc)
        h1 = (r.get("h1") or [""])[0][:40]
        terms = r.get("terms", {})
        cats = ",".join(terms.get("categories", [])[:3])
        tags = ",".join(terms.get("tags", [])[:3])
        dp = (r.get("article_schema") or {}).get("datePublished", "")[:10]
        lines.append(
            f"| {r.get('kind')} | `{r.get('slug')}` | {r.get('http_status')} | "
            f"{' '.join(pt[:4])} | {h1} | {r.get('author_slug') or ''} | "
            f"{cats} | {tags} | {dp} | {r.get('image_count', 0)} |"
        )
    lines.append("")

    lines.append("## JSON-LD @type inventory across samples\n")
    type_counter: Counter[str] = Counter()
    for r in records:
        for t in r.get("json_ld_types", []) or []:
            if isinstance(t, list):
                type_counter.update(t)
            else:
                type_counter[t] += 1
    lines.append("| @type | occurrences |")
    lines.append("|-------|-------------|")
    for t, n in type_counter.most_common():
        lines.append(f"| {t} | {n} |")
    lines.append("")

    lines.append("## Body block-count distribution (main content region)\n")
    lines.append("Median/typical block counts tell us how rich the `content` field is.\n")
    lines.append("| kind | h2 | h3 | p | figure | blockquote | iframe | ul | li | img |")
    lines.append("|------|----|----|---|--------|------------|--------|----|----|-----|")
    for r in records:
        bc = r.get("block_counts", {})
        lines.append(
            f"| {r.get('kind')} | {bc.get('h2',0)} | {bc.get('h3',0)} | {bc.get('p',0)} | "
            f"{bc.get('figure',0)} | {bc.get('blockquote',0)} | {bc.get('iframe',0)} | "
            f"{bc.get('ul',0)} | {bc.get('li',0)} | {bc.get('img',0)} |"
        )
    lines.append("")

    lines.append("## Article schema sample (one record)\n")
    lines.append(
        "```json\n"
        + json.dumps(
            next((r.get("article_schema") for r in records if r.get("article_schema")),
                 None),
            indent=2, ensure_ascii=False,
        )
        + "\n```\n"
    )

    (sources / "02-page-schemas.md").write_text("\n".join(lines))
    print(f"Wrote {sources / '02-page-schemas.md'}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
