package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("create_books", r.CreateBooks)
	api.Delete("delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBooksByID)
	api.Get("/books", r.GetBooks)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Could not load the database. ")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
