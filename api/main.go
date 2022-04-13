package main

import (
	"fmt"
	"log"
	"net/http"
	"onlycakes/controllers"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Menu struct {
	Name      string `gorm:"not null;varchar(255)"`
	MenuItems []MenuItem
	gorm.Model
}

type MenuItem struct {
	Name   string  `gorm:"not null;varchar(255)"`
	Price  float32 `gorm:"not null"`
	Active bool    `gorm:"not null;default:true"`
	MenuId uint    `gorm:"not null"`
	gorm.Model
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
	}

	// Migrate the schema
	db.AutoMigrate(&Menu{})
	db.AutoMigrate(&MenuItem{})

	menuController := controllers.NewMenu()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})

	r.Mount("/menu", menuController.Routes())

	http.ListenAndServe(":8080", r)
}
