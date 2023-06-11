// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package queries

import (
	"time"

	null_v4 "gopkg.in/guregu/null.v4"
)

type Item struct {
	ID         int32     `json:"id"`
	UserID     int32     `json:"user_id"`
	Amount     int32     `json:"amount"`
	TagIds     []int32   `json:"tag_ids"`
	Kind       string    `json:"kind"`
	HappenedAt time.Time `json:"happened_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Tag struct {
	ID        int32          `json:"id"`
	UserID    int32          `json:"user_id"`
	Name      string         `json:"name"`
	Sign      string         `json:"sign"`
	Kind      string         `json:"kind"`
	DeletedAt *time.Time     `json:"deleted_at"`
	X         null_v4.String `json:"x"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type User struct {
	ID        int32     `json:"id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ValidationCode struct {
	ID        int32      `json:"id"`
	Code      string     `json:"code"`
	Email     string     `json:"email"`
	UsedAt    *time.Time `json:"used_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
