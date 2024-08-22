package main

import (
	"log"
	"time"

	"urlshorter/conf"
	"urlshorter/src/controllers"
	"urlshorter/src/models"
	"urlshorter/src/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout: 10 * time.Second,
	})

	url := controllers.NewUrlController()
	db := confiq.Database()
	models.AutoMigrate(db)
	router.NewRoute(app, url)
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
