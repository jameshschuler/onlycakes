package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"onlycakes/models"
	"strconv"

	"github.com/go-chi/chi"
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
		// r.Use(rs.TodoCtx) // lets have a users map, and lets actually load/manipulate
		r.Get("/", ctrl.GetById)
		//r.Put("/", rs.Update)    // PUT /users/{id} - update a single user by :id
		//r.Delete("/", rs.Delete) // DELETE /users/{id} - delete a single user by :id
	})

	return r
}

// TODO:
// func ArticleCtx(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var article *Article
// 		var err error

// 		if articleID := chi.URLParam(r, "articleID"); articleID != "" {
// 			article, err = dbGetArticle(articleID)
// 		} else if articleSlug := chi.URLParam(r, "articleSlug"); articleSlug != "" {
// 			article, err = dbGetArticleBySlug(articleSlug)
// 		} else {
// 			render.Render(w, r, ErrNotFound)
// 			return
// 		}
// 		if err != nil {
// 			render.Render(w, r, ErrNotFound)
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), "article", article)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func (ctrl *MenuController) Create(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Status Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func (ctrl *MenuController) GetById(w http.ResponseWriter, r *http.Request) {
	menuId := chi.URLParam(r, "id")
	fmt.Println(menuId)
	_, err := strconv.ParseUint(menuId, 10, 32)

	if err == nil {
		menu, _ := ctrl.menuService.GetMenuById(1)
		fmt.Println(menu.Name)
	} else {
		fmt.Println(err)
	}
	resp := make(map[string]string)
	resp["message"] = "Status Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
