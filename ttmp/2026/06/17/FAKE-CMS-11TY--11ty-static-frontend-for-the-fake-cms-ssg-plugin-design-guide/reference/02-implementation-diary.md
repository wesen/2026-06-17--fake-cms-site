---
Title: 11ty Frontend — Investigation Diary
Ticket: FAKE-CMS-11TY
Status: active
Topics:
    - 11ty
    - frontend
    - research
DocType: reference
Intent: long-term
Owners: []
RelatedFiles:
    - Path: frontend/_config/fakeCmsClient.cjs
      Note: GraphQL fetcher and current executable schema query shapes
    - Path: frontend/_config/fakeCmsPlugin.cjs
      Note: Project-local Eleventy plugin boundary
    - Path: frontend/_config/renderBlocks.cjs
      Note: Seven-variant block renderer
    - Path: frontend/scripts/integration-build.mjs
      Note: End-to-end frontend acceptance script
ExternalSources: []
Summary: 'Chronological investigation diary: orienting on the CMS contract, downloading 11ty docs as evidence, and writing the design guide.'
LastUpdated: 2026-06-17T18:05:00-04:00
WhatFor: Resume context for the 11ty frontend ticket.
WhenToUse: Read before continuing work on FAKE-CMS-11TY.
---





# Diary

## Goal

Capture how the **11ty frontend** ticket (`FAKE-CMS-11TY`) was researched and
designed: how the 11ty plugin model was investigated, which official docs were
captured as evidence, why the architecture is what it is, and how to continue.

## Step 1: Orient — read the backend ticket and the CMS contract

I started by reading the existing `FAKE-CMS` ticket and the repo to understand
what the frontend must consume. The backend is a read-only, SQLite-backed
GraphQL mock of a WordPress + Yoast media CMS, built to drive an SSG workshop.
Its `internal/doc/workshop-ssg.md` already specifies the deliverable: a static
site that mirrors five URL conventions and renders seven block types. That
document is the acceptance bar for everything in this ticket.

### Prompt Context

**User prompt (verbatim):** "Create a new docmgr ticket to build a frontend for this cms using 11ty. Create a detailed analysis / design / implementation guide that is for a new intern, explaining all the parts of the system needed to understand what it is, with prose paragraphs and bullet points and pseudocode and diagrams and api references and file references. It should be very clear and technical. Store in the ticket and then upload to remarkable. Research how to create a 11ty plugin for that, use defuddle to download pages that are relevant and store them in the sources/ directory of the ticket."

**Assistant interpretation:** Stand up a new docmgr ticket for the 11ty frontend, write a deep intern-facing design/implementation guide backed by real 11ty docs (fetched with defuddle into sources/), keep full docmgr bookkeeping, validate, and ship to reMarkable.

**Inferred user intent:** Reproduce the proven structure of the `FAKE-CMS` ticket (design doc + diary + sources + reMarkable) but for the *frontend half* of the workshop, so a new intern can build the SSG plugin from a self-contained, evidence-backed guide.

### What I did
- Read `README.md`, `schema.graphql`, `internal/doc/{api-reference,workshop-ssg}.md`, `internal/server/server.go`, `internal/domain/domain.go`.
- Read the backend ticket's `index.md`, `tasks.md`, `changelog.md`, and the structure of its `design-doc/01-…md` to match style/depth.
- Confirmed the workshop-ssg contract is the acceptance bar (URL conventions + block types).

### Why
A design guide is only credible if its claims about the CMS are anchored in the actual schema and the actual workshop contract. Starting from the backend ticket avoided re-deriving the requirements.

### What worked
- The backend ticket already pins the five URL conventions and seven block types — I could reference them verbatim instead of re-scraping `20minutes-media.com`.
- The backend's "one-request full render" query shape (`api-reference.md`) mapped directly onto the plugin's full-render fragment.

### What didn't work
- Nothing yet; orientation was clean.

### What I learned
- The pedagogically important constraint is that **the CMS never returns HTML** — the body is a typed block union, so the frontend *must* implement a renderer per variant. That single fact shapes the whole design (§11 of the guide).
- The backend already guards against server-side N+1; the frontend must not reintroduce it client-side by splitting the per-article query.

### What was tricky to build
- N/A for this step (research only).

### What warrants a second pair of eyes
- The mapping of `PostType` enum → URL slug (§4.3 of the guide) is load-bearing; a wrong transform puts every article at the wrong URL. Worth double-checking against the real site's paths.

### What should be done in the future
- Implement the plugin per the phased plan (guide §18) and wire the six acceptance tests (guide §17.2).

### Code review instructions
- Start at `schema.graphql` (the contract), then `internal/doc/workshop-ssg.md` (acceptance).
- Validate: `./fake-cms serve` then run any example query from `api-reference.md` in `/playground`.

### Technical details
- CMS endpoint: `POST http://localhost:8080/graphql`; playground `/playground`; health `/healthz`.

## Step 2: Create the ticket and download 11ty docs as evidence

I created the `FAKE-CMS-11TY` ticket and, rather than trust my memory of 11ty's
API, downloaded the official Eleventy docs into `sources/` with `defuddle`. This
mirrors how the backend ticket used scraped evidence, and it means every API
claim in the design doc (plugin signature, data-cascade precedence, pagination
semantics, programmatic API) is backed by a committed snapshot.

### Prompt Context

**User prompt (verbatim):** (see Step 1)

### What I did
- `docmgr ticket create-ticket --ticket FAKE-CMS-11TY …` → created the workspace.
- Ran `defuddle parse <url> --md | fold -w 100 -s` for 9 official 11ty pages:
  plugins, create-plugin, custom templates, data-cascade, pagination,
  collections, JS data files, programmatic API, config.
- Prepended consistent YAML frontmatter (with `external-url` + re-run command)
  to each source so `docmgr doctor` stays clean and sources are refreshable.
