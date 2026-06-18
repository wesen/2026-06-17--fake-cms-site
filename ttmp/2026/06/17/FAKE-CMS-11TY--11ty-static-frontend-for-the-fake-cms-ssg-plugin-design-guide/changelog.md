# Changelog

## 2026-06-17

- Initial workspace created


## 2026-06-17

Researched 11ty plugin model; downloaded 9 official docs into sources/ via defuddle. Wrote intern-facing 11ty frontend design & implementation guide (21 sections, Mermaid diagrams, pseudocode, JS sketches) + investigation diary. Design: an 'adapter + reference templates' plugin (fetches CMS via addGlobalData, centralizes legacy URLs incl. /rubrique/, renders the block union) + a reference 11ty site. Acceptance lifted from workshop-ssg.

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/ttmp/2026/06/17/FAKE-CMS-11TY--11ty-static-frontend-for-the-fake-cms-ssg-plugin-design-guide/design-doc/01-11ty-frontend-design-implementation-guide.md — Primary deliverable


## 2026-06-18

Reviewed the first 11ty guide against Eleventy 3.1.6 and the executable fake-cms schema. Found and documented concrete mismatches: no site/pages list, no categories(first), invalid direct BlockUnion.order selection, weak SEO og/twitter strings, empty default cms.db. Wrote separate corrected guide recommending a project-local plugin for data/helpers plus visible templates, with client-side taxonomy grouping and optional pageSlugs until backend schema is fixed.

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/ttmp/2026/06/17/FAKE-CMS-11TY--11ty-static-frontend-for-the-fake-cms-ssg-plugin-design-guide/design-doc/02-corrected-11ty-cms-static-site-guide.md — Corrected implementation guide


## 2026-06-18

Committed documentation baseline (33c6fa7) and expanded tasks.md into detailed P0-P7 implementation plan plus A1-A6 acceptance gates. Added .gitignore rule for root fake-cms build binary (010c602).

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/ttmp/2026/06/17/FAKE-CMS-11TY--11ty-static-frontend-for-the-fake-cms-ssg-plugin-design-guide/tasks.md — Detailed implementation checklist


## 2026-06-18

P0 scaffold complete: added frontend/ Eleventy project with hardcoded cms global data, visible index/article templates, base layout, CSS passthrough, npm scripts, and README. Build emits /best-cases/hardcoded-smoke-article/.

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/frontend/eleventy.config.cjs — Initial Eleventy config and hardcoded CMS smoke data


## 2026-06-18

P1/P2 complete: added GraphQL client with cursor pagination and correct BlockUnion fragment, contract smoke script against seeded fake-cms, normalization helpers for URL paths/indexes/sitemap, and Node unit tests. Seeded dataset has 140 articles, so no extra seeding needed now.

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/frontend/_config/fakeCmsClient.cjs — GraphQL client and current executable schema queries


## 2026-06-18

P3 complete: added pure renderBlocks module covering all seven CMS block variants, HTML escaping, heading clamping, order sorting, unknown-block failure, and Node tests.

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/frontend/_config/renderBlocks.cjs — Block renderer used by Eleventy shortcode


## 2026-06-18

P4 complete: added project-local fakeCmsPlugin, replaced hardcoded global data with fetchCms+normalizeCms, registered URL/JSON/renderBlocks helpers, and verified real CMS build generates homepage + 140 article pages.

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/frontend/_config/fakeCmsPlugin.cjs — Eleventy plugin boundary for CMS data and helpers


## 2026-06-18

P5 complete: added visible Nunjucks templates for articles, post-type archives, tags (/rubrique), categories (/archives), authors, optional pages, homepage, shared head partial, article cards, sitemap.xml, and expanded CSS. Build now writes 189 Eleventy pages from seeded CMS data.

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/frontend/src/sitemap.xml.njk — Sitemap template over normalized URLs


## 2026-06-18

P6 acceptance checks complete: added integration script that starts seeded fake-cms, builds Eleventy, checks article count, no /tag/, /rubrique/ pages, sitemap coverage for generated HTML, and parseable JSON-LD. Fixed go-run process cleanup.

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/frontend/scripts/integration-build.mjs — End-to-end acceptance script


## 2026-06-18

P7 complete: documented frontend validation commands, ran npm unit tests, integration acceptance (140 article pages, 190 files), Go package tests excluding ttmp, and docmgr doctor. Noted pre-existing go test ./... failure from historical ttmp debug scripts with duplicate main functions.

### Related Files

- /home/manuel/code/wesen/2026-06-17--fake-cms-site/frontend/README.md — Frontend validation and backend contract notes

