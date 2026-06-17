package repo

import (
	"context"
	"database/sql"
)

// Querier is the subset of *sql.DB / *sql.Tx methods the repo uses. Depending
// on this interface (instead of *sql.DB directly) lets tests inject a counting
// implementation to assert N+1 behavior. *sql.DB satisfies Querier.
type Querier interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	PingContext(ctx context.Context) error
}
