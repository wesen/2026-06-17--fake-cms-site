---
Title: "GraphQL Content Delivery API"
Slug: api-reference
SubTitle: "Read-only content API surface for the SSG workshop"
Short: "Complete reference for every Query, type, enum, scalar, connection, and block variant exposed by the fake CMS."
Topics:
- api
- graphql
- reference
Flags: []
Commands: []
IsTopLevel: true
IsTemplate: false
ShowPerDefault: true
SectionType: GeneralTopic
Order: 2
---

The Content Delivery API is a **read-only GraphQL endpoint** served at `POST /graphql` (and `GET /graphql?query=...`). It is the single contract the static-site-generator plugin talks to. This page documents every field you can query, why it exists, and what real-world (legacy) shape it mirrors.

The API is unauthenticated by design: this is a local workshop mock containing synthetic data. Never expose it to the public internet.

## Endpoint

```text
POST http://localhost:8080/graphql
Content-Type: application/json

{ "query": "<GraphQL document>", "variables": { ... }, "operationName": "..." }
```

A single request returns exactly the fields the document selects — no over-fetching, no under-fetching. Field errors do not change the HTTP status (always `200`); read the top-level `errors` array in the response body, per the GraphQL spec.

## Scalars

| Scalar | Wire format | Notes |
|--------|-------------|-------|
| `ID` | string | Relay global id, e.g. `QXJ0aWNsZTo3Mg==` (base64 of `Article:72`). |
| `String` | string | |
| `Int` | number | |
| `Boolean` | boolean | |
| `DateTime` | ISO-8601 string | `2013-11-19T11:11:20Z`; nullable for drafts. |
| `URL` | string | Absolute URL. |
| `JSON` | any | Opaque JSON blob. Used for `SEO.jsonLd` and OpenGraph/Twitter payloads so the SSG can emit them verbatim. |

## Global ids and `node(id:)`

Every entity that implements the `Node` interface has a globally-unique `id` of the form `base64("<TypeName>:<rowID>")`. The top-level `node(id:)` field decodes the type prefix and dispatches to the right loader, so you can refetch any object by id without knowing its concrete type.

```graphql
{ node(id: "QXJ0aWNsZTo3Mg==") { ... } }
```

To find the id of a concrete type, query its `id` field — it is always the global id, never the raw row id.

## Enums

```graphql
enum PostType {
  ACTUALITES       # /actualites news & announcements
  BEST_CASES       # /best-cases advertising case studies (the bulk of the site)
  ETUDES           # /etudes audience studies
  BLOG             # /blog
  SLIDER_DE_UNE    # homepage hero slider
  CARTOUCHES_HOME  # homepage cards
  NON_CLASSE       # uncategorized (legacy bucket)
}

enum ArticleStatus { DRAFT PUBLISH FUTURE PRIVATE TRASH }
enum OrderField    { PUBLISHED_AT MODIFIED_AT TITLE MENU_ORDER }
enum OrderDirection { ASC DESC }
enum MediaKind     { IMAGE VIDEO PDF AUDIO FILE }
enum Locale        { FR_FR EN_US DE_DE }   # slot only; v1 ships FR_FR
enum Align         { LEFT CENTER RIGHT NONE }
```

## Top-level Query fields

| Field | Arguments | Returns | Purpose |
|-------|-----------|---------|---------|
| `node` | `id: ID!` | `Node` | Refetch any entity by global id. |
| `article` | `id: ID`, `slug: String` | `Article` | Fetch one article by global id or slug. |
| `page` | `id: ID`, `slug: String` | `Page` | Fetch one page by global id or slug. |
| `articles` | `filter`, `first`, `after`, `orderBy` | `ArticleConnection!` | Paginated, filtered article list. See "Connections". |
| `categories` | `first: Int = 20` | `[Category!]!` | All categories. |
| `category` | `slug: String!` | `Category` | One category by slug. |
| `tags` | `first: Int = 20` | `[Tag!]!` | All tags (rendered as `/rubrique/<slug>` on the real site). |
| `tag` | `slug: String!` | `Tag` | One tag by slug. |
| `authors` | `first: Int = 20` | `[Author!]!` | All authors. |
| `author` | `slug: String!` | `Author` | One author by slug. |
| `media` | `id: ID!` | `Media` | One media asset by global id. |
| `site` | — | `Site!` | Site config + menus. |

## Filtering and ordering `articles`