- Wrote `sources/00-sources-readme.md` as a manifest mapping each file to the
  design-doc section that cites it.
- Added vocabulary slugs (`11ty`, `ssg`, `frontend`, `javascript`).

### Why
A design doc that says "use `addGlobalData`" is much stronger when the reader
can open `sources/18-11ty-config.md` and see the exact signature. Snapshots also
make the doc reproducible offline (important for a workshop laptop).

### What worked
- `defuddle` + `fold -w 100 -s` produced clean, re-wrapped markdown for all nine
  pages (the defuddle skill documents the single-line-output quirk; `fold` fixes it).
- Kagi confirmed the canonical URLs (`11ty.dev/docs/...`) before downloading.

### What didn't work
- Each defuddle output starts with ~110 lines of nav chrome. Not an error — the
  real content follows. I accounted for this when extracting API details.

### What I learned
- **As of 11ty v3, `addPlugin` is async-aware** and must be `await`ed if the
  plugin function itself is async (`sources/11-11ty-create-plugin.md`). This
  pushed me toward keeping the plugin function *synchronous* and putting all I/O
  inside `addGlobalData`'s async resolver (guide §9.1, Decision 16.2).
- The **data-cascade precedence ladder** (`sources/13-…`) is what justifies
  injecting fetched content at rung 2 (`addGlobalData`) — global, but any
  template can override.

### What was tricky to build
- Choosing *where* the fetched CMS data enters the cascade. The choice
  (`addGlobalData`, rung 2) has a concrete consequence: per-template fetching
  (the tempting alternative) reintroduces client-side N+1. Documented as
  Decision 16.2 so an intern doesn't "fix" it the wrong way.

### What warrants a second pair of eyes
- Confirm the `defuddle` snapshots are current enough for the 11ty version the
  intern will install (v3). If they target v3 explicitly, the async-`addPlugin`
  note is essential.

### What should be done in the future
- Refresh `sources/` before publishing if 11ty releases breaking changes; the
  re-run command is in each file's frontmatter.

### Code review instructions
- Skim `sources/00-sources-readme.md` (manifest), then spot-check
  `sources/11-11ty-create-plugin.md` and `sources/14-11ty-pagination.md` — the
  two most-cited pages.

### Technical details
- Re-download any source:
  `defuddle parse https://www.11ty.dev/docs/create-plugin/ --md | fold -w 100 -s > sources/11-11ty-create-plugin.md`

## Step 3: Write the design & implementation guide

I wrote
`design-doc/01-11ty-frontend-design-implementation-guide.md` (21 sections) for a
new intern. It opens with a glossary and an 11ty primer, establishes the current
state from the backend's schema, then designs the plugin, the templates, the
block renderer, and the test/acceptance plan — with prose, bullet lists,
pseudocode, Mermaid diagrams, and concrete JS.

### Prompt Context

**User prompt (verbatim):** (see Step 1)

### What I did
- Authored the guide with the same depth/structure as the backend ticket's
  design doc (executive summary → problem/scope → glossary → current state → gap
  → primer → architecture → file-by-file → block renderer → decisions → tests →
  phased plan → risks → references → TL;DR).
- Embedded two Mermaid diagrams (component flow + build sequence).
- Specified `renderBlocks` as a literal `switch` with throw-on-unknown (Decision
  16.4) and the block→markup table as the spec.
- Lifted the six acceptance criteria from `workshop-ssg.md` into §17.2 as tests.

### Why
The user asked for a guide "for a new intern, explaining all the parts of the
system … very clear and technical." That means no assumed 11ty knowledge and no
ungrounded claims. Every API is cited to `sources/`; every CMS claim to the
backend's `schema.graphql`.

### What worked
- The "one-request full render" fragment (§9.3) reused the backend's documented
  query shape verbatim — no invention.
- The pagination-with-`size:1` technique (§6.4/§13.1) cleanly answers "how do
  hundreds of article pages come from one template file," which is the heart of
  the integration.

### What didn't work
- Nothing; drafting was linear once the evidence was in place.

### What I learned
- The cleanest plugin shape is **Option C** (adapter plugin + reference site
  templates): the package owns the un-fun plumbing (fetch, URLs, block
  renderer), the site owns the templates, preserving the workshop's pedagogy
  (guide §7, Decision 16.1).

### What was tricky to build
- Making `cmsPermalink` correct for *all five* entity kinds while keeping it one
  function. Resolved by switching on `__kind` and centralizing the
  `/rubrique/`-not-`/tag/` rule there (Decision 16.5) so a template can't get it
  wrong.

### What warrants a second pair of eyes
- §13.2 (`registerCollections`) reads `cms` off `collectionApi.getAll()[0]?.data`
  — a known-fragile 11ty pattern for reaching global data from a collection
  function. Prefer computing groups inside `fetchAll` (add `cms.articlesByType`)
  before relying on this in production.
- The block renderer's `EmbedBlock` provider map is stubbed (YouTube only);
  extend before claiming "distinct markup" for arbitrary providers.

### What should be done in the future
- Implement Phases 0–5 (guide §18); wire the six acceptance tests; add opt-in
  `@11ty/eleventy-fetch` caching (Decision 16.6).

### Code review instructions
- Read the guide top-to-bottom. Then sanity-check the two most load-bearing code
  blocks: `fetchAllArticles` (§9.2) and `renderOne` (§11).
- Validate by hand: against `fake-cms serve`, run the full-render query from
  §10.3 in `/playground` and confirm the fragment parses.

### Technical details
- Guide size: ~53 KB / 21 sections. Diagrams: Mermaid `flowchart LR` (components)
  and `sequenceDiagram` (build lifecycle). API surface cited: `addPlugin`,
  `addGlobalData`, `addFilter`, `addShortcode`, `addCollection`, `addExtension`,
  pagination front matter, programmatic `new Eleventy()`.

