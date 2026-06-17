---
Title: "The N+1 guard"
Slug: n-plus-one
SubTitle: "How the API avoids the classic GraphQL N+1 problem"
Short: "Why resolving 20 articles with their authors costs 7 SQL statements, not 100, and how the test enforces it."
Topics:
- performance
- n-plus-one
- internals
Commands: []
Flags: []
IsTopLevel: false
IsTemplate: false
ShowPerDefault: false
SectionType: GeneralTopic
Order: 7
---

The N+1 problem is the single most common performance failure in GraphQL APIs that sit over a database. A query that lists 20 articles and selects each article's `author` naively runs one query for the list plus one query per article to fetch the author — 21 queries. Add `categories`, `tags`, `blocks`, and `featuredMedia` and a single request explodes into hundreds of queries.

## How the mock avoids it

The repository exposes **batched loaders**: one method per relationship that accepts a slice of content ids and returns a map, fetching everything in a single SQL statement.

```go
AuthorsByContentIDs(ctx, ids)            // 1 query for N ids
TermsByContentIDs(ctx, ids)              // 1 query
BlocksByContentIDs(ctx, ids)             // 1 query
FeaturedMediaByContentIDs(ctx, ids)      // 1 query
SEOByContentIDs(ctx, ids)                // 1 query
```

The GraphQL resolvers call these instead of looping per node. Resolving a page of 20 articles with all five relationships therefore costs:

```text
1 (list) + 1 (count) + 5 (one batch per relationship) = 7 SQL statements
```

…regardless of page size.

## How the test enforces it

The test `internal/repo/queries_test.go::TestNoNPlus1` injects a counting `Querier` (see `internal/dbcount`) that wraps `*sql.DB` and tallies every query. It resolves 20 articles plus all relationships and asserts the total is small (currently 7, with a CI ceiling of 10). If a future change reintroduces a per-row lookup, the count rises and the test fails.

Run it with:

```bash
go test ./internal/repo -run TestNoNPlus1 -v
```

## Why a `Querier` interface

The repo depends on a small `Querier` interface (the subset of `*sql.DB` methods it uses) rather than on `*sql.DB` directly. `*sql.DB` satisfies it in production; the test swaps in a counting wrapper via `Repo.WithQuerier`. This is the cleanest way to count queries without a custom database driver.

## A note for the SSG plugin

The same trap exists on the client side. If your plugin fetches each article's author in a separate request, you reproduce the N+1 problem over the network. Fetch the full render payload per article in one query (see `help api-reference` → "One-request full render") and batch list traversal through cursor pagination.

## Troubleshooting

| Problem | Cause | Solution |
|---------|-------|----------|
| `TestNoNPlus1` count jumps above the ceiling. | A resolver started calling a per-id loader in a loop. | Route through a batched `*ByContentIDs` loader. |
| You want to count queries in your own test. | Need an instrumented connection. | Use `dbcount.New(r.DB)` + `Repo.WithQuerier`. |

## See Also

`help api-reference`, `help content-model`.