The `articles` field accepts an `ArticleFilter` input and an `ArticleOrder`. Filtering is how you mimic the real site's section pages (`/best-cases`, `/actualites`, archives).

```graphql
input ArticleFilter {
  postType: PostType
  status: ArticleStatus = PUBLISH
  categorySlug: String
  tagSlug: String
  authorSlug: String
  search: String            # legacy: LIKE over title + excerpt, not ranked full-text
  publishedAfter: DateTime
  publishedBefore: DateTime
}
input ArticleOrder { field: OrderField! direction: OrderDirection! = DESC }
```

Example: latest case studies.

```graphql
{
  articles(filter: { postType: BEST_CASES }, first: 10) {
    edges { node { slug title publishedAt } }
    pageInfo { endCursor hasNextPage }
  }
}
```

The `search` argument is intentionally primitive (SQL `LIKE`). Legacy CMS search is rarely ranked; the SSG is expected to do its own relevance handling if it cares.

## Connections (Relay cursor pagination)

List fields that can grow unbounded return a `Connection`, the de-facto Relay pattern. Cursors are **stable under insertion**: paging again with the previous `endCursor` will not skip or duplicate rows even if new articles arrive.

```graphql
type ArticleConnection {
  edges: [ArticleEdge!]!
  pageInfo: PageInfo!
  totalCount: Int!          # total matching the filter, ignoring pagination
}
type ArticleEdge { node: Article! cursor: String! }
type PageInfo {
  hasNextPage: Boolean!
  hasPreviousPage: Boolean!
  startCursor: String
  endCursor: String
}
```

Cursors are opaque (`base64("<publishedAt>:<rowId>")`). Do not parse them; pass them back as `after`. The canonical pagination loop is:

```text
after = null
loop:
  page = query articles(first: 50, after: after)
  emit page.edges
  if not page.pageInfo.hasNextPage: break
  after = page.pageInfo.endCursor
```

## Core types

### Article

```graphql
type Article implements Content & Node {
  id: ID!
  slug: String!
  title: String!
  excerpt: String
  status: ArticleStatus!
  postType: PostType!          # discriminator for the legacy "one table, many types" model
  locale: Locale!
  publishedAt: DateTime
  modifiedAt: DateTime!
  author: Author!
  categories: [Category!]!
  tags: [Tag!]!
  featuredMedia: Media
  blocks: [BlockUnion!]!       # ordered body — see "Content blocks"
  seo: SEO!
  wordCount: Int
  related(first: Int = 5): [Article!]!   # same postType + shared tags
}
```

`author`, `categories`, `tags`, `featuredMedia`, `blocks`, and `seo` are all resolved through the repository's batched loaders, so fetching 20 articles with their authors costs a fixed small number of SQL statements, not one per row (see `help n-plus-one`).

### Page

```graphql
type Page implements Content & Node {
  id: ID!
  slug: String!
  title: String!
  status: ArticleStatus!
  locale: Locale!
  publishedAt: DateTime
  modifiedAt: DateTime!
  parent: Page                # hierarchical pages (nullable for roots)
  menuOrder: Int!
  template: String            # legacy: a PHP filename like "page-display.php"
  featuredMedia: Media
  blocks: [BlockUnion!]!
  seo: SEO!
}
```

`Page.template` is deliberately ugly (a PHP filename) so the SSG plugin learns to map CMS conventions rather than invent its own.

### Author / Category / Tag / Media

```graphql
type Author implements Node {
  id: ID! slug: String! displayName: String!
  email: String               # legacy: may be a raw email, may be null
  description: String
  avatar: Media
  locale: Locale!
}
type Category implements Node {
  id: ID! slug: String! name: String! description: String
  parent: Category            # hierarchical
}
type Tag implements Node {
  id: ID! slug: String! name: String! description: String
}
type Media implements Node {
  id: ID! slug: String! kind: MediaKind! url: URL! sourceUrl: URL!
  alt: String width: Int height: Int mimeType: String fileSize: Int caption: String
  locale: Locale!
}
```

Note the legacy wart preserved on purpose: `Author.email` can be a raw address (the real site has `admin@clic-clic.com` as an author). Do not assume a clean display name.

## Content blocks (`BlockUnion`)

The body of any `Article` or `Page` is an **ordered list of typed blocks**, exposed as a GraphQL union. This is the heart of the workshop: the SSG must render each variant. Blocks implement the `Block` interface (shared `id` / `order`) and are members of `BlockUnion`.

