package api

import (
	"mangosteen/config/queries"

	"gopkg.in/guregu/null.v4"
)

type CreateTagRequest struct {
	Name string      `json:"name" binding:"required"`
	Sign string      `json:"sign" binding:"required"`
	Kind string      `json:"kind" binding:"required"`
	X    null.String `json:"x"`
}
type UpdateTagRequest struct {
	Name string `json:"name"`
	Sign string `json:"sign"`
	Kind string `json:"kind"`
}
type CreateTagResponse struct {
	Resource queries.Tag `json:"resource"`
}
type UpdateTagResponse CreateTagResponse
