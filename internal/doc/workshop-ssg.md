---
Title: "Workshop: build an SSG plugin"
Slug: workshop-ssg
SubTitle: "The contract for the static-site-generator plugin"
Short: "What the SSG plugin must do: query, walk, paginate, and render the CMS content to static HTML matching the real site's URL conventions."
Topics:
- workshop
- ssg
- contract
Commands:
- serve
Flags:
- path
IsTopLevel: true
IsTemplate: false
ShowPerDefault: true
SectionType: Application
Order: 6
---

This is the deliverable specification for the workshop. Participants build a **static site generator plugin** that consumes the fake CMS GraphQL API and emits a static site that mirrors the structure and URL conventions of the real `20minutes-media.com` site. The mock exists precisely to make this exercise realistic.

## Goal

Produce a directory of static files (HTML, plus copied/referenced assets) that, when served by any static host, reproduces the content of the CMS: every article, every taxonomy archive, every author page, plus a working sitemap.

## Step 1 — Introspect the contract

Read `help api-reference` and `schema.graphql`. The schema is the contract; do not hard-code assumptions the schema does not state. If a field is nullable, handle its absence.

## Step 2 — Walk the content graph

Use cursor pagination (see `help api-reference` → "Connections"). The canonical loop:

```text
after = null
loop:
  page = articles(first: 50, after: after) { edges { node { id slug postType } } pageInfo { endCursor hasNextPage } }
  collect page.edges
  if not page.pageInfo.hasNextPage: break
  after = page.pageInfo.endCursor
```

Then, per article, fetch the full render payload in one query (the "one-request full render" shape in `help api-reference`). Fetching per-field in many requests defeats the purpose and risks N+1 behavior on the client side.

## Step 3 — Respect the URL conventions

The real site uses these URL schemes; the SSG must reproduce them so its output is a faithful mirror:

| Entity | URL | Source |
|--------|-----|--------|
| Article | `/<postTypeSlug>/<slug>/index.html` | `article.slug` + the WP slug for `postType` |
| Page | `/<slug>/index.html` | `page.slug` |
| Category archive | `/archives/<slug>/index.html` | `category.slug` |
| Tag archive | `/rubrique/<slug>/index.html` | `tag.slug` (**not** `/tag/`) |
| Author | `/author/<slug>/index.html` | `author.slug` |

Map the GraphQL `PostType` enum to the WP slug (`BEST_CASES` → `best-cases`, etc.). The teaching point: the SSG respects the CMS's URL scheme rather than inventing its own. Tags at `/rubrique/` is a real legacy detail on the target.

## Step 4 — Render the blocks

The body is `[BlockUnion!]!`. Implement a renderer per variant. Pseudocode:

```text
for block in article.blocks:
    switch block.__typename:
        ParagraphBlock -> <p align=...>block.text</p>
        HeadingBlock   -> <h{block.level} id=block.anchor>block.text</h{block.level}>
        ImageBlock     -> <figure><img src=block.media.url ...><figcaption>block.caption</figcaption></figure>
        ListBlock      -> <ul|ol> for item in block.items: <li>item</li> </...>
        QuoteBlock     -> <blockquote>block.text<cite>block.citation</cite></blockquote>
        EmbedBlock     -> provider-specific embed (iframe for youtube, etc.)
        GalleryBlock   -> grid of <img> over block.images
```

This is the core of the exercise. A renderer that dumps raw HTML would not exist, because the API never gives raw HTML.

## Step 5 — Emit the `<head>`

Use the `seo` object verbatim where possible:

```text
<title>{seo.title}</title>
<meta name="description" content="{seo.metaDescription}">
<link rel="canonical" href="{seo.canonical}">
<meta name="robots" content="{seo.robots}">
<!-- OpenGraph / Twitter from seo.og / seo.twitter -->
<script type="application/ld+json">{seo.jsonLd}</script>   <!-- inject verbatim -->
```

Inject `seo.jsonLd` directly; do not reconstruct Schema.org JSON the CMS already produced.

## Step 6 — Build the index pages

For each post type and each taxonomy term, generate a listing page that paginates the related articles (reuse the `articles` query with a filter, e.g. `{ postType: BEST_CASES }` or `{ tagSlug: "audience" }`). Add a homepage that curates the latest items per section.

## Step 7 — Generate a sitemap

Emit `sitemap.xml` from the collected slugs. This closes the loop with the real site's own sitemap (which is what the mock's content model was derived from).

## Step 8 — Verify

A correct plugin should be able to round-trip: the number of article URLs it emits equals `articles(first: 0) { totalCount }`, every emitted page has a non-empty `<title>` from `seo.title`, and every block type renders without an "unknown block" fallback.

## Acceptance criteria

1. Every article in the API has a corresponding static page at the correct URL.
2. Tag pages live at `/rubrique/<slug>/`, not `/tag/<slug>/`.
3. Every page injects `seo.jsonLd` verbatim.
4. All seven block types render distinctly (no two collapse to the same markup).
5. A `sitemap.xml` covers all generated pages.
6. The whole site builds from `fake-cms serve` with no manual data entry.

## Troubleshooting

| Problem | Cause | Solution |
|---------|-------|----------|
| Tag URLs render at `/tag/`. | Inventing conventions. | Tags are `/rubrique/<slug>/`. |
| Pages missing `<title>`. | Not reading `seo.title`. | The `seo` object is non-null; use it. |
| Two block types render identically. | Falling back to a generic renderer. | Implement a distinct branch per `__typename`. |
| Article count mismatch in sitemap. | Not following `hasNextPage`. | Loop pagination to completion. |
| Accents in slugs break the build. | Not URL-encoding. | Slugs contain accented UTF-8; handle encoding. |

## See Also

`help api-reference`, `help content-model`, `help quickstart`, `help commands`.
