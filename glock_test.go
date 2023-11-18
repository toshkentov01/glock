package glock

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"github.com/stretchr/testify/assert"
)

// TestLockAndUnlock tests the Lock and Unlock methods of the Lock struct.
func TestLockAndUnlock(t *testing.T) {
	// Replace with your PostgreSQL connection details
	connStr := "user=your_user password=your_password dbname=your_db sslmode=disable"

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	defer db.Close()

	lock := New(db)
	identifier := int64(123)

	// Test Lock
	err = lock.Lock(context.Background(), identifier)
	assert.NoError(t, err, "Lock should not return an error")

	// Test Unlock
	err = lock.UnLock(context.Background(), identifier)
	assert.NoError(t, err, "Unlock should not return an error")
}

// TestCheckAndLock tests the CheckAndLock method of the Lock struct.
func TestCheckAndLock(t *testing.T) {
	// Replace with your PostgreSQL connection details
	connStr := "user=your_user password=your_password dbname=your_db sslmode=disable"

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	defer db.Close()

	lock := New(db)
	identifier := int64(456)

	// Test CheckAndLock
	locked, err := lock.CheckAndLock(context.Background(), identifier)
	assert.NoError(t, err, "CheckAndLock should not return an error")
	assert.True(t, locked, "The lock should be acquired")

	// Test Unlock
	err = lock.UnLock(context.Background(), identifier)
	assert.NoError(t, err, "Unlock should not return an error")
}
