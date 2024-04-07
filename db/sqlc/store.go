package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Store struct {
	q  *Queries
	db *pgx.Conn
}

func NewStore(db *pgx.Conn) *Store {
	return &Store{
		q:  New(db),
		db: db,
	}
}

func (s *Store) execTx(ctx context.Context, fn func(q *Queries) error) (err error) {
	txOptions := pgx.TxOptions{}
	var tx pgx.Tx
	tx, err = s.db.BeginTx(ctx, txOptions)
	q := New(tx)

	if err != nil {
		return err
	}

	err = fn(q)

	if err != nil {

		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx Error: %v ,err rb Err: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit(ctx)
}

type transferTxParams struct {
	FromAccountId int64
	ToAccountId   int64
	Amount        int32
}

type transferTxResult struct {
	Transfer  Transfer
	FromEntry Entry
	ToEntry   Entry
	From      Account
	To        Account
}

func (s *Store) TransferTx(ctx context.Context, arg transferTxParams) (result transferTxResult, err error) {

	err = s.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: pgtype.Int8{Int64: arg.FromAccountId, Valid: true},
			ToAccountID:   pgtype.Int8{Int64: arg.ToAccountId, Valid: true},
			Amount:        int32(arg.Amount),
		})

		if err != nil {
			fmt.Println("1")
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: pgtype.Int8{Int64: arg.FromAccountId, Valid: true},
			Amount:    -arg.Amount,
		})

		if err != nil {
			fmt.Println("2")

			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: pgtype.Int8{Int64: arg.ToAccountId, Valid: true},
			Amount:    arg.Amount,
		})

		if err != nil {
			fmt.Println("3")

			return err
		}

		//Update Balance
		result.From, err = q.AddAmountToAccount(ctx, AddAmountToAccountParams{
			Amount: arg.Amount,
			ID:     arg.FromAccountId,
		})

		if err != nil {
			fmt.Println("4")

			return err
		}

		result.To, err = q.AddAmountToAccount(ctx, AddAmountToAccountParams{
			Amount: arg.Amount,
			ID:     arg.ToAccountId,
		})

		if err != nil {
			fmt.Println("5")

			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("6")

		return
	}

	return
}
