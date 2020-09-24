package middlewares

import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber"
	. "github.com/iamtraining/url-shortener/app"
	"github.com/iamtraining/url-shortener/libraries"
	"github.com/iamtraining/url-shortener/models"
)

func ValidateUrlPost(c *fiber.Ctx) {
	var (
		g = libraries.Generate(7)
	)

	url := models.UrlForm{
		Url: models.Url{},
	}

	if err := c.BodyParser(&url); err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	if !url.Validate() {
		c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": url.Err,
		})
		return
	}

	if url.Url.Short == "" {
		url.Url.ID = g
	} else {
		url.Url.ID = url.Url.Short
	}

	if err := Store.Create(&url.Url); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	c.Locals("url", url.Url)
	c.Next()
}

func Redirect(c *fiber.Ctx) {
	id := c.Params("id")
	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "url is not exist",
		})
		return
	}

	url, err := Store.Read(id)
	if reflect.DeepEqual(url, models.Url{}) {
		fmt.Println("im idiot")
	}
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	c.Redirect(url.Original, fiber.StatusMovedPermanently)
}
