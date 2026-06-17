// Package migrations embeds the SQL migration files so the binary is
// self-contained. Files live next to this source (go:embed is package-relative).
package migrations

import "embed"

// FS holds all *.sql files under this package directory.
//
//go:embed *.sql
var FS embed.FS
