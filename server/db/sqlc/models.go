// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"time"
)

type Stock struct {
	Username string `json:"username"`
	Ticker   string `json:"ticker"`
	Quantity int64  `json:"quantity"`
}

type Transaction struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Ticker   string `json:"ticker"`
	// can be positive or negative
	Quantity  int64     `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	CreatedAt         time.Time `json:"created_at"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	Balance           float64   `json:"balance"`
}
