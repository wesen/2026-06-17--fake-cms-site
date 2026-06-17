package repo

import (
	"context"
	"testing"
)

// TestSeedDeterministic verifies two independent Seed() runs produce identical
// row counts (a proxy for determinism) and that the dataset is non-trivial.
// Byte-level determinism is asserted via the committed testdata/cms.db sha256
// by the `seed` CLI integration in a later phase.
func TestSeedDeterministic(t *testing.T) {
	ctx := context.Background()

	counts := func(t *testing.T, r *Repo) map[string]int {
		t.Helper()
		out := map[string]int{}
		for _, tbl := range []string{"content_node", "block", "author", "media", "term", "seo"} {
			var n int
			if err := r.DB.QueryRowContext(ctx, "SELECT count(*) FROM "+tbl).Scan(&n); err != nil {
				t.Fatalf("count %s: %v", tbl, err)
			}
			out[tbl] = n
		}
		return out
	}

	r1, err := Open(ctx, ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	if err := r1.Seed(ctx); err != nil {
		t.Fatal(err)
	}
	c1 := counts(t, r1)
	r1.Close()

	r2, err := Open(ctx, ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	if err := r2.Seed(ctx); err != nil {
		t.Fatal(err)
	}
	c2 := counts(t, r2)
	r2.Close()

	for k, v := range c1 {
		if c2[k] != v {
			t.Errorf("determinism: %s count differs: %d vs %d", k, v, c2[k])
		}
		if v == 0 {
			t.Errorf("expected non-zero %s rows", k)
		}
	}
	// Sanity: the bulk is best-cases + actualites.
	if c1["content_node"] < 100 {
		t.Errorf("expected >=100 content nodes, got %d", c1["content_node"])
	}
	if c1["block"] < 500 {
		t.Errorf("expected >=500 blocks, got %d", c1["block"])
	}
}
