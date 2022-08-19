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
ORDER BY created_at DESC;

-- name: ListTransactionsForUserForTicker :many
SELECT * FROM transactions
WHERE 
    user = $1 AND
    ticker = $2
ORDER BY created_at DESC;