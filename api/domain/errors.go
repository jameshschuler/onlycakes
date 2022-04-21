package domain

import (
	"net/http"

	"github.com/go-chi/render"
)

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, Message: "Resource not found."}
var ErrInvalidId = &ErrResponse{HTTPStatusCode: 400, Message: "Invalid Id."}
var ErrInvalidRequest = &ErrResponse{HTTPStatusCode: 400, Message: "Invalid request."}

type ErrResponse struct {
	Err            error `json:"-"`          // low-level runtime error
	HTTPStatusCode int   `json:"statusCode"` // http response status code

	Message   string `json:"message,omitempty"` // user-level status message
	AppCode   int64  `json:"code,omitempty"`    // application-specific error code
	ErrorText string `json:"error,omitempty"`   // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
