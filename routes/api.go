package routes

import (
	"github.com/gofiber/fiber"
	. "github.com/iamtraining/url-shortener/app"
	controllers "github.com/iamtraining/url-shortener/controlles"
	"github.com/iamtraining/url-shortener/middlewares"
)

func ApiRoutes() {
	api := App.Group("")
	ApiVersions(api)
}

func ApiVersions(api fiber.Router) {
	v1Routes(api)
}

func v1Routes(api fiber.Router) {
	api.Post("/create", middlewares.ValidateUrlPost, controllers.UrlPost)
	api.Get("/:id", middlewares.Redirect)
}
