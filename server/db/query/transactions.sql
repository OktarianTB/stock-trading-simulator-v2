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
    user = $1
ORDER BY created_at DESC
FOR NO KEY UPDATE;

-- name: ListTransactionsForUserForTicker :many
SELECT * FROM transactions
WHERE 
    user = $1 AND
    ticker = $2
ORDER BY created_at DESC
FOR NO KEY UPDATE;

-- name: ListStockQuantitiesForUser :many
SELECT username, ticker, SUM(quantity) FROM transactions
WHERE username = $1
GROUP BY username, ticker
ORDER BY ticker;

-- name: GetStockQuantityForUser :one
SELECT username, ticker, SUM(quantity) as total_quantity FROM transactions
WHERE username = $1 AND ticker = $2 GROUP BY username, ticker LIMIT 1;
