---
title: "Sources manifest — 11ty documentation (defuddle)"
doc-type: reference
topics: [11ty, research]
status: active
intent: long-term
summary: "Index of official Eleventy (11ty) docs downloaded with defuddle into sources/. Cited as evidence by the design doc."
---

# Sources — Eleventy (11ty) official documentation

These files are **snapshots of the official Eleventy documentation**, captured
with `defuddle parse <url> --md | fold -w 100 -s` on 2026-06-17 and committed so
the design doc is reproducible offline. They are **the authoritative evidence**
for every claim this ticket makes about how 11ty works: the plugin contract, the
data cascade, pagination, collections, and the programmatic API.

Each file carries YAML frontmatter with its original `external-url` and a one-line
re-run command, so a future maintainer can refresh it.

## Manifest

| File | Topic | Why it is cited |
|------|-------|-----------------|
| `10-11ty-plugins.md` | Plugins overview | What an 11ty plugin *is*; the `addPlugin` entry point. |
| `11-11ty-create-plugin.md` | **Create or use plugins** | The plugin author contract: `function(eleventyConfig){…}`, options object, async `addPlugin` (v3). |
| `12-11ty-custom-templates.md` | Custom template languages | `addExtension(...)` + `compile`/`getData` — how to teach 11ty a new input format (used to discuss "batteries-included" page generation). |
| `13-11ty-data-cascade.md` | **Data cascade** | The precedence ladder (computed → front matter → directory → config-global → global data). Explains *where* the plugin's fetched data sits. |
| `14-11ty-pagination.md` | **Pagination** | Paginating a dataset with `size: 1` + a dynamic `permalink` alias to emit **one output page per item**. This is the single most important technique for mirroring CMS URLs. |
| `15-11ty-collections.md` | Collections | `eleventyConfig.addCollection(...)` to build typed groups (all articles, by post type, by tag…). |
| `16-11ty-js-data-files.md` | JavaScript data files | Async functions as data; the shape our `_data/cms.js` global-data fetcher takes. |
| `17-11ty-programmatic.md` | **Programmatic API** | `new Eleventy(); await elev.write()` — how an integration test drives a build in-process. |
| `18-11ty-config.md` | Configuration API | `addGlobalData`, `addFilter`, `addShortcode`, `setTemplateFormats` — the plugin's toolbox. |
| `19-11ty-config-global-data.md` | **Config global data** | Dedicated `addGlobalData` documentation; confirms function values are evaluated and async-friendly. |
| `20-11ty-fetch.md` | Eleventy Fetch | Official caching helper; useful as an optional later layer for CMS HTTP calls. |
| `21-11ty-virtual-templates.md` | **Virtual templates** | Documents `eleventyConfig.addTemplate`; the correct mechanism if a plugin contributes generated templates. |
| `22-live-schema-check.md` | **Executable fake-cms schema check** | Captured from `./fake-cms serve --path testdata/cms.db`; highlights mismatches between checked-in SDL/API docs and the running schema. |

## How to refresh

```bash
# From the repo root.
SRC=ttmp/2026/06/17/FAKE-CMS-11TY--11ty-static-frontend-for-the-fake-cms-ssg-plugin-design-guide/sources
defuddle parse https://www.11ty.dev/docs/create-plugin/ --md | fold -w 100 -s > "$SRC/11-11ty-create-plugin.md"
```

> **Defuddle quirk (documented in the defuddle skill):** output is emitted on a
> single line. `fold -w 100 -s` re-wraps it to ~100 columns. A file that reports
> `1 line` but tens of KB is *not* a failed extraction — check the byte count.

## Not in sources (read but not mirrored)

- `https://www.11ty.dev/docs/plugins/fetch/` — `@11ty/eleventy-fetch`, the
  recommended caching layer for the plugin's HTTP calls. Summarized inline in
  the design doc (Decision record) rather than mirrored, because the caching
  strategy is a *choice*, not a constraint.
- `https://www.11ty.dev/docs/data-global-custom/` — `addGlobalData`; folded into
  the Configuration (`18-…config.md`) and data-cascade (`13-…`) evidence.
