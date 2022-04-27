package domain

import (
	"errors"
	"net/http"
	"onlycakes/models"
)

type MenuResponse struct {
	Menu *models.Menu `json:"menu"`
}

type MenuRequest struct {
	Name string `json:"name"`
}

func (request *MenuRequest) Bind(r *http.Request) error {
	if request == nil {
		return errors.New("invalid request")
	}

	if request.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
