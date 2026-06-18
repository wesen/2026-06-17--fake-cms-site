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

