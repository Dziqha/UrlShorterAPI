package router

import (
	"urlshorter/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func NewRoute(route fiber.Router, urlController *controllers.UrlController) {

	route.All("/:shortUrl?", urlController.IndexShortenUrl)
}