## Step 4: Review the first 11ty guide against the executable server and write a corrected guide

This step treated the first 11ty design guide as a draft from a well-meaning but too-trusting implementer. The main change was methodological: instead of coding against the checked-in SDL and prose docs alone, I started the actual `fake-cms` server, queried its executable schema, tested the central Eleventy pattern in a temporary project, and then wrote a separate corrected guide that targets the behavior that really runs today.

The corrected guide keeps the good architecture — Eleventy global data, pagination, filters, shortcodes, and a small plugin boundary — but changes the implementation plan where the first guide was wrong. In particular, it removes invalid GraphQL queries (`pages`, `site`, taxonomy filters, direct `order` on a union), makes page rendering optional until pages are enumerable, and recommends visible templates plus a project-local plugin instead of a black-box package.

### Prompt Context

**User prompt (verbatim):** "Go over the guide for building a 11ty plugin / solution to convert our fake cms graphql server data into a static site. It was built by your little brother who is not necessarily the smartest, so you should definitely doublecheck things, and then write your own design and implementation guide in a separate documentation and uploading it to remarkable. Explain things to your brother and an intern, especially a whole section about 11ty itself, and how the flow works and what plugins are (and if plugins are even the right way to build this)."

**Assistant interpretation:** Review the existing 11ty design guide critically, verify its claims against the actual repo/server and Eleventy APIs, write a separate corrected guide for both the original author and an intern, and upload that corrected document to reMarkable.

**Inferred user intent:** The user wants workshop-ready guidance that is accurate, pedagogical, and not merely a polished version of an unverified draft. They specifically want the corrected document to explain 11ty itself and make an explicit decision about whether a plugin is the right abstraction.

### What I did
- Read the existing guide `design-doc/01-11ty-frontend-design-implementation-guide.md`.
- Re-read the backend contract files: `schema.graphql`, `internal/doc/api-reference.md`, `internal/doc/workshop-ssg.md`, and `internal/graphql/schema.go`.
- Installed/tested Eleventy 3.1.6 in temporary directories to verify:
  - async `addGlobalData` can feed `pagination.data`,
  - filters work inside dynamic permalinks,
  - `eleventyConfig.addTemplate` can create virtual paginated pages.
- Downloaded additional official Eleventy docs with defuddle:
  - `sources/19-11ty-config-global-data.md`
  - `sources/20-11ty-fetch.md`
  - `sources/21-11ty-virtual-templates.md`
- Started `./fake-cms serve --path testdata/cms.db --addr :18080` and captured a live schema/query evidence file at `sources/22-live-schema-check.md`.
- Wrote `design-doc/02-corrected-11ty-cms-static-site-guide.md` as a separate second-pass guide.

### Why
- The first guide made several claims that were plausible from the checked-in SDL but false against the executable GraphQL schema.
- The workshop guide should teach the intern to verify executable contracts, not just read docs.
- The plugin decision needed nuance: plugins are right for shared adapter behavior, but not for hiding all page templates from learners.

### What worked
- A temporary Eleventy 3.1.6 project proved the key pattern: `addGlobalData("cms", async () => …)` can feed a paginated template, and the dynamic permalink wrote `_site/actualites/.../index.html` and `_site/best-cases/.../index.html` correctly.
- `eleventyConfig.addTemplate` also worked for virtual paginated templates, confirming that `addTemplate` — not `addExtension` — is the right API if a plugin later contributes pages.
- The seeded backend (`--path testdata/cms.db`) returned `totalCount 140` and a valid renderable article when queried with the corrected block-union fragment.

### What didn't work
- Running `./fake-cms serve` without `--path testdata/cms.db` used the default `cms.db`, which was empty in this working tree:
  - `articles(first:3) { totalCount }` returned `0`.
  - `./fake-cms db query --query "select kind,status,count(*) from content_node group by kind,status"` returned `(0 rows)`.
- The first guide's union query shape failed exactly as GraphQL should fail:
  - Query: `blocks { __typename order ... on ParagraphBlock { text } }`
  - Error: `Cannot query field "order" on type "BlockUnion". Did you mean to use an inline fragment on "Block", "EmbedBlock", "GalleryBlock", "HeadingBlock", or "ImageBlock"?`
- The first guide's executable-schema assumptions failed:
  - `Cannot query field "site" on type "Query".`
  - `Cannot query field "pages" on type "Query". Did you mean "page" or "tags"?`
  - `Unknown argument "first" on field "categories" of type "Query".`
- The executable schema returns `seo.og` and `seo.twitter` as Go `fmt` map strings, not structured objects; only `seo.jsonLd` is a reliable structured JSON value today.

### What I learned
- The checked-in SDL and docs are aspirational in places; the running `internal/graphql/schema.go` is the contract the frontend must survive today.
- The best workshop architecture is not “publish a black-box plugin first.” It is a project-local plugin for fetch/normalization/helpers plus visible templates in `src/` so the intern learns the page-generation flow.
- Eleventy 3's `addTemplate` gives a clean later path for virtual templates if a one-command demo becomes useful.

### What was tricky to build
- The trickiest part was separating good high-level architecture from bad concrete query shapes. The first guide's overall idea survived, but several code snippets would fail immediately against the server.
- Another sharp edge is the backend contract mismatch: `schema.graphql` says `site`, `pages`, typed SEO objects, and taxonomy filters exist, while the executable schema omits or weakens them. The corrected guide therefore documents workarounds and recommends backend fixes before a polished workshop.

### What warrants a second pair of eyes
- Confirm whether the frontend should target the executable schema as-is or whether we should first fix the backend to match `schema.graphql`.
- Review the recommendation to keep templates visible rather than shipping all pages as plugin virtual templates; this is a pedagogy decision as much as a technical one.
- Review `seo.og/twitter/breadcrumbs` handling. The frontend guide currently treats only `jsonLd` as reliable because the executable schema exposes the others as strings.

