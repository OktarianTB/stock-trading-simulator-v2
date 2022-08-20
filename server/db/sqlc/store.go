package db

import (
	"context"
	"database/sql"
	"fmt"
)

type TradingTxResult struct {
	User        User        `json:"user"`
	Transaction Transaction `json:"transaction"`
	StockEntry  Stock       `json:"stock_entry"`
}

type Store interface {
	Querier
	PurchaseTx(ctx context.Context, arg CreateTransactionParams) (TradingTxResult, error)
	SellTx(ctx context.Context, arg CreateTransactionParams) (TradingTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transaction
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func (store *SQLStore) PurchaseTx(ctx context.Context, arg CreateTransactionParams) (TradingTxResult, error) {
	var result TradingTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Check that the user has a sufficient balance to purchase these stocks
		user, err := q.GetUser(ctx, arg.Username)
		if err != nil {
			return err
		}

		amount := arg.Price * float64(arg.Quantity)

		if user.Balance < amount {
			return fmt.Errorf("insufficient funds to purchase %v of %v", arg.Quantity, arg.Ticker)
		}

		// Then updates balances as appropriate
		_, err = q.GetStockForUser(ctx, GetStockForUserParams{
			Username: arg.Username,
			Ticker:   arg.Ticker,
		})
		if err != nil {
			if err == sql.ErrNoRows {
				result.StockEntry, err = q.CreateStockEntryForUser(ctx, CreateStockEntryForUserParams{
					Username: arg.Username,
					Ticker:   arg.Ticker,
					Quantity: arg.Quantity,
				})
				if err != nil {
					return err
				}
			}

			return err
		} else {
			result.StockEntry, err = q.AddStockQuantityForUser(ctx, AddStockQuantityForUserParams{
				Username: arg.Username,
				Ticker:   arg.Ticker,
				Quantity: arg.Quantity,
			})
			if err != nil {
				return err
			}
		}

		result.User, err = q.RemoveUserBalance(ctx, RemoveUserBalanceParams{
			Username: arg.Username,
			Amount:   amount,
		})
		if err != nil {
			return err
		}

		// Record transaction for reference
		result.Transaction, err = q.CreateTransaction(ctx, CreateTransactionParams{
			Username: arg.Username,
			Ticker:   arg.Ticker,
			Quantity: arg.Quantity,
			Price:    arg.Price,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

func (store *SQLStore) SellTx(ctx context.Context, arg CreateTransactionParams) (TradingTxResult, error) {
	var result TradingTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Check that the user has the required number of stocks in their portfolio
		stock, err := q.GetStockForUser(ctx, GetStockForUserParams{
			Username: arg.Username,
			Ticker:   arg.Ticker,
		})
		if err != nil {
			return err
		}

		if stock.Quantity < arg.Quantity {
			return fmt.Errorf("user does not have the required quantity of the stock %v", arg.Ticker)
		}

		// Then updates balances as appropriate
		result.StockEntry, err = q.RemoveStockQuantityForUser(ctx, RemoveStockQuantityForUserParams{
			Username: arg.Username,
			Ticker:   arg.Ticker,
			Quantity: arg.Quantity,
		})
		if err != nil {
			return err
		}

		result.User, err = q.AddUserBalance(ctx, AddUserBalanceParams{
			Username: arg.Username,
			Amount:   arg.Price * float64(arg.Quantity),
		})
		if err != nil {
			return err
		}

		// Record transaction for reference
		result.Transaction, err = q.CreateTransaction(ctx, CreateTransactionParams{
			Username: arg.Username,
			Ticker:   arg.Ticker,
			Quantity: -arg.Quantity,
			Price:    arg.Price,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
