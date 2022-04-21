package domain

import "net/http"

type APIResponse[T any] struct {
	Data  *T           `json:"data"`
	Error *ErrResponse `json:"error"`
}

func (rd *APIResponse[T]) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
