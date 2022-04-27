package domain

import "net/http"

type APIResponse[T any] struct {
	Data  *T           `json:"data,omitempty"`
	Error *ErrResponse `json:"error,omitempty"`
}

func (rd *APIResponse[T]) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
