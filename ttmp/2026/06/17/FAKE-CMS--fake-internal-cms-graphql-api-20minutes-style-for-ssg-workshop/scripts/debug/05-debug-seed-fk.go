// 05-debug-seed-fk.go — Reproduce the seed FOREIGN KEY failure in isolation.
//
// Run from the repo root:
//   go run ttmp/2026/06/17/FAKE-CMS--*/scripts/debug/05-debug-seed-fk.go
//
// Why this exists:
//   The seed returned: "insert article node #1: constraint failed: FOREIGN KEY
//   constraint failed (787)". The article insert references author_id, so we
//   isolate: insert a locale, an author, then a content_node referencing that
//   author id. If the FK still fails, the cause is column ordering / type
//   mismatch in the INSERT column list (a stray uppercase typo, a wrong
//   position, etc.), not the author id capture.
package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-go-golems/fake-cms/internal/repo"
	_ "modernc.org/sqlite"
)

func main() {
	ctx := context.Background()
	r, err := repo.Open(ctx, ":memory:")
	if err != nil {
		fmt.Println("open:", err)
		return
	}
	defer r.Close()

	tx, _ := r.DB.BeginTx(ctx, nil)

	// show locales present
	rows, qerr := tx.Query("SELECT id, code FROM locale")
	fmt.Println("locale query err:", qerr)
	if qerr == nil {
		fmt.Print("locales: ")
		for rows.Next() {
			var i int
			var c string
			rows.Scan(&i, &c)
			fmt.Printf("%s ", c)
		}
		rows.Close()
		fmt.Println()
	}

	res, _ := tx.Exec("INSERT INTO author(slug,display_name,email,locale) VALUES('x','X','x@x','fr_FR')")
	aid, _ := res.LastInsertId()
	fmt.Println("author id:", aid)

	_, err = tx.Exec(`INSERT INTO content_node
		(kind, slug, title, status, locale, published_at, modified_at, author_id)
		VALUES('ARTICLE','s','t','PUBLISH','fr_FR','2020-01-01T00:00:00Z','2020-01-01T00:00:00Z',?)`, aid)
	fmt.Println("node err:", err)

	// Also probe: does the exact seed INSERT signature work? (matches seed.go)
	_, err = tx.Exec(`INSERT INTO content_node
		(kind, slug, title, excerpt, status, locale, published_at, modified_at, author_id, word_count)
		VALUES ('ARTICLE', 's2', 't2', 'exc', 'PUBLISH', 'fr_FR', '2020', '2020', ?, 99)`, aid)
	fmt.Println("full-signature node err:", err)

	_ = tx.Rollback()
	_ = sql.ErrNoRows
}
