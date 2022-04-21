package domain

import (
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
	// a.Article is nil if no Article fields are sent in the request. Return an
	// error to avoid a nil pointer dereference.
	// if a.Article == nil {
	// 	return errors.New("missing required Article fields.")
	// }

	// a.User is nil if no Userpayload fields are sent in the request. In this app
	// this won't cause a panic, but checks in this Bind method may be required if
	// a.User or futher nested fields like a.User.Name are accessed elsewhere.

	// just a post-process after a decode..
	// a.ProtectedID = ""                                 // unset the protected ID
	// a.Article.Title = strings.ToLower(a.Article.Title) // as an example, we down-case
	return nil
}
