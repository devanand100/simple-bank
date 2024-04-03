// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
  balance, owner , currency
) VALUES (
  $1, $2 , $3
)
RETURNING id, owner, currency, balance, created_at
`

type CreateAccountParams struct {
	Balance  int32
	Owner    string
	Currency string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, createAccount, arg.Balance, arg.Owner, arg.Currency)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Currency,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAccount, id)
	return err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, owner, currency, balance, created_at FROM accounts
ORDER BY owner
LIMIT $1 
OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.Query(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Currency,
			&i.Balance,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE accounts
  set balance = $2
WHERE id = $1
RETURNING id, owner, currency, balance, created_at
`

type UpdateAccountParams struct {
	ID      int64
	Balance int32
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateAccount, arg.ID, arg.Balance)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Currency,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}
