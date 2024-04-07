package db

import (
	"context"
	"fmt"
	"simple-bank/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	CreateTestAccount(t)
}

func CreateTestAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomString(10),
		Currency: util.RandomCurrency(),
		Balance:  int32(util.RandomInt(0, 1000)),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	fmt.Println("created account....", account)
	assert.NoError(t, err)
	assert.Equal(t, account.Balance, arg.Balance)
	assert.Equal(t, account.Currency, arg.Currency)
	assert.Equal(t, account.Owner, arg.Owner)

	assert.NotZero(t, account.ID)
	assert.NotZero(t, account.CreatedAt)

	return account
}
