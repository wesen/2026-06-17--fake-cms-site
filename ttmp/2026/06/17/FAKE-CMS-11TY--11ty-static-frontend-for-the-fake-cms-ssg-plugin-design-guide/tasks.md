# Tasks

## TODO

### Research & design (this ticket)
- [x] Create ticket + workspace
- [x] Download 11ty official docs into sources/ with defuddle (9 pages)
- [x] Write sources manifest (00-sources-readme.md)
- [x] Add vocabulary (11ty, ssg, frontend, javascript)
- [x] Write intern-facing design & implementation guide (design-doc/01-…)
- [x] Write investigation diary (reference/02-…)
- [x] Relate files + update changelog
- [x] docmgr doctor (clean)
- [x] Upload design+diary bundle to reMarkable

### Implementation (future — see guide §18)
- [ ] Phase 0: scaffold 11ty site + vitest; smoke build
- [ ] Phase 1: lib/graphql-client.js (gql + fetchAllArticles cursor loop)
- [ ] Phase 2: plugin core (index.js, global-data.js, filters.js); first article pages
- [ ] Phase 3: lib/blocks.js (7 branches + throw-on-unknown); head.njk SEO/JSON-LD
- [ ] Phase 4: archives/rubriques/authors/pages templates; registerCollections; homepage
- [ ] Phase 5: sitemap.xml; opt-in @11ty/eleventy-fetch caching; README; E2E vs fake-cms serve

### Acceptance gate (guide §17.2)
- [ ] Article-page count == articles.totalCount
- [ ] Tags at /rubrique/, never /tag/
- [ ] seo.jsonLd injected verbatim
- [ ] All 7 block types render distinctly; renderBlocks throws on unknown
- [ ] sitemap covers all pages
- [ ] Full build from `fake-cms serve` with no manual data
