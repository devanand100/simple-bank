package db

import (
	"context"
	"fmt"
	"log"
	"simple-bank/util"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

var DB *pgx.Conn

func TestTransferTx(t *testing.T) {
	databaseUrl := "postgresql://postgres:123456@localhost:8000/simple-bank?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		log.Fatal("Database connection error")
	}
	store := NewStore(conn)

	account1 := CreateTestAccount(t)
	account2 := CreateTestAccount(t)

	var amount int32 = int32(util.RandomInt(0, 10))
	result, err := store.TransferTx(context.Background(), transferTxParams{FromAccountId: int64(account1.ID), ToAccountId: int64(account2.ID), Amount: amount})
	fmt.Println("result............", result)
	fmt.Println("result............", result.Transfer)
	fmt.Println("result............", result.FromEntry)
	fmt.Println("result............", result.ToEntry)
	fmt.Println("err............", err)
	assert.NoError(t, err)
	assert.Equal(t, result.FromEntry.Amount, -amount)
	assert.Equal(t, result.ToEntry.Amount, amount)
	assert.Equal(t, result.Transfer.FromAccountID, account1.ID)
	assert.Equal(t, result.Transfer.ToAccountID, account2.ID)

}
