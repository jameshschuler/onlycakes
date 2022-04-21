package controllers

import (
	"onlycakes/domain"

	"github.com/go-chi/render"
)

func ErrRender(err error) render.Renderer {
	return &domain.ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		Message:        "Error rendering response.",
		ErrorText:      err.Error(),
	}
}
