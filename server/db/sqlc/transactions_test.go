package db

import (
	"context"
	"testing"

	util "github.com/OktarianTB/stock-trading-simulator-golang/utils"
	"github.com/stretchr/testify/require"
)

func createPurchaseTransaction(t *testing.T, username, ticker string, quantity int64) Transaction {
	arg := CreateTransactionParams{
		Username: username,
		Ticker:   ticker,
		Quantity: quantity,
		Price:    util.RandomFloat(1, 100),
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transaction)
	require.Equal(t, arg.Username, transaction.Username)
	require.Equal(t, arg.Ticker, transaction.Ticker)
	require.Equal(t, arg.Quantity, transaction.Quantity)

	return transaction
}

func TestCreateTransaction(t *testing.T) {
	user := createRandomUser(t)
	createPurchaseTransaction(t, user.Username, util.RandomString(4), util.RandomInt(-10, 10))
}

func TestGetStockQuantityForUser(t *testing.T) {
	user := createRandomUser(t)
	ticker := util.RandomString(4)
	var quantity int64 = 0

	for i := 0; i < 5; i++ {
		q := util.RandomInt(-10, 10)
		quantity += q
		createPurchaseTransaction(t, user.Username, ticker, q)
	}

	stock, err := testQueries.GetStockQuantityForUser(context.Background(), GetStockQuantityForUserParams{
		Username: user.Username,
		Ticker:   ticker,
	})

	require.NoError(t, err)
	require.NotEmpty(t, stock)
	require.Equal(t, user.Username, stock.Username)
	require.Equal(t, ticker, stock.Ticker)
	require.Equal(t, quantity, stock.TotalQuantity)
}