```graphql
interface Block { id: ID! order: Int! }
union BlockUnion =
    ParagraphBlock | HeadingBlock | ImageBlock
  | ListBlock | QuoteBlock | EmbedBlock | GalleryBlock
```

Because it is a union, you **cannot** select `order` directly on `BlockUnion`; spread the interface to read shared fields, then inline-fragment each variant for its specific fields:

```graphql
blocks {
  ... on Block { id order }
  ... on HeadingBlock   { level text anchor }
  ... on ParagraphBlock { text align }
  ... on ImageBlock     { media { url width height } caption size }
  ... on ListBlock      { ordered items }
  ... on QuoteBlock     { text citation }
  ... on EmbedBlock     { provider url caption }
  ... on GalleryBlock   { images { url } columns }
}
```

| Block | Type-specific fields | Notes |
|-------|----------------------|-------|
| `ParagraphBlock` | `text`, `align: Align` | The most common block. |
| `HeadingBlock` | `level: Int` (`2`–`4`), `text`, `anchor` | `anchor` is the in-page slug for deep links. |
| `ImageBlock` | `media: Media!`, `caption`, `link: URL`, `size` | `media` is non-null and resolves to a real `Media` node. |
| `ListBlock` | `ordered: Boolean!`, `items: [String!]!` | |
| `QuoteBlock` | `text`, `citation` | |
| `EmbedBlock` | `provider`, `url`, `caption` | `provider` e.g. `youtube`. |
| `GalleryBlock` | `images: [Media!]!`, `columns: Int` | Each image resolves to a `Media` node. |

## SEO (Yoast-style)

Every content node has a non-null `seo` object mirroring the Yoast SEO plugin the real site runs.

```graphql
type SEO {
  title: String!
  metaDescription: String
  canonical: URL
  robots: String              # e.g. "index,follow"
  og: OpenGraph               # OpenGraph card (emitted as JSON in v1)
  twitter: TwitterCard        # Twitter card (emitted as JSON in v1)
  jsonLd: JSON                # raw Yoast JSON-LD graph for the resource
  breadcrumbs: [Breadcrumb!]!
}
type Breadcrumb { label: String! url: URL! }
```

`jsonLd` is the exact structured-data graph (Schema.org `Article`, `BreadcrumbList`, etc.). The SSG should inject it verbatim into the page `<head>`; do not re-derive it.

## One-request full render

The API is designed so the SSG can fetch a complete render payload in a single query. This is the recommended query shape for an article page:

```graphql
query RenderArticle($slug: String!) {
  article(slug: $slug) {
    id slug title excerpt postType wordCount
    publishedAt modifiedAt status locale
    author { id slug displayName description }
    categories { id slug name }
    tags { id slug name }
    featuredMedia { id url width height alt }
    seo { title metaDescription canonical robots jsonLd breadcrumbs }
    blocks {
      ... on Block { id order }
      ... on ParagraphBlock { text }
      ... on HeadingBlock { level text anchor }
      ... on ImageBlock { media { id url width height } caption size }
      ... on ListBlock { ordered items }
      ... on QuoteBlock { text citation }
      ... on EmbedBlock { provider url caption }
      ... on GalleryBlock { images { id url } columns }
    }
  }
}
```

## Troubleshooting

| Problem | Cause | Solution |
|---------|-------|----------|
| `Cannot query field "order" on type "BlockUnion"` | Selecting an interface field directly on a union. | Spread the interface: `... on Block { order }`. |
| `Cannot query field "media" on type "Block"` | Same — `media` lives on `ImageBlock`, not the interface. | Use `... on ImageBlock { media { url } }`. |
| Nested `author` returns `null`. | The content node has no author (e.g. some pages). | `Article.author` is non-null; only `Page` may lack one. |
| Pagination repeats/overlaps rows. | Using offsets instead of the returned cursor. | Always pass `pageInfo.endCursor` as `after`. |
| `totalCount` differs from `edges.length`. | Expected — `totalCount` ignores pagination. | Use it for "N results" UI; use `edges` for the page. |
| HTTP 200 but response has `errors`. | Correct per GraphQL spec; field-level failure. | Inspect `errors[0].message` and `path`. |

## See Also

- `help quickstart` — get the server running in five minutes.
- `help content-model` — why blocks, taxonomy, and SEO are modeled this way.
- `help commands` — the `serve`, `seed`, and `db` operator commands.
- `help workshop-ssg` — the SSG plugin contract and required output.
