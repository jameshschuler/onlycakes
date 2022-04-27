package domain

import (
	"net/http"

	"github.com/go-chi/render"
)

type APIResponse[T any] struct {
	Data T `json:"data,omitempty"`
}

func (resp *APIResponse[T]) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type BaseResponse struct {
	HTTPStatusCode int `json:"-"`
}

func (resp *BaseResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, resp.HTTPStatusCode)
	return nil
}
