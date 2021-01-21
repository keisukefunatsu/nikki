package actions

import "github.com/gobuffalo/buffalo"

// baseAPIResponse extend this struct to generate API responses
type baseAPIResponse struct {
	Message string `json:"message"`
}

// baseErrorAPIResponse extend this struct to generate API responses
type baseErrorAPIResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type baseAPIResources interface {
	List(buffalo.Context) error
	Show(buffalo.Context) error
}
