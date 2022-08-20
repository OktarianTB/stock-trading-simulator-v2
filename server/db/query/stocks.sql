-- name: CreateStockEntryForUser :one
INSERT INTO stocks (
  username,
  ticker,
  quantity
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetStockForUser :one
SELECT * FROM stocks
WHERE username = $1 AND ticker = $2 LIMIT 1;

-- name: ListUserStocks :many
SELECT * FROM stocks
WHERE username = $1
ORDER BY ticker;

-- name: AddStockQuantityForUser :one
UPDATE stocks
SET quantity = quantity + sqlc.arg(quantity)
WHERE username = sqlc.arg(username) AND ticker = sqlc.arg(ticker)
RETURNING *;

-- name: RemoveStockQuantityForUser :one
UPDATE stocks
SET quantity = quantity - sqlc.arg(quantity)
WHERE username = sqlc.arg(username) AND ticker = sqlc.arg(ticker)
RETURNING *;
