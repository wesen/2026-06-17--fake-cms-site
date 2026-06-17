package repo

import (
	"context"
	"testing"
)

// TestOpenMigrate verifies Open creates the schema on a fresh in-memory DB,
// that Migrate is idempotent (re-running is a no-op), and that Recreate
// cleanly drops + re-applies.
func TestOpenMigrate(t *testing.T) {
	ctx := context.Background()
	r, err := Open(ctx, ":memory:")
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	defer r.Close()

	// Core tables must exist.
	for _, table := range []string{"content_node", "author", "term", "media",
		"block", "seo", "menu", "article_post_type", "schema_version"} {
		var name string
		err := r.DB.QueryRowContext(ctx,
			"SELECT name FROM sqlite_master WHERE type='table' AND name=?", table).Scan(&name)
		if err != nil {
			t.Fatalf("table %s missing after migrate: %v", table, err)
		}
	}

	// Idempotency: running Migrate again must not error.
	if err := r.Migrate(ctx); err != nil {
		t.Fatalf("re-migrate not idempotent: %v", err)
	}

	// Locale/post_type seed rows from 0001_init.sql must be present.
	var n int
	if err := r.DB.QueryRowContext(ctx, "SELECT count(*) FROM locale").Scan(&n); err != nil {
		t.Fatalf("count locale: %v", err)
	}
	if n == 0 {
		t.Fatalf("expected seeded locale rows, got 0")
	}

	// Recreate then migrate should leave schema_version with exactly version 1.
	if err := r.Recreate(ctx); err != nil {
		t.Fatalf("recreate: %v", err)
	}
	var v int
	if err := r.DB.QueryRowContext(ctx,
		"SELECT version FROM schema_version").Scan(&v); err != nil {
		t.Fatalf("read schema_version after recreate: %v", err)
	}
	if v != 1 {
		t.Fatalf("expected schema_version=1 after recreate, got %d", v)
	}
}
