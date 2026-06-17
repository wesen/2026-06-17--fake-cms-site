# fake-cms

A **fake internal CMS GraphQL API** used as the data source for a workshop in
which participants build a **static site generator (SSG) plugin** that renders
CMS content to static HTML.

The mock **mirrors a legacy WordPress + Yoast** media CMS (modelled on
`20minutes-media.com`). It is **read-only**, **SQLite-backed**, and
**deterministic** (a committed seed file).

> ⚠️ **FAKE DATA — DO NOT EXPOSE PUBLICLY.** This is an unauthenticated,
> read-only mock. It is intended for local workshops only.

## Quick start

```bash
make build                 # build ./fake-cms
./fake-cms seed --db-path testdata/cms.db   # (re)generate the deterministic DB
./fake-cms serve           # serve GraphQL at http://localhost:8080/graphql
```

Open the playground at <http://localhost:8080/playground>.

## Architecture (layers point up only)

```
cmd/fake-cms          glazed Cobra CLI (serve, seed, db)
internal/server       net/http: /graphql, /playground, /healthz
internal/graphql      graphql-go schema-first resolvers + DataLoader
internal/repo         database/sql -> domain structs (no graphql)
internal/domain       plain structs
migrations/0001_init.sql
testdata/cms.db       committed deterministic seed
schema.graphql        THE contract (single source of truth)
```

See `ttmp/2026/06/17/FAKE-CMS--*/design-doc/01-*.md` for the full design.

## Workshop contract (SSG plugin)

Participants query `articles(first:, after:)` (Relay cursor pagination),
follow `pageInfo.endCursor`, then render per `postType`/`author`/`category`/
`tag` (tags live at `/rubrique/<slug>/`, mirroring the real CMS URLs). The body
is a typed `Block` union — render each variant.
