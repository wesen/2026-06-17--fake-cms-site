---
Title: "Commands"
Slug: commands
SubTitle: "serve, seed, db, version — the operator surface"
Short: "Reference for every fake-cms CLI command, its flags, and when to use it."
Topics:
- cli
- commands
- operations
Commands:
- serve
- seed
- db
- version
Flags:
- path
- addr
- recreate
- query
IsTopLevel: true
IsTemplate: false
ShowPerDefault: true
SectionType: GeneralTopic
Order: 4
---

`fake-cms` is a glazed Cobra CLI. Every subcommand shares a consistent flag vocabulary because flags come from glazed **field sections** (groups like `db`, `http`, `sql`). This page lists each command, its flags, and a realistic invocation.

To see the exact flag names for a command, run `fake-cms <command> --help`. Flags are prefixed by their section, so the `db` section's `path` field becomes `--path`.

## `fake-cms version`

Prints the version and commit. Useful to confirm the binary was built from the right source.

```bash
fake-cms version
# fake-cms 0.0.0-dev (commit none)
```

## `fake-cms seed`

Creates or recreates the SQLite database and fills it with the **deterministic** dataset. Use this to bootstrap `testdata/cms.db` or a fresh DB for a workshop attendee.

```bash
fake-cms seed --path testdata/cms.db
# seeded testdata/cms.db with 146 content nodes
```

| Flag | Section | Default | Meaning |
|------|---------|---------|---------|
| `--path` | `db` | `cms.db` | Path to the SQLite file (created/migrated if missing). |
| `--recreate` | `db` | `false` | Drop and recreate the DB first. `seed` always recreates internally for determinism. |

Because the generator is seeded (`rand.NewSource(42)`), running `seed` twice on an empty target yields identical content. The committed artifact ships with a content hash at `testdata/cms.db.content-sha256`.

## `fake-cms serve`

Opens the database, builds the GraphQL schema, and serves the HTTP API. This is the command participants run during the workshop.

```bash
fake-cms serve --path testdata/cms.db
# listening on :8080 (playground at http://localhost:8080/playground)
```

| Flag | Section | Default | Meaning |
|------|---------|---------|---------|
| `--path` | `db` | `cms.db` | SQLite file to open (migrated if missing). |
| `--addr` | `http` | `:8080` | Listen address. |

The server exposes:

- `POST /graphql` — the GraphQL endpoint (also `GET /graphql?query=...`).
- `GET /playground` — an embedded GraphiQL query browser (no install needed).
- `GET /healthz` — returns `ok` (for liveness checks).
- `GET /` — a landing page linking to the playground.

Press `Ctrl+C` to stop; the server shuts down gracefully.

## `fake-cms db query`

A read-only SQL escape hatch. Runs a `SELECT` against the database and prints tab-separated rows. Useful for poking at the seed while developing the SSG, or for debugging why a GraphQL result looks the way it does.

```bash
fake-cms db query --path testdata/cms.db \
  --query "SELECT kind, count(*) AS n FROM content_node GROUP BY kind"
# kind    n
# ARTICLE 140
# PAGE    6
# (2 rows)
```

| Flag | Section | Default | Meaning |
|------|---------|---------|---------|
| `--path` | `db` | `cms.db` | SQLite file. |
| `--query` | `sql` | *(required)* | The SQL `SELECT` to run. |

This command is intentionally read-only and tab-separated so its output pipes cleanly into `awk`, `sort`, or a CSV converter.

## `fake-cms help`

The glazed help browser. With no argument it lists all help topics; with a slug it renders that page.

```bash
fake-cms help                 # list topics
fake-cms help api-reference   # the detailed API spec
fake-cms help quickstart
```

The help entries are embedded in the binary (see `internal/doc/doc.go`), so `fake-cms help` works with no network access.

## Troubleshooting

| Problem | Cause | Solution |
|---------|-------|----------|
| `serve` prints `--help` instead of starting. | An unknown flag (e.g. `--db-path`). | Flags are section-prefixed: use `--path`, not `--db-path`. |
| `db query` returns "sql.query is required". | The flag is `--query`, not `--sql`. | Use `--query "SELECT ..."`. |
| `serve` exits immediately with "address already in use". | Another process holds the port. | Change with `--addr :8090`. |
| `seed` overwrote my DB. | `seed` always recreates. | Point it at a throwaway path, e.g. `--path /tmp/workshop.db`. |

## See Also

`help quickstart`, `help api-reference`, `help overview`.
