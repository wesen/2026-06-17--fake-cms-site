# 11ty static frontend for the fake CMS (SSG plugin + design guide)

This is the document workspace for ticket FAKE-CMS-11TY.

## Structure

- **design/**: Design documents and architecture notes
- **reference/**: Reference documentation and API contracts
- **playbooks/**: Operational playbooks and procedures
- **scripts/**: Utility scripts and automation
- **sources/**: External sources and imported documents
- **various/**: Scratch or meeting notes, working notes
- **archive/**: Optional space for deprecated or reference-only artifacts

## Getting Started

Use docmgr commands to manage this workspace:

- Add documents: `docmgr doc add --ticket FAKE-CMS-11TY --doc-type design-doc --title "My Design"`
- Import sources: `docmgr import file --ticket FAKE-CMS-11TY --file /path/to/doc.md`
- Update metadata: `docmgr meta update --ticket FAKE-CMS-11TY --field Status --value review`
