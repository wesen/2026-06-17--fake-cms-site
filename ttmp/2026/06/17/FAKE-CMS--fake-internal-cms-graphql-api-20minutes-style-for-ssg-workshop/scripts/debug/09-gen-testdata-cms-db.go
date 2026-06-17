// 09-gen-testdata-cms-db.go — Generate the committed testdata/cms.db from the
// deterministic seed, and emit a CONTENT-level sha256 (over normalized, ordered
// rows) rather than the opaque file bytes. SQLite's on-disk format is not
// byte-stable across writes (page allocator / header counters), but the seeded
// DATA is fully deterministic, so we hash the data, not the file.
//
// Run from the repo root:
//   go run ttmp/2026/06/17/FAKE-CMS--*/scripts/debug/09-gen-testdata-cms-db.go
package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"os"

	"github.com/go-go-golems/fake-cms/internal/repo"
	_ "modernc.org/sqlite"
)

// dataTables: tables whose rows are part of the deterministic seed payload,
// in a stable order, with the columns we hash (deterministic columns only —
// exclude nothing here because the seed injects no time.Now()).
var dataTables = []struct {
	name    string
	columns string
}{
	{"author", "id,slug,display_name,email,description,locale"},
	{"media", "id,slug,kind,source_url,alt,width,height,mime_type,locale"},
	{"term", "id,taxonomy_slug,slug,name,parent_term_id"},
	{"content_node", "id,kind,slug,title,excerpt,status,locale,published_at,modified_at,author_id,template,word_count"},
	{"article_post_type", "content_id,post_type"},
	{"block", "id,content_id,order_index,type,data"},
	{"content_term", "content_id,term_id"},
	{"content_media", "content_id,media_id,role"},
	{"seo", "content_id,title,meta_description,canonical,robots,og_json,twitter_json,json_ld,breadcrumbs_json"},
	{"menu_item", "id,menu_slug,label,url,order_index"},
}

func main() {
	ctx := context.Background()
	path := "testdata/cms.db"
	_ = os.MkdirAll("testdata", 0o755)

	r, err := repo.Open(ctx, path)
	if err != nil {
		fmt.Println("open:", err)
		os.Exit(1)
	}
	if err := r.Seed(ctx); err != nil {
		fmt.Println("seed:", err)
		os.Exit(1)
	}

	h, err := contentHash(ctx, r.DB)
	if err != nil {
		fmt.Println("hash:", err)
		os.Exit(1)
	}

	if err := r.Close(); err != nil {
		fmt.Println("close:", err)
		os.Exit(1)
	}

	if err := os.WriteFile(path+".content-sha256", []byte(h+"  "+path+" (content)\n"), 0o644); err != nil {
		fmt.Println("write hash:", err)
		os.Exit(1)
	}
	fmt.Println("content-sha256:", h)
	fmt.Println("wrote:", path, "and", path+".content-sha256")
}

// contentHash builds a sha256 over every row of every data table, ordered by
// primary key, concatenated as "table|col=val|col=val\n".
func contentHash(ctx context.Context, db *sql.DB) (string, error) {
	h := sha256.New()
	for _, t := range dataTables {
		q := fmt.Sprintf("SELECT %s FROM %s ORDER BY 1", t.columns, t.name)
		rows, err := db.QueryContext(ctx, q)
		if err != nil {
			return "", fmt.Errorf("query %s: %w", t.name, err)
		}
		cols, _ := rows.Columns()
		for rows.Next() {
			vals := make([]any, len(cols))
			ptrs := make([]any, len(cols))
			for i := range vals {
				ptrs[i] = &vals[i]
			}
			if err := rows.Scan(ptrs...); err != nil {
				return "", err
			}
			fmt.Fprintf(h, "%s|", t.name)
			for i, c := range cols {
				fmt.Fprintf(h, "%s=%v|", c, vals[i])
			}
			fmt.Fprintln(h)
		}
		rows.Close()
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
