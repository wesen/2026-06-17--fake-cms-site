---
Title: "Quickstart"
Slug: quickstart
SubTitle: "Run the mock and send your first query in five minutes"
Short: "Build, seed, serve, and query the fake CMS end to end."
Topics:
- quickstart
- getting-started
- tutorial
Commands:
- serve
- seed
Flags:
- path
- addr
IsTopLevel: true
IsTemplate: false
ShowPerDefault: true
SectionType: Tutorial
Order: 5
---

This tutorial takes you from a clean checkout to a running GraphQL server answering real queries. You need Go 1.25+ and nothing else.

## 1. Build the binary

```bash
make build
```

This produces `./fake-cms`. Verify it runs:

```bash
./fake-cms version
# fake-cms 0.0.0-dev (commit none)
```

## 2. Seed the database

The repository ships a committed seed at `testdata/cms.db`, so you can skip this step. If you want a fresh copy (or a throwaway one), generate it:

```bash
./fake-cms seed --path testdata/cms.db
# seeded testdata/cms.db with 146 content nodes
```

`seed` is deterministic: re-running it reproduces identical content. See `help commands` for the flag vocabulary.

## 3. Start the server

```bash
make serve
# fake-cms listening on :8080 (playground at http://localhost:8080/playground)
```

`make serve` is shorthand for `./fake-cms serve --path testdata/cms.db`. The server stays in the foreground; press `Ctrl+C` to stop.

## 4. Send your first query

Open the playground in a browser:

```text
http://localhost:8080/playground
```

Paste this query and run it:

```graphql
{
  articles(first: 3) {
    totalCount
    edges {
      node {
        slug
        title
        postType
        publishedAt
        author { slug displayName }
      }
    }
    pageInfo { endCursor hasNextPage }
  }
}
```

You should get three articles with nested authors and a `totalCount` around 140.

## 5. Fetch one article in full

```graphql
{
  article(slug: "reportage-engagé--oenobiol-72") {
    title
    excerpt
    seo { canonical jsonLd }
    blocks {
      ... on Block { order }
      ... on ParagraphBlock { text }
      ... on HeadingBlock { level text }
      ... on ImageBlock { media { url width height } caption }
    }
  }
}
```

If the slug above does not exist in your seed, first list slugs with `articles(first: 1) { edges { node { slug } } }` and substitute. Notice the block union: shared fields like `order` come via `... on Block`, type-specific fields via `... on <Type>`.

## 6. Drive pagination

```graphql
query($after: String) {
  articles(first: 50, after: $after) {
    edges { node { slug } }
    pageInfo { endCursor hasNextPage }
  }
}
```

Loop: take the page's `endCursor`, feed it back as `$after`, repeat until `hasNextPage` is false. This is the exact loop an SSG uses to mirror the whole site.

## 7. From the command line

If you prefer `curl` to a browser:

```bash
curl -s localhost:8080/graphql \
  -H 'Content-Type: application/json' \
  -d '{"query":"{ articles(first:2){ totalCount } }"}'
```

And to inspect the underlying data:

```bash
./fake-cms db query --path testdata/cms.db \
  --query "SELECT post_type, count(*) FROM article_post_type GROUP BY post_type"
```

## What you have now

A running, read-only GraphQL API over a deterministic SQLite dataset, plus a help system (`./fake-cms help`) and a SQL escape hatch. The next step is the workshop: build an SSG plugin that renders this content to static HTML. See `help workshop-ssg` for the contract.

## Troubleshooting

| Problem | Cause | Solution |
|---------|-------|----------|
| `make build` downloads for a long time. | First run fetches Go modules. | One-time; subsequent builds are fast. |
| Playground loads but queries hang. | Server not running, or wrong port. | Confirm `curl localhost:8080/healthz` returns `ok`. |
| `cannot find module` for `modernc.org/sqlite`. | Stale module cache. | `go mod tidy && go mod download`. |
| Article slug with accents 404s in your SSG. | URL encoding. | Slugs contain accented UTF-8; encode them. |

## See Also

`help api-reference`, `help commands`, `help workshop-ssg`, `help content-model`.
