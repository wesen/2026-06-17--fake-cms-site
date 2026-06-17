# Changelog

## 2026-06-17

- Initial workspace created


## 2026-06-17

Researched 20minutes-media.com (WordPress+Yoast) via scripts; captured evidence in sources/. Wrote SQLite-backed GraphQL mock design + intern implementation guide. Started diary.

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/ttmp/2026/06/17/FAKE-CMS--fake-internal-cms-graphql-api-20minutes-style-for-ssg-workshop/design-doc/01-fake-cms-graphql-api-design-implementation-guide.md — Primary deliverable
- /home/manuel/code/wesen/2026-06-17--fake-cms-site/ttmp/2026/06/17/FAKE-CMS--fake-internal-cms-graphql-api-20minutes-style-for-ssg-workshop/reference/01-investigation-diary.md — Investigation diary
- /home/manuel/code/wesen/2026-06-17--fake-cms-site/ttmp/2026/06/17/FAKE-CMS--fake-internal-cms-graphql-api-20minutes-style-for-ssg-workshop/sources/00-sitemap-inventory.md — Proves WP+Yoast and post-type/taxonomy surface
- /home/manuel/code/wesen/2026-06-17--fake-cms-site/ttmp/2026/06/17/FAKE-CMS--fake-internal-cms-graphql-api-20minutes-style-for-ssg-workshop/sources/03-article-blocks.md — Article type + Block union modeled from this


## 2026-06-17

Validated (doctor clean), added frontmatter to sources/, uploaded design+diary bundle to reMarkable at /ai/2026/06/17/FAKE-CMS.

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/ttmp/2026/06/17/FAKE-CMS--fake-internal-cms-graphql-api-20minutes-style-for-ssg-workshop/design-doc/01-fake-cms-graphql-api-design-implementation-guide.md — Uploaded to reMarkable


## 2026-06-17

P0 done (commit c4876e0): go module + glazed CLI skeleton; version command works; go build clean

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/cmd/fake-cms/main.go — Root entrypoint wired via glazed
- /home/manuel/code/wesen/2026-06-17--fake-cms-site/internal/cli/cli.go — Shared field sections + command registration


## 2026-06-17

P1 done (commits dfb10da, 8e8935d): migrations + repo.Open + deterministic seed; testdata/cms.db committed (146 nodes, 1378 blocks); content-deterministic

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/internal/repo/repo.go — Open/Migrate/Recreate
- /home/manuel/code/wesen/2026-06-17--fake-cms-site/internal/repo/seed.go — Deterministic seeder
- /home/manuel/code/wesen/2026-06-17--fake-cms-site/testdata/cms.db — Committed seed artifact


## 2026-06-17

Implemented full mock (P0-P5, commits c4876e0..a79c029): glazed CLI, SQLite repo+seed, schema-first GraphQL, block union, net/http server. N+1 guard = 7 SQL ops for 20 articles. Live server smoke green.

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/internal/server/server.go — HTTP server
- /home/manuel/code/wesen/2026-06-17--fake-cms-site/schema.graphql — GraphQL contract


## 2026-06-17

Added glazed help entries (detailed API spec) bundled into the app via embedded HelpSystem; rewrote README (commit 7b79207). 'fake-cms help api-reference' works offline.

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/README.md — Polished README
- /home/manuel/code/wesen/2026-06-17--fake-cms-site/internal/doc/api-reference.md — Detailed API spec help entry


## 2026-06-17

Backfilled diary Step 10 (glazed help API spec + README, commit 7b79207). Diary now covers all 10 steps; noted that earlier steps were backfilled rather than incremental.

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/ttmp/2026/06/17/FAKE-CMS--fake-internal-cms-graphql-api-20minutes-style-for-ssg-workshop/reference/01-investigation-diary.md — Diary backfill

