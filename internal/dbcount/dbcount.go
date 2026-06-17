// Package dbcount provides a Querier that wraps *sql.DB and counts every
// Query/QueryRow/Exec operation. It is used by the N+1 regression test to
// assert the repo's list + relationship resolution stays at a fixed small
// number of SQL statements.
package dbcount

import (
	"context"
	"database/sql"
	"sync/atomic"
)

// Counter is a Querier wrapper that records the number of read operations.
type Counter struct {
	DB  *sql.DB
	n   int64
}

// New wraps db. Operations are counted via the returned Counter and also
// delegated to db.
func New(db *sql.DB) *Counter {
	return &Counter{DB: db}
}

// Count returns the number of observed operations so far.
func (c *Counter) Count() int64 { return atomic.LoadInt64(&c.n) }

// Reset zeroes the counter.
func (c *Counter) Reset() { atomic.StoreInt64(&c.n, 0) }

// ExecContext implements repo.Querier.
func (c *Counter) ExecContext(ctx context.Context, q string, args ...any) (sql.Result, error) {
	atomic.AddInt64(&c.n, 1)
	return c.DB.ExecContext(ctx, q, args...)
}

// QueryContext implements repo.Querier.
func (c *Counter) QueryContext(ctx context.Context, q string, args ...any) (*sql.Rows, error) {
	atomic.AddInt64(&c.n, 1)
	return c.DB.QueryContext(ctx, q, args...)
}

// QueryRowContext implements repo.Querier. We count at row-scan resolution by
// returning a wrapper; the count is incremented here (one QueryRow == one SQL
// statement in the underlying driver).
func (c *Counter) QueryRowContext(ctx context.Context, q string, args ...any) *sql.Row {
	atomic.AddInt64(&c.n, 1)
	return c.DB.QueryRowContext(ctx, q, args...)
}

// PingContext implements repo.Querier.
func (c *Counter) PingContext(ctx context.Context) error {
	return c.DB.PingContext(ctx)
}
