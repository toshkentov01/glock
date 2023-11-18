package glock

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// Lock represents a structure for managing advisory locks in PostgreSQL.
type Lock struct {
	db *sqlx.DB
}

// New creates a new Lock instance with the provided SQL connection.
func New(db *sqlx.DB) *Lock {
	return &Lock{
		db: db,
	}
}

// Lock acquires an advisory lock identified by the given identifier.
func (l *Lock) Lock(ctx context.Context, identifier int64) error {
	query := "SELECT pg_advisory_lock($1)"

	// Execute the advisory lock query with the provided identifier.
	_, err := l.db.ExecContext(ctx, query, identifier)
	return err
}

// CheckAndLock attempts to acquire an advisory lock identified by the given identifier,
// and it returns true if the lock is acquired successfully, false otherwise.
func (l *Lock) CheckAndLock(ctx context.Context, identifier int64) (bool, error) {
	var (
		locked bool
		err    error
	)

	query := "SELECT pg_try_advisory_lock($1)"

	// Execute the advisory lock query with the provided identifier,
	// and scan the result into the 'locked' variable.
	err = l.db.QueryRowContext(ctx, query, identifier).Scan(&locked)
	return locked, err
}

// Unlock releases the advisory lock identified by the given identifier.
func (l *Lock) UnLock(ctx context.Context, identifier int64) error {
	query := "SELECT pg_advisory_unlock($1)"

	// Execute the advisory unlock query with the provided identifier.
	_, err := l.db.ExecContext(ctx, query, identifier)
	return err
}
