// 08-debug-seedauthors-ids.go — Confirm whether repo.Seed's seedAuthors
// captures non-zero author IDs. We reproduce seedAuthors inline using the
// same s.exec pattern and print the IDs returned.
//
// Run from the repo root:
//   go run ttmp/2026/06/17/FAKE-CMS--*/scripts/debug/08-debug-seedauthors-ids.go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/go-go-golems/fake-cms/internal/repo"
	_ "modernc.org/sqlite"
)

type seeder struct{ tx *sql.Tx }

func (s *seeder) exec(q string, args ...any) (int64, error) {
	res, err := s.tx.Exec(q, args...)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return id, nil
}

func main() {
	ctx := context.Background()
	r, err := repo.Open(ctx, ":memory:")
	if err != nil {
		fmt.Println("open:", err)
		os.Exit(1)
	}
	defer r.Close()
	r.Recreate(ctx)
	tx, _ := r.DB.BeginTx(ctx, nil)
	s := &seeder{tx: tx}

	specs := []string{"a", "b", "c", "d"}
	for _, slug := range specs {
		id, err := s.exec(`INSERT INTO author (slug, display_name, email, description, locale)
			VALUES (?,?,?,?,?)`, slug, slug, slug+"@x", "d", "fr_FR")
		fmt.Printf("slug=%s id=%d err=%v\n", slug, id, err)
	}
	_ = tx.Rollback()
}
