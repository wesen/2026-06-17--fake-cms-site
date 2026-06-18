---
Title: "11ty Frontend — Investigation Diary"
Ticket: FAKE-CMS-11TY
Status: active
Topics: [11ty, frontend, research]
DocType: reference
Intent: long-term
Owners: []
Summary: "Chronological investigation diary: orienting on the CMS contract, downloading 11ty docs as evidence, and writing the design guide."
LastUpdated: 2026-06-17T18:05:00-04:00
WhatFor: "Resume context for the 11ty frontend ticket."
WhenToUse: "Read before continuing work on FAKE-CMS-11TY."
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
