#!/usr/bin/env python3
"""
01-fetch-sitemaps.py — Fetch the 20minutes-media.com WordPress/Yoast sitemaps
and produce a compact, structured inventory of all known URLs.

Writes:
  sources/00-sitemap-inventory.md   (human-readable summary, kept small)
  sources/_raw/sitemap-urls.json    (machine-readable list of {type, url})

Design notes:
  - The site is WordPress + Yoast SEO. Its sitemap index exposes five
    content facets that map cleanly onto a typical legacy CMS taxonomy:
      post-sitemap      -> articles (custom post types: actualites, best-cases, ...)
      page-sitemap      -> static pages (solutions, presse, mentions-legales, ...)
      category-sitemap  -> category archive terms
      post_tag-sitemap  -> tag ("rubrique") archive terms
      author-sitemap    -> author archive terms
  - We deliberately do NOT fetch the page bodies here. This script only
    inventories URLs so later scripts can pick representative samples.
"""
from __future__ import annotations

import json
import re
import sys
import urllib.request
from html import unescape
from pathlib import Path

BASE = "https://www.20minutes-media.com"
SITEMAPS = [
    "post-sitemap.xml",
    "page-sitemap.xml",
    "category-sitemap.xml",
    "post_tag-sitemap.xml",
    "author-sitemap.xml",
]
# Friendly names for the content-type each sitemap represents.
TYPE_LABEL = {
    "post-sitemap.xml": "post",      # WordPress "post" + custom post types
    "page-sitemap.xml": "page",      # static pages
    "category-sitemap.xml": "category",
    "post_tag-sitemap.xml": "post_tag",  # rendered as /rubrique/<slug>
    "author-sitemap.xml": "author",
}

LOC_RE = re.compile(r"<loc>([^<]+)</loc>")


def fetch(url: str) -> str:
    req = urllib.request.Request(url, headers={"User-Agent": "fake-cms-research/1.0"})
    with urllib.request.urlopen(req, timeout=30) as resp:
        return resp.read().decode("utf-8", errors="replace")


def main() -> int:
    here = Path(__file__).resolve().parent
    ticket = here.parent
    sources = ticket / "sources"
    raw = sources / "_raw"
    raw.mkdir(parents=True, exist_ok=True)

    all_entries: list[dict] = []
    per_type_counts: dict[str, int] = {}
    per_first_segment: dict[str, dict[str, int]] = {}

    for sm in SITEMAPS:
        xml = fetch(f"{BASE}/{sm}")
        urls = LOC_RE.findall(xml)
        urls = [unescape(u) for u in urls]
        ctype = TYPE_LABEL[sm]
        per_type_counts[ctype] = len(urls)
        seg_counts: dict[str, int] = {}
        for u in urls:
            path = u.removeprefix(BASE).lstrip("/")
            seg = path.split("/")[0] if path else "<root>"
            seg_counts[seg] = seg_counts.get(seg, 0) + 1
            all_entries.append({"type": ctype, "url": u, "path": path})
        per_first_segment[ctype] = dict(
            sorted(seg_counts.items(), key=lambda kv: -kv[1])
        )

    (raw / "sitemap-urls.json").write_text(
        json.dumps(all_entries, indent=2, ensure_ascii=False)
    )

    # Human-readable summary (kept compact — counts + samples, not all 666 URLs)
    lines: list[str] = []
    lines.append("# 20minutes-media.com sitemap inventory\n")
    lines.append(
        "Source: WordPress + Yoast SEO sitemap index at `/sitemap.xml`.\n"
        "This file is a compact summary; the full URL list is in "
        "`sources/_raw/sitemap-urls.json`.\n"
    )
    lines.append("## Counts by content type\n")
    lines.append("| type | count |")
    lines.append("|------|-------|")
    for ctype, n in per_type_counts.items():
        lines.append(f"| {ctype} | {n} |")
    lines.append("")

    lines.append("## Top-level path segments per content type\n")
    lines.append(
        "The first path segment reveals custom post types (e.g. `best-cases`, "
        "`actualites`) and CMS routing conventions.\n"
    )
    for ctype, segs in per_first_segment.items():
        lines.append(f"### {ctype}\n")
        lines.append("| first segment | count |")
        lines.append("|--------------|-------|")
        for seg, n in list(segs.items())[:20]:
            lines.append(f"| `{seg}` | {n} |")
        lines.append("")

    lines.append("## Representative sample URLs (for body scraping)\n")
    lines.append(
        "Pick a handful per content type / segment for deeper HTML parsing "
        "(scripts 02+). These are the URLs the sample parsers target.\n"
    )
    samples: list[dict] = []
    by_seg: dict[tuple[str, str], list[str]] = {}
    for e in all_entries:
        key = (e["type"], e["path"].split("/")[0] if e["path"] else "<root>")
        bucket = by_seg.setdefault(key, [])
        if len(bucket) < 3:
            bucket.append(e["url"])
            samples.append(e)
    lines.append("| type | url |")
    lines.append("|------|-----|")
    for e in samples:
        lines.append(f"| {e['type']} | {e['url']} |")
    lines.append("")

    (sources / "00-sitemap-inventory.md").write_text("\n".join(lines))
    print(f"Wrote {sources / '00-sitemap-inventory.md'}")
    print(f"Wrote {raw / 'sitemap-urls.json'} ({len(all_entries)} URLs)")
    for ctype, n in per_type_counts.items():
        print(f"  {ctype}: {n}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
