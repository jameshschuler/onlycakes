package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"onlycakes/domain"
	"onlycakes/models"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type MenuController struct {
	menuService *models.MenuService
}

func NewMenuController(ms *models.MenuService) *MenuController {
	return &MenuController{
		menuService: ms,
	}
}

func (ctrl MenuController) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	//r.Get("/", rs.List)    // GET /users - read a list of users
	r.Post("/", ctrl.Create) // POST /users - create a new user and persist it
	//r.Put("/", rs.Delete)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(ctrl.MenuCtx) // lets have a users map, and lets actually load/manipulate
		r.Get("/", ctrl.GetMenu)
		//r.Put("/", rs.Update)    // PUT /users/{id} - update a single user by :id
		//r.Delete("/", rs.Delete) // DELETE /users/{id} - delete a single user by :id
	})

	return r
}

func (ctrl *MenuController) MenuCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var menu *models.Menu
		var err error

		if id := chi.URLParam(r, "id"); id != "" {
			menuId, e := strconv.ParseUint(id, 10, 64)
			if e != nil {
				resp := &domain.APIResponse[domain.MenuResponse]{
					Error: domain.ErrInvalidId,
				}
				render.Render(w, r, resp)
				return
			}

			menu, err = ctrl.menuService.GetMenuById(menuId)
		} else {
			resp := &domain.APIResponse[domain.MenuResponse]{
				Error: domain.ErrNotFound,
			}
			render.Render(w, r, resp)
			return
		}

		if err != nil {
			resp := &domain.APIResponse[domain.MenuResponse]{
				Error: domain.ErrNotFound,
			}
			render.Render(w, r, resp)
			return
		}

		ctx := context.WithValue(r.Context(), "menu", menu)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (ctrl *MenuController) Create(w http.ResponseWriter, r *http.Request) {
	data := &domain.MenuRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, domain.ErrInvalidRequest)
		return
	}

	resp := make(map[string]string)
	resp["message"] = "Status Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func (ctrl *MenuController) GetMenu(w http.ResponseWriter, r *http.Request) {
	menu := r.Context().Value("menu").(*models.Menu)

	response := &domain.APIResponse[domain.MenuResponse]{
		Data: &domain.MenuResponse{Menu: menu},
	}

	if err := render.Render(w, r, response); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}
