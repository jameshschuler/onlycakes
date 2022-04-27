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
		return &domain.APIErrorResponse{
			Error: *domain.ErrInvalidRequest,
		}
	}

	var errors = []domain.ValidationError{
		{Message: err.Error()},
	}

	return &domain.APIErrorResponse{
		Error: domain.ErrResponse{
			HTTPStatusCode:   400,
			Message:          "One or more validation errors found",
			ValidationErrors: errors},
	}
}

func GetErrorResponse(err domain.ErrResponse) render.Renderer {
	return &domain.APIErrorResponse{
		Error: err,
	}
}