### What should be done in the future
- Either update the backend executable schema to match `schema.graphql` or update the SDL/docs to match reality.
- If the backend remains unchanged, implement the frontend using the corrected guide's workarounds: local `site` config, optional explicit `pageSlugs`, client-side taxonomy grouping, and a corrected block-union fragment.
- Add a contract test that starts `fake-cms --path testdata/cms.db` and verifies the frontend's GraphQL fragment before any Eleventy build runs.

### Code review instructions
- Start with `design-doc/02-corrected-11ty-cms-static-site-guide.md`, especially sections 1, 3, 4, and 5.
- Compare its corrections against `sources/22-live-schema-check.md` and `internal/graphql/schema.go`.
- Validate the central Eleventy claim by repeating the temporary-project test: async `addGlobalData` + `pagination.data: cms.articles` + dynamic permalink.

### Technical details
- Eleventy version tested: `@11ty/eleventy@3.1.6`.
- Backend smoke command: `./fake-cms serve --path testdata/cms.db --addr :18080`.
- Correct block query pattern: `blocks { __typename ... on Block { id order } ... on ParagraphBlock { text align } ... }`.
- Correct current list queries: `categories { ... }`, `tags { ... }`, `authors { ... }` with no `first` arguments.

## Step 5: Commit the documentation baseline and expand implementation tasks

This step turned the research package into a clean implementation starting point. I first committed the ticket documentation and sources so the design baseline is recoverable, then replaced the coarse task list with a phase-by-phase checklist that can drive small implementation commits.

The task list now separates scaffolding, GraphQL fetching, normalization, block rendering, plugin wiring, templates, integration tests, and acceptance checks. That matters because the frontend has several independent risk areas: the executable GraphQL schema differs from the aspirational docs, Eleventy pagination is a build-time mechanism, and block rendering needs its own pure tests.

### Prompt Context

**User prompt (verbatim):** "Commit at appropriate intervals, and now commit all the ticket work and documents. Then add detailed tasks to the ticket, and then implement it task by task, keeping a detailed diary as you work.

If the site doesnot contain enough data, feel free to seed more from the 20minutes-media.com research we did."

**Assistant interpretation:** Commit the existing ticket docs first, make the implementation backlog detailed enough to work from, then implement the 11ty frontend in focused commits while updating the diary after each meaningful step.

**Inferred user intent:** The user wants disciplined implementation: no giant unreviewable change, no untracked documentation, and a diary that makes the path understandable for review and continuation.

### What I did
- Read the git-commit hygiene skill.
- Observed an untracked built `fake-cms` binary and added `/fake-cms` to `.gitignore`.
- Committed `.gitignore` as `010c602` — `chore: ignore local fake-cms binary`.
- Staged and committed the FAKE-CMS-11TY ticket docs, sources, diary, changelog, and vocabulary as `33c6fa7` — `docs: add FAKE-CMS 11ty frontend design ticket`.
- Rewrote `tasks.md` into a detailed implementation checklist with phases P0–P7 and acceptance gates A1–A6.

### Why
- The existing ticket work was valuable and should be committed before code changes.
- The untracked `fake-cms` binary is a build artifact and should not be accidentally staged later.
- A detailed task list lets implementation proceed in reviewable increments and keeps the diary aligned with the work.

### What worked
- The docs commit was focused and included only the ticket workspace and vocabulary.
- The task list now maps directly to the corrected guide's implementation phases.

### What didn't work
- No implementation failures yet. This step was planning and commit hygiene.

### What I learned
- The repo did not previously ignore the root `fake-cms` build binary, even though `make build` creates it.

### What was tricky to build
- Keeping the docs baseline separate from implementation was the main ordering constraint. If code and docs were mixed, review would be harder and the implementation diary would lose a clean starting point.

### What warrants a second pair of eyes
- The task list assumes the frontend should target the current executable schema rather than fixing the backend first. Reviewers may decide to reorder work and align the backend schema before building the frontend.

### What should be done in the future
- Implement P0 next: create the `frontend/` skeleton, prove Eleventy pagination with a tiny local smoke path, and commit that separately.

### Code review instructions
- Review `.gitignore` for the binary ignore rule.
- Review `tasks.md` to ensure the implementation phases and acceptance gates match the corrected guide.
- Validate with `git log --oneline -3` and `git status --porcelain`.

### Technical details
- Code/doc commits:
  - `010c602` — `chore: ignore local fake-cms binary`
  - `33c6fa7` — `docs: add FAKE-CMS 11ty frontend design ticket`

## Step 6: Scaffold the Eleventy frontend and prove pagination with hardcoded data

This step created the first real frontend files under `frontend/`. The goal was intentionally small: prove the central Eleventy mechanism before adding GraphQL. A hardcoded `cms.articles` global-data object feeds `src/articles.njk`, and Eleventy pagination emits a static article page at the CMS-style path `/best-cases/hardcoded-smoke-article/`.

This unlocks the rest of the implementation because the static-site skeleton, package scripts, layout, passthrough CSS, and output directory are now real. The next step can replace the hardcoded data with a GraphQL client without also debugging Eleventy project setup.

**Commit (code):** <pending> — "feat(frontend): scaffold Eleventy site"

### What I did
- Added frontend ignore rules to `.gitignore`: `/frontend/_site/`, `/frontend/.cache/`, `/frontend/node_modules/`.
- Created `frontend/package.json` with scripts for `build`, `serve`, `clean`, and `test`.
- Installed `@11ty/eleventy` and generated `frontend/package-lock.json`.
- Added `frontend/eleventy.config.cjs` with hardcoded `cms` global data, a temporary `json` filter, and a temporary `renderBlocks` shortcode.
- Added visible templates:
  - `frontend/src/index.njk`
  - `frontend/src/articles.njk`
  - `frontend/src/_includes/base.njk`
  - `frontend/src/styles.css`
