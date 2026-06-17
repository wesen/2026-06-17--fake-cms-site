// Package doc embeds the glazed help entries under doc/ so the binary is
// self-contained: `fake-cms help api-reference` works with no network access.
package doc

import "embed"

// FS holds all help markdown files under the project doc/ directory.
//
//go:embed *.md
var FS embed.FS
