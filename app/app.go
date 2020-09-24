package app

import (
	"github.com/gofiber/fiber"
	"github.com/iamtraining/url-shortener/store"
)

var App *fiber.App //nolint:gochecknoglobals

var Store *store.Store
