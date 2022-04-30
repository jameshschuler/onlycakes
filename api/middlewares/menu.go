package middlewares

import (
	"context"
	"net/http"
	"onlycakes/domain"
	"onlycakes/models"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// I have no idea what I'm doing 😐
func MenuCtx(menuService *models.MenuService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var menuId uint64
			var menu *models.Menu
			var err error

			if id := chi.URLParam(r, "id"); id != "" {
				menuId, err = strconv.ParseUint(id, 10, 64)
				if err != nil {
					render.Render(w, r, domain.GetErrorResponse(*domain.ErrInvalidId))
					return
				}

				menu, err = menuService.GetMenuById(menuId)
			}

			if err != nil {
				render.Render(w, r, domain.GetErrorResponse(*domain.ErrNotFound))
				return
			}

			ctx := context.WithValue(r.Context(), "menu", menu)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
