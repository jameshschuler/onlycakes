package main

import (
	"fmt"
	"log"
	"net/http"
	"onlycakes/controllers"
	"onlycakes/models"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := makeDbConnection()

	// Migrate the schema
	fmt.Println("Running migrations...")
	db.AutoMigrate(&models.Menu{}, &models.MenuItem{}, &models.MenuItemStep{}, &models.MenuItemStepOption{})
	fmt.Println("Finished!")

	// Setup services
	menuService := models.NewMenuService(db)
	menuItemService := models.NewMenuItemService(db)

	// Setup controllers
	menuController := controllers.NewMenuController(menuService)
	menuItemController := controllers.NewMenuItemController(menuItemService, menuService)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})

	r.Mount("/api/menu", menuController.Routes())
	r.Mount("/api/menu/{id}/menuItem", menuItemController.Routes())

	http.ListenAndServe(":8080", r)
}

func makeDbConnection() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection successful!")
		return db
	}
}