- Added `frontend/README.md` with the seeded fake-cms command.
- Ran `npm run clean` and `npm run build` from `frontend/`.
- Marked P0.1, P0.2, and P0.3 complete in `tasks.md`.

### Why
- The corrected design depends on Eleventy pagination over build-time data. Proving that with one hardcoded article isolates Eleventy setup from GraphQL/API work.
- Keeping the templates visible now preserves the workshop teaching goal: an intern can open the files and see how one template generates many static pages.

### What worked
- `npm run build` succeeded with Eleventy 3.1.6.
- Eleventy wrote:
  - `_site/best-cases/hardcoded-smoke-article/index.html`
  - `_site/index.html`
- Passthrough CSS copied successfully.

### What didn't work
- No failure in this step. I deliberately did not run `npm test` yet because the test directory is not implemented until later phases.

### What I learned
- The corrected guide's core assumption is valid in this repo, not just in a throwaway test: `addGlobalData` plus `pagination.data` produces the desired CMS-style path.

### What was tricky to build
- The temporary `renderBlocks` shortcode and hardcoded data are intentionally disposable. The tricky part is preventing them from becoming permanent architecture. They are only a scaffold so the next commits can replace them with `_config/renderBlocks.cjs` and `_config/fakeCmsPlugin.cjs`.

### What warrants a second pair of eyes
- Confirm the frontend should live under `frontend/` rather than at the repo root. This keeps the Go backend untouched but introduces a nested npm project.

### What should be done in the future
- Replace hardcoded global data with the project-local plugin and GraphQL client in P1/P4.
- Add tests before relying on `npm test` in CI.

### Code review instructions
- Start with `frontend/eleventy.config.cjs`, then `frontend/src/articles.njk`.
- Validate with:
  - `cd frontend && npm install`
  - `npm run clean && npm run build`
  - `find _site -type f | sort`

### Technical details
- Build command: `npm run build`
- Output observed: `[11ty] Copied 2 Wrote 2 files in 0.05 seconds (v3.1.6)`

## Step 7: Add the GraphQL client, contract smoke, and CMS normalization

This step replaced guesswork about the backend with executable frontend code. I added a GraphQL client that follows the current fake-cms schema, including the correct block-union fragment, and a contract smoke script that proves the seeded server returns articles while deliberately catching the invalid `blocks { order }` query from the first guide.

I also added normalization helpers that make templates simpler and safer: URL paths are derived once, tag URLs are locked to `/rubrique/`, accented slugs are URL-encoded, article indexes are built once, and sitemap URLs are assembled from normalized entities.

**Commit (code):** <pending> — "feat(frontend): add CMS client and normalization"

### What I did
- Added `frontend/_config/fakeCmsClient.cjs`:
  - `gql(endpoint, query, variables)`
  - `fetchAllArticles()` with cursor pagination
  - `fetchKnownPages()` for optional configured page slugs
  - `fetchCms()` for articles/categories/tags/authors/pages
- Added `frontend/_config/normalizeCms.cjs`:
  - `postTypeSlug`
  - `pathSegment`
  - `normalizeCms`
  - taxonomy/author/post-type indexes
  - `sitemapUrls`
- Added `frontend/scripts/contract-smoke.mjs`.
- Added `frontend/test/normalizeCms.test.mjs`.
- Added `contract:smoke` npm script.
- Ran `npm test` successfully.
- Started the seeded backend with `./fake-cms serve --path testdata/cms.db --addr :18080` and ran `CMS_ENDPOINT=http://localhost:18080/graphql npm run contract:smoke` successfully.
- Marked P1.1–P1.4 and P2.1–P2.4 complete.

### Why
- The frontend must target the executable GraphQL schema, not the aspirational SDL.
- URL derivation and grouping are correctness-sensitive enough to live in testable JavaScript helpers rather than ad hoc template snippets.

### What worked
- `npm test` passed 3 normalization tests.
- Contract smoke reported: `fake-cms contract smoke ok: 140 articles at http://localhost:18080/graphql`.
- The contract smoke confirms both the positive query and the expected failure for direct `BlockUnion.order` selection.

### What didn't work
- No new failures. The known backend mismatch is handled by querying `categories`, `tags`, and `authors` without `first` and by not relying on `site` or `pages` enumeration.

### What I learned
- The seeded dataset has enough content for the workshop right now: 140 articles. No extra 20minutes-media.com seeding is needed at this point.

### What was tricky to build
- The current executable schema exposes `articles(filter: String)` but not the typed filter object from `schema.graphql`. The client avoids that path and fetches all articles once, then normalization builds the taxonomy indexes client-side.
- `pathSegment` deliberately uses `encodeURIComponent`; this preserves output correctness for accented slugs, but it means generated file paths contain URL-encoded segments.

### What warrants a second pair of eyes
- Confirm whether encoded path segments are preferred over raw UTF-8 output directories for the final workshop site. The backend troubleshooting text says to handle accents; URL encoding is safe, but raw UTF-8 paths may look nicer.
- Review whether `fetchAllArticles` should assert `articles.length === totalCount` directly. The integration test will cover this later.

### What should be done in the future
- P3: implement the real block renderer with seven branches and tests.
- Later P4: wire `fetchCms + normalizeCms` into the project-local plugin, replacing the hardcoded smoke data.

### Code review instructions
- Start with `frontend/_config/fakeCmsClient.cjs`, especially `ARTICLE_FIELDS` and `fetchAllArticles`.
- Then review `frontend/_config/normalizeCms.cjs` for URL conventions and indexes.
- Validate with:
  - `cd frontend && npm test`
  - `./fake-cms serve --path testdata/cms.db --addr :18080`
  - `CMS_ENDPOINT=http://localhost:18080/graphql npm run contract:smoke`

