// 06-debug-seed-step.go — Reproduce the seed by calling the REAL repo.Seed
// and, on failure, dump the row counts per table so we can see which FK
// target is empty at the moment of failure.
//
// Run from the repo root:
//   go run ttmp/2026/06/17/FAKE-CMS--*/scripts/debug/06-debug-seed-step.go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-go-golems/fake-cms/internal/repo"
	_ "modernc.org/sqlite"
)

func main() {
	ctx := context.Background()
	r, err := repo.Open(ctx, ":memory:")
	if err != nil {
		fmt.Println("open:", err)
		os.Exit(1)
	}
	defer r.Close()

	err = r.Seed(ctx)
	if err != nil {
		fmt.Println("SEED ERR:", err)
	}
	// dump counts
	for _, t := range []string{"locale", "post_type", "taxonomy", "media",
		"author", "term", "content_node", "article_post_type", "block",
		"content_term", "content_media", "seo", "menu", "menu_item"} {
		var n int
		if e := r.DB.QueryRowContext(ctx, "SELECT count(*) FROM "+t).Scan(&n); e != nil {
			fmt.Printf("%-20s ERR %v\n", t, e)
			continue
		}
		fmt.Printf("%-20s %d\n", t, n)
	}
}
