# glock - Advisory Locks in Golang for PostgreSQL

The `glock` package provides a simple interface for working with advisory locks in PostgreSQL using Golang.

## Installation

To use the `glock` package in your project, you need to install it:

```bash
    go get -u github.com/toshkentov01/glock
```

## Usage

```golang
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/yourusername/glock"
)

const (
	dbConnStr = "user=your_user password=your_password dbname=your_db sslmode=disable"
)

func main() {
	// Connect to PostgreSQL
	db, err := sqlx.Connect("postgres", dbConnStr)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	defer db.Close()

	// Create a Lock instance
	lock := glock.New(db)
	defer lock.Close()

	// Example 1: Acquire and Release Lock
	identifier := int64(123)

	if err := lock.Lock(context.Background(), identifier); err != nil {
		log.Fatal("Failed to acquire lock:", err)
	}

	fmt.Println("Lock acquired successfully.")

	// ... Perform operations under the lock ...

	if err := lock.Unlock(context.Background(), identifier); err != nil {
		log.Fatal("Failed to release lock:", err)
	}

	fmt.Println("Lock released.")

	// Example 2: Check and Acquire Lock
	identifier2 := int64(456)

	locked, err := lock.CheckAndLock(context.Background(), identifier2)
	if err != nil {
		log.Fatal("Failed to check and acquire lock:", err)
	}

	if locked {
		fmt.Println("Lock acquired successfully.")
		// ... Perform operations under the lock ...
		if err := lock.Unlock(context.Background(), identifier2); err != nil {
			log.Fatal("Failed to release lock:", err)
		}

		fmt.Println("Lock released.")
        
	} else {
		fmt.Println("Lock not acquired.")
	}
}

```