### Technical details
- Correct current list queries: `categories { ... }`, `tags { ... }`, `authors { ... }`.
- Correct block union query pattern: `blocks { __typename ... on Block { id order } ... }`.

## Step 8: Implement the seven-variant CMS block renderer

This step implemented the real block renderer that the workshop contract requires. The temporary renderer from the scaffold only handled paragraphs; the new `renderBlocks` module renders all seven current CMS block variants, escapes text and attributes, sorts by `order`, clamps heading levels, and throws when the schema introduces an unknown block type.

The block renderer is deliberately pure JavaScript with unit tests. It does not fetch data and does not depend on Eleventy. That separation matters because block rendering is the content-correctness core of the exercise and should fail fast before a full static build is involved.

**Commit (code):** <pending> — "feat(frontend): render CMS blocks"

### What I did
- Added `frontend/_config/renderBlocks.cjs` with:
  - `escapeHtml`
  - `attrs`
  - `clampHeadingLevel`
  - `renderBlocks`
  - `renderBlock`
- Implemented render branches for:
  - `ParagraphBlock`
  - `HeadingBlock`
  - `ImageBlock`
  - `ListBlock`
  - `QuoteBlock`
  - `EmbedBlock`
  - `GalleryBlock`
- Added `frontend/test/renderBlocks.test.mjs`.
- Ran `npm test`; all 8 tests passed.
- Marked P3.1–P3.4 complete.

### Why
- The CMS body is a GraphQL union, not raw HTML. The frontend must explicitly render each variant.
- Unknown blocks should break the build rather than silently dropping content.

### What worked
- The renderer produces distinct semantic markup/classes for all seven block types.
- The tests confirm order sorting and unknown-block failure.
- HTML escaping is centralized and covered by tests.

### What didn't work
- No failures in this step.

### What I learned
- `ImageBlock` and `EmbedBlock` both naturally render as `<figure>`, so distinctness is better asserted through classes (`block-image`, `block-embed-*`) rather than only root tag names.

### What was tricky to build
- Heading levels need clamping. CMS data is synthetic and currently sane, but a renderer should not emit `<h99>` or `<h1>` accidentally inside article bodies.
- Embed rendering is intentionally conservative: it outputs a linked figure instead of provider-specific iframes. This keeps the first implementation safe and distinct; richer provider rendering can come later.

### What warrants a second pair of eyes
- Review whether `EmbedBlock` should render YouTube iframes now or remain a safe linked embed for the workshop baseline.
- Review escaping coverage before allowing any CMS-provided text into `| safe` template output.

### What should be done in the future
- P4 should wire this renderer into the real plugin shortcode, replacing the scaffold's temporary inline shortcode.

### Code review instructions
- Start with `frontend/_config/renderBlocks.cjs`.
- Validate with `cd frontend && npm test`.
- Inspect each branch and confirm it maps to the fields in `schema.graphql` / the live schema.

### Technical details
- Unknown block error text: `Unknown CMS block type: <typename>`.
- Current tests: 8 total across normalization and block rendering.

## Step 9: Wire the project-local plugin to real fake-cms data

This step replaced the scaffold's hardcoded `cms` object with a real project-local Eleventy plugin. The plugin now fetches the seeded GraphQL API, normalizes the result, registers URL/JSON filters, and exposes the tested block renderer as a shortcode. With the seeded backend running, `npm run build` generated the homepage plus all 140 article pages from real CMS data.

This is the first end-to-end static build from the fake CMS. It does not yet render tag/category/author archive pages or a sitemap, but the central data path is complete: fake-cms GraphQL → fetch client → normalizer → Eleventy data cascade → paginated article template → `_site/` files.

**Commit (code):** <pending> — "feat(frontend): wire Eleventy plugin to CMS data"

### What I did
- Added `frontend/_config/fakeCmsPlugin.cjs`.
- Rewrote `frontend/eleventy.config.cjs` to register the project-local plugin instead of hardcoded data.
- Registered filters/shortcodes:
  - `postTypeSlug`
  - `pathSegment`
  - `json`
  - `absoluteUrl`
  - `renderBlocks`
- Updated `frontend/README.md` to document real CMS build commands and backend contract notes.
- Ran `npm test` successfully.
- Started `./fake-cms serve --path testdata/cms.db --addr :18080` and ran `CMS_ENDPOINT=http://localhost:18080/graphql npm run build` successfully.
- Marked P4.1–P4.3 complete.

### Why
- The plugin is the boundary recommended by the corrected guide: shared data fetching and helper behavior live in `_config`, while page templates remain visible in `src/`.
- Wiring the real API before adding archive templates reduces risk; article generation is the core path and now works independently.

### What worked
- The build generated 141 Eleventy pages (homepage + 140 articles) and copied CSS.
- The output included encoded accented slugs such as `reportage-engag%C3%A9--oenobiol-72`.
- The seeded dataset is sufficient for development; no extra seeding needed.

### What didn't work
- No new failures. The build requires a running CMS server, as expected.

### What I learned
- Eleventy build output is verbose when writing 140 pages; this is useful during development but may warrant `--quiet` in scripted integration tests.

### What was tricky to build
- The plugin must remain synchronous at registration time while using async global data. The correct pattern is `eleventyConfig.addGlobalData("cms", async () => { ... })`, not an async plugin function.
- The build is now coupled to the backend process. Integration tests need to start the backend explicitly or skip with a clear message.

### What warrants a second pair of eyes
- Review whether the plugin should support an offline fixture mode for workshops where the backend is not running. I did not add this because acceptance requires building from `fake-cms serve`.

