package configs

import (
	"github.com/gofiber/fiber"
	. "github.com/iamtraining/url-shortener/app"
)

func BootApp() {
	App = fiber.New()
}
