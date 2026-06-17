// Package repo is the storage layer: it owns the *sql.DB, applies migrations,
// seeds deterministic data, and exposes typed queries returning domain structs.
//
// It knows nothing about GraphQL. Dependencies point UP only.
package repo

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"sort"
	"strings"

	"github.com/go-go-golems/fake-cms/internal/domain"
	"github.com/go-go-golems/fake-cms/internal/migrations"
	_ "modernc.org/sqlite" // register the "sqlite" driver (pure Go, no CGO)
)

// Repo wraps a *sql.DB with typed CMS queries.
type Repo struct {
	DB *sql.DB
	// Q is the read Querier used by query methods. It is DB by default but
	// can be replaced with a counting querier in tests (N+1 guard).
	Q Querier
}

// Open opens (or creates) the SQLite file, enables WAL + foreign_keys for file
// DBs, and applies any pending migrations under migrations/.
//
// Use ":memory:" for ephemeral DBs (WAL is skipped there).
func Open(ctx context.Context, path string) (*Repo, error) {
	// busy_timeout avoids "database is locked" under light concurrency.
	dsn := path
	if path != ":memory:" {
		dsn = path + "?_pragma=busy_timeout(5000)"
	}
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("open sqlite %q: %w", path, err)
	}
	// modernc benefits from a small connection pool; SQLite is single-writer.
	db.SetMaxOpenConns(1)
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping sqlite: %w", err)
	}

	// WAL + FK only make sense for file-backed DBs.
	if path != ":memory:" {
		for _, stmt := range []string{
			"PRAGMA journal_mode=WAL",
			"PRAGMA foreign_keys=ON",
		} {
			if _, err := db.ExecContext(ctx, stmt); err != nil {
				_ = db.Close()
				return nil, fmt.Errorf("pragma %q: %w", stmt, err)
			}
		}
	}

	r := &Repo{DB: db, Q: db}
	if err := r.Migrate(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}
	return r, nil
}

// Close releases the database handle.
func (r *Repo) Close() error { return r.DB.Close() }

// New wraps an existing *sql.DB (already open) and ensures migrations are
// applied. Tests use it to inject an instrumented connection.
func New(ctx context.Context, db *sql.DB) (*Repo, error) {
	r := &Repo{DB: db, Q: db}
	if err := r.Migrate(ctx); err != nil {
		return nil, err
	}
	return r, nil
}

// WithQuerier returns a shallow copy of the repo whose read path uses the
// provided Querier (e.g. a counting wrapper for the N+1 test). The underlying
// DB handle is shared.
func (r *Repo) WithQuerier(q Querier) *Repo {
	cp := *r
	cp.Q = q
	return &cp
}

// Migrate applies embedded SQL files ordered by filename, skipping ones already
// recorded in schema_version. It is idempotent.
func (r *Repo) Migrate(ctx context.Context) error {
	if _, err := r.DB.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS schema_version (
		version INTEGER PRIMARY KEY, applied_at TEXT NOT NULL)`); err != nil {
		return fmt.Errorf("create schema_version: %w", err)
	}

	files, err := fs.ReadDir(migrations.FS, ".")
	if err != nil {
		return fmt.Errorf("read migrations: %w", err)
	}
	names := make([]string, 0, len(files))
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".sql") {
			names = append(names, f.Name())
		}
	}
	sort.Strings(names)

	for _, name := range names {
		version, ok := migrationVersion(name) // "0001_init.sql" -> 1
		if !ok {
			continue
		}
		applied, err := r.isApplied(ctx, version)
		if err != nil {
			return err
		}
		if applied {
			continue
		}
		stmt, err := migrations.FS.ReadFile(name)
		if err != nil {
			return fmt.Errorf("read %s: %w", name, err)
		}
		if _, err := r.DB.ExecContext(ctx, string(stmt)); err != nil {
			return fmt.Errorf("apply %s: %w", name, err)
		}
		if _, err := r.DB.ExecContext(ctx,
			"INSERT INTO schema_version (version, applied_at) VALUES (?, datetime('now'))",
			version); err != nil {
			return fmt.Errorf("record %s: %w", name, err)
		}
	}
	return nil
}

func (r *Repo) isApplied(ctx context.Context, version int) (bool, error) {
	var v int
	err := r.DB.QueryRowContext(ctx,
		"SELECT version FROM schema_version WHERE version = ?", version).Scan(&v)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("check migration %d: %w", version, err)
	}
	return true, nil
}

// migrationVersion parses the leading integer of a filename like "0001_init.sql".
func migrationVersion(name string) (int, bool) {
	i := 0
	for i < len(name) && name[i] >= '0' && name[i] <= '9' {
		i++
	}
	if i == 0 {
		return 0, false
	}
	var v int
	_, err := fmt.Sscanf(name[:i], "%d", &v)
	if err != nil {
		return 0, false
	}
	return v, true
}

// Recreate drops all tables so a seed run starts from a clean slate. Used by
// the `seed` command and by tests.
func (r *Repo) Recreate(ctx context.Context) error {
	tables := []string{
		"menu_item", "menu", "seo", "content_media", "content_term",
		"block", "article_post_type", "content_node", "term", "author",
		"media", "taxonomy", "post_type", "locale", "schema_version",
	}
	for _, t := range tables {
		if _, err := r.DB.ExecContext(ctx, fmt.Sprintf("DROP TABLE IF EXISTS %s", t)); err != nil {
			return fmt.Errorf("drop %s: %w", t, err)
		}
	}
	return r.Migrate(ctx)
}

// Sanity: ensure domain import is used (prevents accidental removal).
var _ = domain.KindArticle