### What should be done in the future
- P5 should add visible archive, author, homepage, and sitemap templates using the normalized indexes that already exist.

### Code review instructions
- Start with `frontend/_config/fakeCmsPlugin.cjs`, then `frontend/eleventy.config.cjs`.
- Validate with:
  - `cd frontend && npm test`
  - `./fake-cms serve --path testdata/cms.db --addr :18080`
  - `CMS_ENDPOINT=http://localhost:18080/graphql npm run build`

### Technical details
- Build output: `Copied 2 Wrote 141 files in 0.38 seconds (v3.1.6)`.
- File count after build: 142 files under `frontend/_site` including copied CSS.

## Step 10: Add visible archive, author, page, head, and sitemap templates

This step turned the article-only build into a fuller static site. The frontend now renders post-type archive pages, tag pages under `/rubrique/`, category archives under `/archives/`, author pages, an optional page template, a homepage, and `sitemap.xml`. The shared `<head>` partial now serializes `seo.jsonLd` into the document head rather than leaving JSON-LD inside the body.

The implementation still keeps the teaching surface visible: all page families are normal Nunjucks templates under `frontend/src/`, not hidden virtual templates. The project-local plugin supplies data and helpers; the templates show how static pages are generated from normalized CMS data.

**Commit (code):** <pending> — "feat(frontend): add archive templates and sitemap"

### What I did
- Added `frontend/src/_includes/head.njk`.
- Rewrote `frontend/src/_includes/base.njk` to include the head partial and section navigation.
- Added `frontend/src/_includes/article-card.njk`.
- Rewrote `frontend/src/articles.njk` to use real CMS fields, taxonomy links, featured media, and the shared head partial.
- Rewrote `frontend/src/index.njk` as a homepage over normalized CMS data.
- Added archive templates:
  - `frontend/src/tag.njk`
  - `frontend/src/category.njk`
  - `frontend/src/author.njk`
  - `frontend/src/post-type.njk`
  - `frontend/src/pages.njk`
  - `frontend/src/sitemap.xml.njk`
- Expanded `frontend/src/styles.css` for cards, layout, taxonomy links, and block rendering.
- Ran `npm test` successfully.
- Built against `./fake-cms serve --path testdata/cms.db --addr :18080` successfully.
- Spot-checked that JSON-LD appears before `<body>` and no `/tag/` directory exists.
- Marked P5.1–P5.4 complete.

### Why
- Acceptance requires more than article pages: tags must live at `/rubrique/`, sitemap must cover generated URLs, and SEO JSON-LD must be injected into page documents.
- Visible templates are part of the workshop value; they show the intern how Eleventy pagination creates each page family.

### What worked
- Build output: `Copied 2 Wrote 189 files in 0.60 seconds (v3.1.6)`.
- There are 30 generated `/rubrique/<slug>/` tag pages.
- JSON-LD appears in the article `<head>` before `<body>`.
- No `/tag/` directory was generated.

### What didn't work
- No failures. The optional `pages.njk` currently emits no pages unless `CMS_PAGE_SLUGS` is configured, because the backend still lacks page enumeration.

### What I learned
- Nunjucks layouts can see the pagination alias (`article`, `tag`, `category`, etc.), so the shared head partial can derive article/page SEO without trying to serialize object-valued `eleventyComputed` data.

### What was tricky to build
- The SEO object cannot be passed around as a string. The head partial detects `article.seo` or `cmsPage.seo` directly from the rendering context and serializes `jsonLd` with the `json` filter.
- Sitemap generation needs both normalized CMS entity URLs and local generated pages such as `/` and post-type archive pages.

### What warrants a second pair of eyes
- Review the exact sitemap scope. It currently includes homepage, post-type archives, normalized articles/categories/tags/authors/pages. It does not include pagination subpages because archives are not paginated beyond one page each.
- Review whether encoded slugs in generated paths are acceptable for the visual workshop output.

### What should be done in the future
- P6 should automate the spot checks: article count, no `/tag/`, `/rubrique/` exists, sitemap includes expected URLs, and JSON-LD parses.

### Code review instructions
- Start with `frontend/src/_includes/head.njk`, then `frontend/src/tag.njk`, `frontend/src/category.njk`, and `frontend/src/sitemap.xml.njk`.
- Validate with a running seeded CMS:
  - `cd frontend && npm test`
  - `CMS_ENDPOINT=http://localhost:18080/graphql SITE_URL=http://example.test npm run build`
  - `grep -R "application/ld+json" _site/actualites _site/best-cases | head`

### Technical details
- Build file count after P5: 190 files under `_site` including copied CSS.
- Generated tag-page count: 30.

## Step 11: Add integration build acceptance checks

This step automated the end-to-end checks that were previously manual spot checks. The new integration script starts the seeded Go backend with `go run`, waits for `/healthz`, builds the Eleventy frontend against that live GraphQL endpoint, and inspects `_site/` for the workshop acceptance criteria.

The script now verifies the important invariants directly: the generated article-page count equals the CMS `totalCount`, no `/tag/` directory exists, `/rubrique/` pages exist, every generated HTML page is represented in `sitemap.xml`, and a representative article contains parseable JSON-LD.

**Commit (code):** <pending> — "test(frontend): add integration build acceptance checks"

### What I did
- Added `frontend/scripts/integration-build.mjs`.
- Added `test:integration` npm script.
- The integration script:
  - starts `go run ./cmd/fake-cms serve --path testdata/cms.db --addr :18081`,
  - waits for `/healthz`,
  - queries article `totalCount`,
  - runs the Eleventy build with `CMS_ENDPOINT` and `SITE_URL`,
  - counts article pages,
  - asserts no `/tag/` output,
  - asserts `/rubrique/` output exists,
  - verifies sitemap coverage for generated HTML pages,
  - parses JSON-LD from a representative article page.
