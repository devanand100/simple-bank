// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID        int64
	Owner     string
	Currency  string
	Balance   int32
	CreatedAt pgtype.Timestamptz
}

type Entry struct {
	ID        int32
	AccountID pgtype.Int8
	// can be negative
	Amount    int32
	CreatedAt pgtype.Timestamptz
}

type Transfer struct {
	ID            int32
	FromAccountID pgtype.Int8
	ToAccountID   pgtype.Int8
	// must be positive
	Amount    int32
	CreatedAt pgtype.Timestamptz
}
