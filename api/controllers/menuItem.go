package controllers

import (
	"net/http"
	"onlycakes/domain"
	"onlycakes/middlewares"
	"onlycakes/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type MenuItemController struct {
	menuItemService *models.MenuItemService
	menuService     *models.MenuService
}

func NewMenuItemController(menuItemService *models.MenuItemService, menuService *models.MenuService) *MenuItemController {
	return &MenuItemController{
		menuItemService: menuItemService,
		menuService:     menuService,
	}
}

func (ctrl MenuItemController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middlewares.MenuCtx(ctrl.menuService))

	r.Get("/", ctrl.GetAllMenuItems)

	r.Route("/{id}", func(r chi.Router) {
		//r.Use(ctrl.MenuItemCtx)
		//r.Post("/", ctrl.CreateMenuItem)
		//r.Get("/", ctrl.GetMenuItem)
		//r.Put("/", rs.Update)    // PUT /users/{id} - update a single user by :id
		//r.Delete("/", ctrl.DeleteMenu)
	})

	return r
}

func (ctrl *MenuItemController) GetAllMenuItems(w http.ResponseWriter, r *http.Request) {
	menu := r.Context().Value("menu").(*models.Menu)

	menuItems, err := ctrl.menuItemService.GetAll(menu.ID, false)

	if err != nil {
		render.Render(w, r, domain.GetErrorResponse(*domain.ErrNotFound))
		return
	}

	response := &domain.APIResponse[domain.MenuItemsResponse]{
		Data: domain.MenuItemsResponse{MenuItems: *menuItems},
	}

	if err := render.Render(w, r, response); err != nil {
		render.Render(w, r, ErrRender(err))
	}
}
