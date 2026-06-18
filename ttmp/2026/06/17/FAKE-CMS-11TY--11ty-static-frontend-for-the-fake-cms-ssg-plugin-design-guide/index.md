---
Title: 11ty static frontend for the fake CMS (SSG plugin + design guide)
Ticket: FAKE-CMS-11TY
Status: active
Topics:
    - 11ty
    - ssg
    - frontend
    - workshop
    - javascript
DocType: index
Intent: long-term
Owners: []
RelatedFiles:
    - Path: internal/doc/api-reference.md
      Note: Per-field API + the one-request full-render query shape
    - Path: internal/doc/workshop-ssg.md
      Note: The acceptance criteria (URL conventions + 7 block types)
    - Path: schema.graphql
      Note: The GraphQL contract the plugin queries (single source of truth)
    - Path: ttmp/2026/06/17/FAKE-CMS-11TY--11ty-static-frontend-for-the-fake-cms-ssg-plugin-design-guide/design-doc/01-11ty-frontend-design-implementation-guide.md
      Note: Primary intern-facing design & implementation guide for the 11ty frontend
    - Path: ttmp/2026/06/17/FAKE-CMS-11TY--11ty-static-frontend-for-the-fake-cms-ssg-plugin-design-guide/design-doc/02-corrected-11ty-cms-static-site-guide.md
      Note: Second-pass corrected guide based on executable fake-cms schema and tested Eleventy 3 APIs
    - Path: ttmp/2026/06/17/FAKE-CMS-11TY--11ty-static-frontend-for-the-fake-cms-ssg-plugin-design-guide/reference/02-implementation-diary.md
      Note: Investigation diary (research steps 1-3)
    - Path: ttmp/2026/06/17/FAKE-CMS-11TY--11ty-static-frontend-for-the-fake-cms-ssg-plugin-design-guide/sources/00-sources-readme.md
      Note: Manifest of defuddle-downloaded 11ty docs
    - Path: ttmp/2026/06/17/FAKE-CMS-11TY--11ty-static-frontend-for-the-fake-cms-ssg-plugin-design-guide/sources/11-11ty-create-plugin.md
      Note: The plugin author contract (addPlugin signature
    - Path: ttmp/2026/06/17/FAKE-CMS-11TY--11ty-static-frontend-for-the-fake-cms-ssg-plugin-design-guide/sources/13-11ty-data-cascade.md
      Note: Cascade precedence ladder; justifies addGlobalData at rung 2
    - Path: ttmp/2026/06/17/FAKE-CMS-11TY--11ty-static-frontend-for-the-fake-cms-ssg-plugin-design-guide/sources/14-11ty-pagination.md
      Note: size:1 + dynamic permalink = one page per item (the core technique)
    - Path: ttmp/2026/06/17/FAKE-CMS-11TY--11ty-static-frontend-for-the-fake-cms-ssg-plugin-design-guide/sources/17-11ty-programmatic.md
      Note: new Eleventy(); await elev.write()/toJSON() for in-process tests
    - Path: ttmp/2026/06/17/FAKE-CMS-11TY--11ty-static-frontend-for-the-fake-cms-ssg-plugin-design-guide/sources/19-11ty-config-global-data.md
      Note: Official addGlobalData source used by corrected guide
    - Path: ttmp/2026/06/17/FAKE-CMS-11TY--11ty-static-frontend-for-the-fake-cms-ssg-plugin-design-guide/sources/21-11ty-virtual-templates.md
      Note: Official addTemplate/virtual-template source used to correct plugin-template guidance
    - Path: ttmp/2026/06/17/FAKE-CMS-11TY--11ty-static-frontend-for-the-fake-cms-ssg-plugin-design-guide/sources/22-live-schema-check.md
      Note: Live evidence from ./fake-cms serve --path testdata/cms.db; documents executable schema mismatches
ExternalSources: []
Summary: ""
LastUpdated: 2026-06-17T18:01:37.342576511-04:00
WhatFor: ""
WhenToUse: ""
---















# 11ty static frontend for the fake CMS (SSG plugin + design guide)

## Overview

<!-- Provide a brief overview of the ticket, its goals, and current status -->

## Key Links

- **Related Files**: See frontmatter RelatedFiles field
- **External Sources**: See frontmatter ExternalSources field

## Status

Current status: **active**

## Topics

- 11ty
- ssg
- frontend
- workshop
- javascript

## Tasks

See [tasks.md](./tasks.md) for the current task list.

## Changelog

See [changelog.md](./changelog.md) for recent changes and decisions.

## Structure

- design/ - Architecture and design documents
- reference/ - Prompt packs, API contracts, context summaries
- playbooks/ - Command sequences and test procedures
- scripts/ - Temporary code and tooling
- various/ - Working notes and research
- archive/ - Deprecated or reference-only artifacts
