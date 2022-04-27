package controllers

import (
	"context"
	"fmt"
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

	r.Get("/", ctrl.GetAll)
	r.Post("/", ctrl.CreateMenu)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(ctrl.MenuCtx)
		r.Get("/", ctrl.GetMenu)
		//r.Put("/", rs.Update)    // PUT /users/{id} - update a single user by :id
		r.Delete("/", ctrl.DeleteMenu)
	})

	return r
}

func (ctrl *MenuController) MenuCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var menuId uint64
		var menu *models.Menu
		var err error

		if id := chi.URLParam(r, "id"); id != "" {
			menuId, err = strconv.ParseUint(id, 10, 64)
			if err != nil {
				render.Render(w, r, GetErrorResponse(*domain.ErrInvalidId))
				return
			}

			menu, err = ctrl.menuService.GetMenuById(menuId)
		}

		if err != nil {
			render.Render(w, r, GetErrorResponse(*domain.ErrNotFound))
			return
		}

		ctx := context.WithValue(r.Context(), "menu", menu)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (ctrl *MenuController) CreateMenu(w http.ResponseWriter, r *http.Request) {
	request := &domain.MenuRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, GetValidationErrorResponse(err))
		return
	}

	menu := models.Menu{Name: request.Name}
	err := ctrl.menuService.CreateMenu(&menu)
	if err != nil {
		render.Render(w, r, GetErrorResponse(*domain.ErrInvalidRequest))
		return
	}

	w.Header().Set("location", fmt.Sprintf("/menu/%v", menu.ID))
	render.Render(w, r, &domain.BaseResponse{HTTPStatusCode: http.StatusCreated})
}

func (ctrl *MenuController) GetAll(w http.ResponseWriter, r *http.Request) {
	menus, err := ctrl.menuService.GetAll()

	if err != nil {
		render.Render(w, r, GetErrorResponse(*domain.ErrNotFound))
		return
	}

	// TODO: Simplify this?
	response := &domain.APIResponse[domain.MenusResponse]{
		Data: domain.MenusResponse{Menus: *menus},
	}

	if err := render.Render(w, r, response); err != nil {
		render.Render(w, r, ErrRender(err))
	}
}

func (ctrl *MenuController) GetMenu(w http.ResponseWriter, r *http.Request) {
	menu := r.Context().Value("menu").(*models.Menu)

	// TODO: Simplify this?
	response := &domain.APIResponse[domain.MenuResponse]{
		Data: domain.MenuResponse{Menu: *menu},
	}

	if err := render.Render(w, r, response); err != nil {
		render.Render(w, r, ErrRender(err))
	}
}

func (ctrl *MenuController) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	menu := r.Context().Value("menu").(*models.Menu)

	err := ctrl.menuService.DeleteMenu(menu.ID)

	if err != nil {
		render.Render(w, r, GetErrorResponse(*domain.ErrInvalidRequest))
		return
	}

	render.Render(w, r, &domain.BaseResponse{HTTPStatusCode: http.StatusNoContent})
}
