---
Title: Fake internal CMS GraphQL API (20minutes-style) for SSG workshop
Ticket: FAKE-CMS
Status: active
Topics:
    - graphql
    - cms
    - backend
    - workshop
    - go
DocType: index
Intent: long-term
Owners: []
RelatedFiles:
    - Path: cmd/fake-cms/main.go
      Note: Root entrypoint
    - Path: internal/graphql/blocks.go
      Note: Block union + media resolution
    - Path: internal/graphql/schema.go
      Note: Programmatic GraphQL schema
    - Path: internal/repo/repo.go
      Note: Open/Migrate/Recreate + Querier
    - Path: internal/repo/seed.go
      Note: Deterministic SQLite seeder
    - Path: internal/server/server.go
      Note: net/http GraphQL server + GraphiQL
    - Path: schema.graphql
      Note: The GraphQL SDL contract (single source of truth)
    - Path: testdata/cms.db
      Note: Committed deterministic seed
    - Path: ttmp/2026/06/17/FAKE-CMS--fake-internal-cms-graphql-api-20minutes-style-for-ssg-workshop/scripts/01-fetch-sitemaps.py
      Note: Reproducible sitemap inventory generator
    - Path: ttmp/2026/06/17/FAKE-CMS--fake-internal-cms-graphql-api-20minutes-style-for-ssg-workshop/scripts/02-parse-page-schema.py
      Note: Reproducible per-page schema extractor
    - Path: ttmp/2026/06/17/FAKE-CMS--fake-internal-cms-graphql-api-20minutes-style-for-ssg-workshop/scripts/03-parse-article-blocks.py
      Note: Reproducible block-model extractor
    - Path: ttmp/2026/06/17/FAKE-CMS--fake-internal-cms-graphql-api-20minutes-style-for-ssg-workshop/sources/00-sitemap-inventory.md
      Note: Proves WordPress+Yoast; post-type and taxonomy breakdown
    - Path: ttmp/2026/06/17/FAKE-CMS--fake-internal-cms-graphql-api-20minutes-style-for-ssg-workshop/sources/02-page-schemas.md
      Note: Per-page SEO/OG/JSON-LD and body-class extraction
    - Path: ttmp/2026/06/17/FAKE-CMS--fake-internal-cms-graphql-api-20minutes-style-for-ssg-workshop/sources/03-article-blocks.md
      Note: Resolved Article JSON-LD node + block histogram; basis for Article+Block types
ExternalSources: []
Summary: ""
LastUpdated: 2026-06-17T15:44:25.091855419-04:00
WhatFor: ""
WhenToUse: ""
---



# Fake internal CMS GraphQL API (20minutes-style) for SSG workshop

## Overview

<!-- Provide a brief overview of the ticket, its goals, and current status -->

## Key Links

- **Related Files**: See frontmatter RelatedFiles field
- **External Sources**: See frontmatter ExternalSources field

## Status

Current status: **active**

## Topics

- graphql
- cms
- backend
- workshop
- go

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
