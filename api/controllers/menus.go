package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Menu struct {
}

func NewMenu() *Menu {
	return &Menu{}
}

func (m Menu) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	//r.Get("/", rs.List)    // GET /users - read a list of users
	r.Post("/", m.Create) // POST /users - create a new user and persist it
	//r.Put("/", rs.Delete)

	r.Route("/{id}", func(r chi.Router) {
		// r.Use(rs.TodoCtx) // lets have a users map, and lets actually load/manipulate
		r.Get("/", m.GetById)
		//r.Put("/", rs.Update)    // PUT /users/{id} - update a single user by :id
		//r.Delete("/", rs.Delete) // DELETE /users/{id} - delete a single user by :id
	})

	return r
}

func (menu *Menu) Create(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Status Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func (menu *Menu) GetById(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Status Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
