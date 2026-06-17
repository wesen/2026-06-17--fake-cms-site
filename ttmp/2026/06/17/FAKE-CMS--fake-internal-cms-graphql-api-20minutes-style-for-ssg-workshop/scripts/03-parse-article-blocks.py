#!/usr/bin/env python3
"""
03-parse-article-blocks.py — Resolve JSON-LD @id references and extract a
proper "blocks" content model for one rich article (best-cases post type),
plus one actualites post.

This models the legacy `content`/`blocks` field that a real CMS GraphQL API
exposes: an ordered list of typed blocks (paragraph, heading, image, list,
quote, embed) rather than a single HTML blob.

Outputs:
  sources/_raw/article-blocks.json
  sources/03-article-blocks.md   (compact: schema + block-type histogram
                                  + a truncated block list; full body kept raw-only)

We keep prose out of the summary on purpose (only block *types* and counts,
plus the resolved Schema.org node). Raw extracted blocks live in JSON only.
"""
from __future__ import annotations

import json
import re
import sys
import urllib.request
from collections import Counter
from pathlib import Path

from bs4 import BeautifulSoup, NavigableString, Tag

SAMPLES = [
    ("best-cases", "https://www.20minutes-media.com/best-cases/suez-environnement-2"),
    ("actualites", "https://www.20minutes-media.com/actualites/reuters-institute-20-minutes-2e-marque-la-plus-utilisee-pour-sinformer-en-ligne"),
    ("etudes", "https://www.20minutes-media.com/etudes/dela-de-laudience-lattachement-francais-aux-marques-dinfo"),
]


def fetch(url: str) -> str:
    req = urllib.request.Request(url, headers={"User-Agent": "fake-cms-research/1.0"})
    with urllib.request.urlopen(req, timeout=30) as resp:
        return resp.read().decode("utf-8", errors="replace")


def collect_jsonld(soup: BeautifulSoup) -> tuple[list[dict], dict[str, dict]]:
    """Return (top_level_nodes, by_id_index)."""
    nodes: list[dict] = []
    for s in soup.find_all("script", type="application/ld+json"):
        raw = s.string or s.get_text() or ""
        try:
            data = json.loads(raw)
        except json.JSONDecodeError:
            continue
        items = data if isinstance(data, list) else [data]
        for it in items:
            nodes.append(it)
            if isinstance(it, dict) and isinstance(it.get("@graph"), list):
                nodes.extend(it["@graph"])
    by_id = {n["@id"]: n for n in nodes if isinstance(n, dict) and n.get("@id")}
    return nodes, by_id


def resolve(node, by_id: dict[str, dict], depth=0):
    """Recursively replace {"@id": ...} dicts with their resolved node (bounded)."""
    if depth > 4:
        return node
    if isinstance(node, dict):
        if "@id" in node and len(node) == 1:
            return resolve(by_id.get(node["@id"], node), by_id, depth + 1)
        return {k: resolve(v, by_id, depth + 1) for k, v in node.items()}
    if isinstance(node, list):
        return [resolve(v, by_id, depth + 1) for v in node]
    return node


def find_article_node(nodes: list[dict]) -> dict | None:
    article_types = {"Article", "NewsArticle", "BlogPosting"}
    for n in nodes:
        if not isinstance(n, dict):
            continue
        t = n.get("@type")
        types = t if isinstance(t, list) else [t]
        if any(tt in article_types for tt in types):
            return n
    return None


def content_root(soup: BeautifulSoup) -> Tag | None:
    """Best-effort main content container for a WP post."""
    # Common WP theme wrappers, in priority order.
    for sel in [
        "div.entry-content",
        "div.post-content",
        "div.entry",
        "article",
        "main",
        "div#content",
        "div.content",
    ]:
        node = soup.select_one(sel)
        if node:
            return node
    return soup.body


def classify_block(tag: Tag) -> str:
    name = tag.name
    if name in ("h1", "h2", "h3", "h4", "h5", "h6"):
        return "heading"
    if name == "p":
        return "paragraph"
    if name == "figure":
        return "image"
    if name == "img" and tag.parent and tag.parent.name != "figure":
        return "image"
    if name == "blockquote":
        return "quote"
    if name in ("ul", "ol"):
        return "list"
    if name == "iframe":
        return "embed"
    if name == "table":
        return "table"
    if name in ("div", "section"):
        # gallery / columns wrappers
        if tag.find("img"):
            return "gallery"
        return "group"
    return "other"


