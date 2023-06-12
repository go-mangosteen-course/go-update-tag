package api

import (
	"mangosteen/config/queries"
)

type GetPagedTagsRequest struct {
	Page int32  `json:"page"`
	Kind string `json:"kind"`
}
type GetPagesTagsResponse struct {
	Resources []queries.Tag `json:"resources"`
	Pager     Pager         `json:"pager"`
}
type GetTagResponse struct {
	Resource queries.Tag `json:"resource"`
}
type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
	Sign string `json:"sign" binding:"required"`
	Kind string `json:"kind" binding:"required"`
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
