---
Title: "Content model"
Slug: content-model
SubTitle: "Why blocks, taxonomies, and SEO look the way they do"
Short: "Explains the WordPress + Yoast content model the API mirrors: post types, taxonomy, blocks, and the SEO layer."
Topics:
- concepts
- content-model
- cms
Commands: []
Flags: []
IsTopLevel: true
IsTemplate: false
ShowPerDefault: true
SectionType: GeneralTopic
Order: 3
---

The schema is not invented from scratch. It mirrors the data model of a **WordPress + Yoast SEO** CMS, because that is what the evidence (the target site's sitemaps and JSON-LD) proves the real system runs. This page explains each modeling choice so the schema stops looking arbitrary.

## Post types: one table, many content kinds

In WordPress, articles, case studies, and studies are not separate tables. They are **custom post types (CPTs)** stored in one `wp_posts` table, distinguished by a `post_type` column. The real target exposes these CPTs as URL prefixes:

| URL prefix | CPT | GraphQL `PostType` |
|------------|-----|--------------------|
| `/best-cases` | `best-cases` | `BEST_CASES` (the bulk of the site) |
| `/actualites` | `actualites` | `ACTUALITES` |
| `/etudes` | `etudes` | `ETUDES` |
| `/blog` | `blog` | `BLOG` |
| `/slider-de-une` | `slider-de-une` | `SLIDER_DE_UNE` |
| `/cartouches-home` | `cartouches-home` | `CARTOUCHES_HOME` |
| `/non-classe` | `non-classe` | `NON_CLASSE` |

The mock stores all of these in a single `content_node` table with `kind = 'ARTICLE'`, linked to `article_post_type`. The GraphQL `Article.postType` field is the discriminator. This is the first piece of "legacy complexity": a query for "latest case studies" must filter by post type.

## Taxonomy: two flavors

WordPress has two built-in taxonomies, and the target uses both:

- **Categories** (`taxonomy = 'category'`) are **hierarchical**. They have a `parent`. On the target they live under `/archives/...`. The mock's `Category.parent` is nullable for roots.
- **Tags** (`taxonomy = 'post_tag'`) are **flat**. On the target they are rendered at `/rubrique/<slug>` — note the unusual URL convention. The SSG must respect the CMS's URL scheme, not invent `/tag/<slug>`.

Both are exposed as `Term` rows in storage and surfaced as distinct GraphQL types (`Category`, `Tag`) so the SSG can treat them differently (breadcrumb trees vs. flat clouds).

## The body is blocks, not HTML

The most important design decision. Modeling `content` as a raw HTML string would be the lazy choice and the wrong one: it would make the SSG exercise trivial and pointless. Instead the body is an **ordered list of typed blocks** (`ParagraphBlock`, `HeadingBlock`, `ImageBlock`, `ListBlock`, `QuoteBlock`, `EmbedBlock`, `GalleryBlock`).

This mirrors how modern WordPress (the block editor, "Gutenberg") stores content, and it forces the SSG to implement a renderer per block variant — which is the actual skill being taught. See `help api-reference` → "Content blocks" for the union contract.

Blocks are stored as rows in a `block` table with an `order_index` and a JSON `data` payload. The GraphQL layer decodes the JSON and dispatches to a concrete type via `ResolveType`.

## The SEO layer (Yoast)

Yoast SEO emits, for every content node: a `<title>`, meta description, canonical URL, `robots` directives, OpenGraph and Twitter cards, and a JSON-LD structured-data graph. The mock reproduces all of this in a 1:1 `seo` table.

The SSG needs these fields to produce correct `<head>` output, so they are first-class GraphQL fields, not afterthoughts. `SEO.jsonLd` is passed through verbatim — the SSG should not reconstruct Schema.org data the CMS already provides.

## Media as first-class nodes

Images (and video, PDFs) are not inline strings. They are `Media` nodes with their own global ids, dimensions, alt text, and MIME type. Blocks reference media by id (`ImageBlock.media`, `GalleryBlock.images`), and `Article.featuredMedia` points at the hero image. This lets the SSG deduplicate assets and generate responsive `srcset` markup from real dimensions.

## Relationships and the N+1 trap

Every article links to one author, several terms, several media, and a list of blocks. A naive resolver that fetches each relationship per row turns "list 20 articles" into 100+ SQL queries. The repository exposes **batched loaders** (`AuthorsByContentIDs`, `TermsByContentIDs`, …) that fetch all relationships for a page of articles in one query each. The GraphQL resolvers call these, keeping the total at a fixed small number. See `help n-plus-one` for the exact counts.

## Legacy warts, kept on purpose

These are intentional, not bugs:

- `Author.email` may be a raw email (`admin@clic-clic.com`) — legacy CMS authorship is messy.
- Case studies are image-led (`wordCount` can be ~80) — the SSG must handle low-text, media-heavy content.
- `Page.template` is a PHP filename (`page-display.php`) — the SSG must map it.
- `search` is `LIKE`-based, not ranked full-text.

These are the texture that makes the mock feel like a real legacy system.

## Troubleshooting

| Problem | Cause | Solution |
|---------|-------|----------|
| Tag URLs differ from the real site. | You used `/tag/` instead of `/rubrique/`. | Respect the CMS convention; tags live at `/rubrique/<slug>`. |
| A category appears twice in breadcrumbs. | You ignored the `parent` hierarchy. | Walk `Category.parent` to build the path. |

## See Also

`help api-reference`, `help n-plus-one`, `help workshop-ssg`.
