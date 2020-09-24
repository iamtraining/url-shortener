package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/iamtraining/url-shortener/models"
)

func UrlPost(c *fiber.Ctx) {
	url := c.Locals("url").(models.Url)
	c.JSON(fiber.Map{
		"id":            url.ID,
		"original":      url.Original,
		"short":         url.Short,
		"shortened_url": "localhost:3000/" + url.ID,
	})
	return
}
