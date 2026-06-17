# Tasks

## TODO


- [x] P0.1 Init go module (github.com/go-go-golems/fake-cms), create Makefile, .gitignore (done), README stub
- [x] P0.2 Add deps: glazed, graphql-go/graphql, modernc.org/sqlite; pin go.mod; go mod tidy
- [x] P0.3 Scaffold package dirs: internal/{domain,repo,graphql,server,cli}, cmd/fake-cms, migrations, testdata
- [x] P0.4 Root cobra command (fake-cms) wired via glazed; --help lists serve/seed/db
- [ ] P1.1 Write migrations/0001_init.sql (content_node, author, term, media, block, seo, menu) per design doc §8
- [ ] P1.2 repo.Open: open modernc sqlite, apply migrations, enable WAL + foreign_keys; migration idempotency
- [ ] P1.3 repo.Seed: deterministic PRNG generator building authors/terms/media/articles/blocks/seo from sources/_raw
- [ ] P1.4 Commit testdata/cms.db + sha256; verify re-seed reproduces identical hash
- [ ] P2.1 domain structs: Article, Page, Author, Term, Media, Block, SEO (no graphql/sql tags)
- [ ] P2.2 repo: ListArticles/GetArticleBySlug with ArticleFilter + cursor pagination + totalCount
- [ ] P2.3 repo: batched lookups AuthorByID(s), AuthorsByContentIDs, TermsByContentID, MediaByContentID, BlocksByContentID
- [ ] P2.4 repo unit tests against testdata/cms.db + N+1 counting-driver test
- [ ] P3.1 Write schema.graphql (the SDL contract from design §7)
- [ ] P3.2 Global ID encode/decode (base64 Type:rowID) + node(id:) resolver
- [ ] P3.3 Root Query resolvers: article/page/articles/category/tag/author/media/site + DataLoader wiring
- [ ] P3.4 Article field resolvers: author/categories/tags/featuredMedia/related; Connection helpers
- [ ] P3.5 Golden query fixtures: testdata/queries/*.graphql + *.expected.json; schema Exec E2E test
- [ ] P4.1 Block union/interface resolvers (paragraph/heading/image/list/quote/embed/gallery)
- [ ] P4.2 SEO resolver (title/meta/canonical/og/twitter/jsonLd/breadcrumbs) + full-article golden test
- [ ] P5.1 server.New: net/http mux (/graphql, /playground, /healthz) + serve glazed command
- [ ] P5.2 seed + 'db query' glazed commands (structured output) + README quickstart + workshop contract
- [ ] P5.3 Smoke: build, run serve, curl a query, golden E2E; finalize diary + reMarkable re-upload
