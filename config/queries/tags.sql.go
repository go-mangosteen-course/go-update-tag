// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: tags.sql

package queries

import (
	"context"
)

const createTag = `-- name: CreateTag :one
INSERT INTO tags (
  user_id,
  name,
  sign,
  kind,
  x
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING id, user_id, name, sign, kind, deleted_at, x, created_at, updated_at
`

type CreateTagParams struct {
	UserID int32   `json:"user_id"`
	Name   string  `json:"name"`
	Sign   string  `json:"sign"`
	Kind   string  `json:"kind"`
	X      *string `json:"x"`
}

func (q *Queries) CreateTag(ctx context.Context, arg CreateTagParams) (Tag, error) {
	row := q.db.QueryRowContext(ctx, createTag,
		arg.UserID,
		arg.Name,
		arg.Sign,
		arg.Kind,
		arg.X,
	)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Sign,
		&i.Kind,
		&i.DeletedAt,
		&i.X,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTag = `-- name: UpdateTag :one
UPDATE tags
SET
  user_id = $1,
  name = CASE WHEN $2::varchar = '' THEN name ELSE $2 END,
  sign = CASE WHEN $3::varchar = '' THEN sign ELSE $3 END,
  kind = CASE WHEN $4::varchar = '' THEN kind ELSE $4 END
WHERE id = $5
RETURNING id, user_id, name, sign, kind, deleted_at, x, created_at, updated_at
`

type UpdateTagParams struct {
	UserID int32  `json:"user_id"`
	Name   string `json:"name"`
	Sign   string `json:"sign"`
	Kind   string `json:"kind"`
	ID     int32  `json:"id"`
}

func (q *Queries) UpdateTag(ctx context.Context, arg UpdateTagParams) (Tag, error) {
	row := q.db.QueryRowContext(ctx, updateTag,
		arg.UserID,
		arg.Name,
		arg.Sign,
		arg.Kind,
		arg.ID,
	)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Sign,
		&i.Kind,
		&i.DeletedAt,
		&i.X,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}