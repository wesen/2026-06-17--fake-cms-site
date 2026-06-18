# Tasks

## TODO

### Documentation / ticket baseline
- [x] D0.1 Create FAKE-CMS-11TY ticket workspace.
- [x] D0.2 Download official Eleventy docs with defuddle into `sources/`.
- [x] D0.3 Write first intern-facing 11ty design guide.
- [x] D0.4 Review first guide against executable fake-cms schema and Eleventy 3.1.6.
- [x] D0.5 Write corrected 11ty + fake-cms static-site guide.
- [x] D0.6 Upload corrected guide bundle to reMarkable.
- [x] D0.7 Commit ticket docs and sources (`33c6fa7`).

### Implementation plan
- [x] P0.1 Add frontend project skeleton under `frontend/` with package.json, Eleventy config, source/layout directories, and ignore rules for `_site`, `.cache`, and `node_modules`.
- [x] P0.2 Add a tiny hardcoded CMS global-data smoke path proving Eleventy pagination can emit one article page.
- [x] P0.3 Add npm scripts for `build`, `serve`, `test`, `clean`, and a README with the seeded fake-cms command.

- [x] P1.1 Add a GraphQL client (`gql`) that POSTs `{query, variables}` and throws on HTTP or GraphQL errors.
- [x] P1.2 Add `fetchAllArticles` with Relay cursor pagination, using the correct `BlockUnion` fragment (`... on Block { id order }`).
- [x] P1.3 Add fetchers for current executable schema lists: `categories`, `tags`, `authors`, and optional configured `pageSlugs`.
- [x] P1.4 Add a contract test or smoke command that verifies `testdata/cms.db` returns article data and catches the invalid old `blocks { order }` query.

- [x] P2.1 Add normalization helpers: `postTypeSlug`, `pathSegment`, URL derivation, and entity `kind` tagging.
- [x] P2.2 Build indexes once in normalization: `articlesByPostType`, `articlesByTagSlug`, `articlesByCategorySlug`, `articlesByAuthorSlug`.
- [x] P2.3 Build `sitemapUrls` from normalized articles, pages, categories, tags, and authors.
- [x] P2.4 Unit-test normalization, especially `/rubrique/` tag URLs and accented slugs.

- [x] P3.1 Add `renderBlocks` with distinct markup for all seven block types.
- [x] P3.2 Escape HTML attributes/text and clamp heading levels safely.
- [x] P3.3 Throw on unknown block types.
- [x] P3.4 Unit-test every block branch plus unknown-block failure.

- [x] P4.1 Convert the local plugin to real CMS mode: `addGlobalData("cms", async () => normalizeCms(await fetchCms(...)))`.
- [x] P4.2 Register filters/shortcodes: `postTypeSlug`, `pathSegment`, `json`, `absoluteUrl`, `renderBlocks`.
- [x] P4.3 Keep the plugin project-local (`frontend/_config/fakeCmsPlugin.cjs`) until the backend schema stabilizes.

- [x] P5.1 Add visible Nunjucks templates: `articles.njk`, `tag.njk`, `category.njk`, `author.njk`, `post-type.njk`, `index.njk`, `sitemap.xml.njk`.
- [x] P5.2 Add `_includes/base.njk`, `_includes/head.njk`, and reusable cards/partials.
- [x] P5.3 Ensure `seo.jsonLd` is serialized with a JSON filter and injected as `<script type="application/ld+json">`.
- [x] P5.4 Add minimal CSS and passthrough copy so pages are readable.

- [x] P6.1 Add Eleventy programmatic integration test that builds against a running seeded fake-cms server.
- [x] P6.2 Assert article-page count equals `articles.totalCount`.
- [x] P6.3 Assert no `/tag/` URLs and at least one `/rubrique/` URL.
- [x] P6.4 Assert `sitemap.xml` includes generated content URLs.
- [x] P6.5 Assert representative article HTML includes valid JSON-LD.

- [ ] P7.1 Document known backend contract mismatches in `frontend/README.md`.
- [ ] P7.2 Update ticket diary after each committed implementation step.
- [ ] P7.3 Run `npm test`, `npm run build`, `go test ./...`, and `docmgr doctor` before final handoff.
- [ ] P7.4 Commit implementation in focused intervals matching completed phases.

### Acceptance gate
- [x] A1 Every article in `testdata/cms.db` has a corresponding static HTML page at `/<postTypeSlug>/<slug>/`.
- [x] A2 Tag pages live at `/rubrique/<slug>/`, never `/tag/<slug>/`.
- [x] A3 Article pages inject `seo.jsonLd` as valid JSON-LD.
- [x] A4 All seven block types render distinctly; unknown block types fail the build/test.
- [x] A5 `sitemap.xml` covers generated article, archive, tag, author, and optional page URLs.
- [x] A6 The site builds from `./fake-cms serve --path testdata/cms.db` with no manual data entry.
