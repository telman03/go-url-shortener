package main

import (
	"github.com/telman03/go-url-service/cache"
	"github.com/telman03/go-url-service/config"
	"github.com/telman03/go-url-service/database"
	"github.com/gofiber/fiber/v2"
	"github.com/telman03/go-url-service/models"
	"github.com/telman03/go-url-service/handlers"
)

func main() {
	// load env
	config.LoadEnv()

	database.ConnectDB()
	database.DB.AutoMigrate(&models.URL{})
	cache.InitRedis()
	app := fiber.New()


	app.Post("/shorten", handlers.ShortenURL)
	app.Get("/:shortcode", handlers.RedirectURL)
	app.Get("/analytics/:shortcode", handlers.GetAnalytics)

	app.Listen(":8080")
}
