package controllers

import (
	"onlycakes/domain"

	"github.com/go-chi/render"
)

func ErrRender(err error) render.Renderer {
	return &domain.ErrResponse{
		HTTPStatusCode: 422,
		Message:        "Error rendering response.",
	}
}

func GetValidationErrorResponse(err error) render.Renderer {
	if err == nil {
		return domain.ErrInvalidRequest
	}

	var errors = []domain.AppError{
		{Message: err.Error()},
	}
	return &domain.ErrResponse{HTTPStatusCode: 400, Errors: errors}
}