def extract_blocks(root: Tag) -> list[dict]:
    """Walk top-level children of the content root and classify them."""
    blocks: list[dict] = []
    if root is None:
        return blocks
    for child in root.children:
        if isinstance(child, NavigableString):
            text = str(child).strip()
            if text:
                blocks.append({"type": "text", "chars": len(text)})
            continue
        if isinstance(child, Tag):
            btype = classify_block(child)
            text = child.get_text(" ", strip=True)
            blocks.append({
                "type": btype,
                "tag": child.name,
                "text_chars": len(text),
                "text_preview": text[:80],
                "has_image": bool(child.find("img")),
                "has_embed": bool(child.find("iframe")),
            })
    return blocks


def parse_one(kind: str, url: str) -> dict:
    html = fetch(url)
    soup = BeautifulSoup(html, "lxml")
    nodes, by_id = collect_jsonld(soup)
    article = find_article_node(nodes)
    resolved_article = resolve(article, by_id) if article else None

    root = content_root(soup)
    blocks = extract_blocks(root)

    # also collect images inside content for media modeling
    images = []
    if root:
        for img in root.find_all("img"):
            images.append({
                "src": img.get("src") or img.get("data-src"),
                "alt": (img.get("alt") or "")[:60],
                "width": img.get("width"),
                "height": img.get("height"),
            })

    return {
        "kind": kind,
        "url": url,
        "content_selector": root.name if root else None,
        "content_class": " ".join(root.get("class", [])) if root else None,
        "resolved_article_schema": resolved_article,
        "block_count": len(blocks),
        "block_types": dict(Counter(b["type"] for b in blocks)),
        "blocks": blocks,
        "images": images[:10],
        "image_count": len(images),
    }


def main() -> int:
    here = Path(__file__).resolve().parent
    ticket = here.parent
    sources = ticket / "sources"
    raw = sources / "_raw"
    raw.mkdir(parents=True, exist_ok=True)

    records = [parse_one(k, u) for k, u in SAMPLES]

    (raw / "article-blocks.json").write_text(
        json.dumps(records, indent=2, ensure_ascii=False)
    )

    lines: list[str] = []
    lines.append("# Article content-block model (compact)\n")
    lines.append(
        "Models the legacy `content` field as an ordered block list. "
        "Full blocks + resolved Schema.org node are in "
        "`sources/_raw/article-blocks.json`.\n"
    )

    lines.append("## Resolved Schema.org Article node (first sample)\n")
    first = records[0].get("resolved_article_schema")
    # Trim long fields for readability.
    if first:
        slim = {}
        for k, v in first.items():
            if isinstance(v, (str, int, float, type(None))):
                slim[k] = v
            elif isinstance(v, list) and v and isinstance(v[0], str):
                slim[k] = v[:5]
            elif isinstance(v, dict):
                slim[k] = {kk: vv for kk, vv in list(v.items())[:3]}
            else:
                slim[k] = f"<{type(v).__name__}>"
        lines.append("```json\n" + json.dumps(slim, indent=2, ensure_ascii=False) + "\n```\n")

    lines.append("## Block-type histograms\n")
    lines.append("| kind | content container | total blocks | heading | paragraph | image/gallery | list | quote | embed | group/other | images |")
    lines.append("|------|-------------------|--------------|---------|-----------|---------------|------|-------|-------|-------------|--------|")
    for r in records:
        bt = r["block_types"]
        img_like = bt.get("image", 0) + bt.get("gallery", 0)
        group = bt.get("group", 0) + bt.get("other", 0) + bt.get("text", 0)
        lines.append(
            f"| {r['kind']} | `{r.get('content_class') or r.get('content_selector')}` | "
            f"{r['block_count']} | {bt.get('heading',0)} | {bt.get('paragraph',0)} | "
            f"{img_like} | {bt.get('list',0)} | {bt.get('quote',0)} | {bt.get('embed',0)} | "
            f"{group} | {r['image_count']} |"
        )
    lines.append("")

    lines.append("## Block sequence (first 25 blocks, type + preview only)\n")
    for r in records:
        lines.append(f"### {r['kind']} — {r['url']}\n")
        lines.append("| # | type | tag | chars | preview |")
        lines.append("|---|------|-----|-------|---------|")
        for i, b in enumerate(r["blocks"][:25], 1):
            prev = (b.get("text_preview") or "").replace("|", "/")[:60]
            marks = []
            if b.get("has_image"):
                marks.append("img")
            if b.get("has_embed"):
                marks.append("embed")
            lines.append(f"| {i} | {b['type']} | {b['tag']} | {b['text_chars']} | {prev} {' '.join(marks)} |")
        lines.append("")

    (sources / "03-article-blocks.md").write_text("\n".join(lines))
    print(f"Wrote {sources / '03-article-blocks.md'}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
