-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: AddUserBalance :one
UPDATE users
SET balance = balance + sqlc.arg(amount)
WHERE username = sqlc.arg(username)
RETURNING *;

-- name: RemoveUserBalance :one
UPDATE users
SET balance = balance - sqlc.arg(amount)
WHERE username = sqlc.arg(username)
RETURNING *;