-- name: CreateTransaction :one
INSERT INTO transactions (
  username,
  ticker,
  quantity,
  price
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: ListTransactionsForUser :many
SELECT * FROM transactions
WHERE 
    username = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: ListTransactionsForUserForTicker :many
SELECT * FROM transactions
WHERE 
    username = $1 AND
    ticker = $2
ORDER BY created_at DESC
LIMIT $3
OFFSET $4;

-- name: GetPurchasePriceForTicker :one
SELECT SUM(price * quantity)::float as purchase_price FROM transactions
WHERE 
    username = $1 AND
    ticker = $2;

-- name: ListStockQuantitiesForUser :many
SELECT ticker, SUM(quantity) as quantity FROM transactions
WHERE username = $1
GROUP BY username, ticker
ORDER BY ticker;

-- name: GetStockQuantityForUser :one
SELECT username, ticker, SUM(quantity) as total_quantity FROM transactions
WHERE username = $1 AND ticker = $2 GROUP BY username, ticker LIMIT 1;
