package domain

import (
	"net/http"

	"github.com/go-chi/render"
)

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, Message: "Resource not found."}
var ErrInvalidId = &ErrResponse{HTTPStatusCode: 400, Message: "Invalid Id."}
var ErrInvalidRequest = &ErrResponse{HTTPStatusCode: 400, Message: "Invalid request."}

type ErrResponse struct {
	HTTPStatusCode int        `json:"statusCode"`
	Message        string     `json:"message,omitempty"`
	Errors         []AppError `json:"errors,omitempty"`
}

type AppError struct {
	Message string `json:"message,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
