package db

import (
	"context"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPurchaseTx(t *testing.T) {
	store := NewStore(testDB)
	user := createRandomUser(t)

	n := 5
	amount := 171.52
	ticker := "AAPL"

	errs := make(chan error)
	results := make(chan TradingTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.PurchaseTx(context.Background(), CreateTransactionParams{
				Username: user.Username,
				Ticker:   ticker,
				Quantity: 1,
				Price:    amount,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check user balance
		require.True(t, math.Abs(user.Balance-float64(i+1)*amount-result.User.Balance) < 0.1)
	}

	updatedUser, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.True(t, math.Abs(user.Balance-float64(n)*amount-updatedUser.Balance) < 0.1)

	stock, err := testQueries.GetStockQuantityForUser(context.Background(), GetStockQuantityForUserParams{
		Username: user.Username,
		Ticker:   ticker,
	})
	require.NoError(t, err)
	require.Equal(t, int64(n), stock.TotalQuantity)
}

func TestSellTx(t *testing.T) {
	store := NewStore(testDB)
	user := createRandomUser(t)

	var quantity int64 = 10
	amount := 171.52
	ticker := "AAPL"

	result, err := store.PurchaseTx(context.Background(), CreateTransactionParams{
		Username: user.Username,
		Ticker:   ticker,
		Quantity: quantity,
		Price:    amount,
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.True(t, math.Abs(user.Balance-float64(quantity)*amount-result.User.Balance) < 0.1)

	updatedUser, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)

	errs := make(chan error)
	results := make(chan TradingTxResult)

	for i := 0; i < int(quantity); i++ {
		go func() {
			result, err := store.SellTx(context.Background(), CreateTransactionParams{
				Username: user.Username,
				Ticker:   ticker,
				Quantity: 1,
				Price:    amount,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < int(quantity); i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check user balance
		require.True(t, math.Abs(updatedUser.Balance+float64(i+1)*amount-result.User.Balance) < 0.1)
	}

	updatedUser, err = testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.True(t, math.Abs(user.Balance-updatedUser.Balance) < 0.1)

	stock, err := testQueries.GetStockQuantityForUser(context.Background(), GetStockQuantityForUserParams{
		Username: user.Username,
		Ticker:   ticker,
	})
	require.NoError(t, err)
	require.Equal(t, int64(0), stock.TotalQuantity)
}

func TestPurchaseWithInsufficientBalance(t *testing.T) {
	store := NewStore(testDB)
	user := createRandomUser(t)

	ticker := "AAPL"

	_, err := store.PurchaseTx(context.Background(), CreateTransactionParams{
		Username: user.Username,
		Ticker:   ticker,
		Quantity: 9,
		Price:    10000.5,
	})
	require.NoError(t, err)

	_, err = store.PurchaseTx(context.Background(), CreateTransactionParams{
		Username: user.Username,
		Ticker:   ticker,
		Quantity: 1,
		Price:    20000.,
	})
	require.Error(t, err)
}

func TestSellWithInsufficientStockQuantity(t *testing.T) {
	store := NewStore(testDB)
	user := createRandomUser(t)

	ticker := "AAPL"

	_, err := store.PurchaseTx(context.Background(), CreateTransactionParams{
		Username: user.Username,
		Ticker:   ticker,
		Quantity: 10,
		Price:    100.2,
	})
	require.NoError(t, err)

	_, err = store.SellTx(context.Background(), CreateTransactionParams{
		Username: user.Username,
		Ticker:   ticker,
		Quantity: 11,
		Price:    100.5,
	})
	require.Error(t, err)
}