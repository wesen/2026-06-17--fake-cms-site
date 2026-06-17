// 07-debug-author-id.go — Call the unexported seeder path indirectly: open a
// repo, recreate, then manually run the first two seed steps (authors, media)
// by reimplementing them inline, to confirm LastInsertId for the author row
// matches what a content_node FK check accepts.
//
// Run from the repo root:
//   go run ttmp/2026/06/17/FAKE-CMS--*/scripts/debug/07-debug-author-id.go
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
	if err := r.Recreate(ctx); err != nil {
		fmt.Println("recreate:", err)
		os.Exit(1)
	}
	tx, _ := r.DB.BeginTx(ctx, nil)

	res, err := tx.Exec(`INSERT INTO author (slug, display_name, email, description, locale)
		VALUES (?,?,?,?,?)`, "eking20minutes-fr", "Élise King", "eking@20minutes.fr", "desc", "fr_FR")
	if err != nil {
		fmt.Println("author insert:", err)
		_ = tx.Rollback()
		return
	}
	aid, _ := res.LastInsertId()
	fmt.Println("author lastinsertid:", aid)

	// verify the row exists with that id
	var got int64
	if e := tx.QueryRow("SELECT id FROM author WHERE slug=?", "eking20minutes-fr").Scan(&got); e != nil {
		fmt.Println("lookup author:", e)
	} else {
		fmt.Println("author row id:", got)
	}

	// now the EXACT insert the seed does for article #1
	_, err = tx.Exec(`INSERT INTO content_node
		(kind, slug, title, excerpt, status, locale, published_at, modified_at, author_id, word_count)
		VALUES ('ARTICLE', ?, ?, ?, 'PUBLISH', ?, ?, ?, ?, ?)`,
		"slug-1", "Title", "excerpt", "fr_FR", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", aid, 40)
	fmt.Println("article node insert err:", err)
	_ = tx.Rollback()
}