- Fixed process cleanup by spawning `go run` as a detached process group and killing the process group in `finally`.
- Ran `npm test` and `npm run test:integration` successfully.
- Marked P6.1–P6.5 and acceptance gates A1–A6 complete.

### Why
- The frontend's correctness depends on the interaction between Go backend, GraphQL client, Eleventy global data, templates, and generated files. Unit tests cannot prove the full path.
- The previous script version proved the checks but timed out because `go run` left a child process alive; reliable cleanup is part of making the integration test usable.

### What worked
- Integration output: `integration ok: 140 article pages, 190 files`.
- The script catches the highest-risk regressions: article count mismatch, `/tag/` URL regression, sitemap drift, and missing/invalid JSON-LD.

### What didn't work
- First integration run printed `integration ok` but the command timed out after 180 seconds. Cause: killing the parent `go run` process did not terminate the compiled child server process.
- Fix: spawn with `detached: true` and call `process.kill(-server.pid, "SIGTERM")` to terminate the process group.

### What I learned
- `go run` process management matters in Node integration tests. Treat it as a process group, not a single process.
- Verifying sitemap coverage by comparing generated HTML paths against sitemap contents is stronger than checking only a few sample URLs.

### What was tricky to build
- Counting article pages required distinguishing `/post-type/<slug>/index.html` article pages from `/post-type/index.html` archive pages and from `/archives`, `/rubrique`, and `/author` pages.
- The script must be robust enough to run from `frontend/` while starting the Go backend from the repository root.

### What warrants a second pair of eyes
- Review the integration script's article-file classifier. It assumes article pages have exactly `/<postType>/<slug>/index.html` shape and excludes `archives`, `rubrique`, and `author` prefixes.
- Review whether the integration test should be part of default `npm test` or stay as explicit `npm run test:integration` because it starts a Go server.

### What should be done in the future
- P7 should update the frontend README with test commands and run final `go test ./...` plus `docmgr doctor`.

### Code review instructions
- Start with `frontend/scripts/integration-build.mjs`.
- Validate with:
  - `cd frontend && npm test`
  - `npm run test:integration`

### Technical details
- Integration backend endpoint: `http://localhost:18081/graphql`.
- Current verified output: 140 article pages, 190 total files.

## Step 12: Final frontend documentation, validation, and ticket cleanup

This step completed the implementation loop by documenting the validation commands, running the full frontend unit and integration checks, running the meaningful Go package tests, and verifying the docmgr ticket. The frontend README now explains how to run unit tests, the contract smoke, and the integration acceptance script.

The only validation caveat is pre-existing repository state: `go test ./...` includes old debug programs under `ttmp/.../scripts/debug` that all declare `package main` with duplicate `main` functions. The real application packages pass when `ttmp` is excluded. I recorded this because it matters for review: the frontend did not introduce the Go failure, but a future cleanup should exclude or reorganize debug scripts.

**Commit (code):** <pending> — "docs(frontend): document validation and finish ticket"

### What I did
- Updated `frontend/README.md` with validation commands and current integration assertions.
- Ran frontend unit tests: `cd frontend && npm test`.
- Ran frontend integration acceptance: `cd frontend && npm run test:integration`.
- Ran `go test ./...`, observed a pre-existing failure in `ttmp/.../scripts/debug` due duplicate `main` functions.
- Ran `go test $(go list ./... | grep -v '/ttmp/')`, which passed for application packages.
- Ran `docmgr doctor --ticket FAKE-CMS-11TY --stale-after 30`, which passed.
- Marked P7.1–P7.4 complete.

### Why
- A finished implementation needs reproducible validation commands, not just passing local state.
- Recording the `go test ./...` caveat prevents reviewers from attributing an old debug-script issue to the new frontend.

### What worked
- `npm test` passed 8 tests.
- `npm run test:integration` reported `integration ok: 140 article pages, 190 files`.
- `go test $(go list ./... | grep -v '/ttmp/')` passed.
- `docmgr doctor` passed.

### What didn't work
- `go test ./...` failed because old debug scripts under `ttmp/2026/06/17/FAKE-CMS--fake-internal-cms-graphql-api-20minutes-style-for-ssg-workshop/scripts/debug` contain multiple files with `func main()` in the same package:
  - `04-debug-lastinsertid.go`
  - `05-debug-seed-fk.go`
  - `06-debug-seed-step.go`
  - and others.
- Exact representative error: `main redeclared in this block`.

### What I learned
- The repo's documentation workspace currently contains Go debug scripts that are visible to `go test ./...`. Future Go-wide validation should either exclude `ttmp` or add build tags to those debug scripts.

### What was tricky to build
- Final validation spans two ecosystems: Go for the fake CMS and Node/Eleventy for the frontend. The integration script is the bridge, but the repository-level Go test command still sees historical ticket artifacts.

### What warrants a second pair of eyes
- Decide whether to add build tags to historical debug scripts or change documented Go validation to exclude `ttmp`.
- Review whether `npm run test:integration` should be part of CI; it starts a Go process and builds the whole site, which is valuable but slower than unit tests.

### What should be done in the future
- Optional cleanup: add build tags to debug scripts under old ticket workspaces so `go test ./...` is clean.
- Optional frontend polish: richer embed rendering and image optimization.

### Code review instructions
- Review commits from `ed4d647` through the final frontend commit.
- Validate with:
  - `cd frontend && npm test`
  - `cd frontend && npm run test:integration`
  - `go test $(go list ./... | grep -v '/ttmp/')`
  - `docmgr doctor --ticket FAKE-CMS-11TY --stale-after 30`

### Technical details
- Final generated output verified by integration: 140 article pages, 190 files.
- No extra seeding from `20minutes-media.com` was needed because `testdata/cms.db` already contains enough data for the workshop.
