package controllers

import (
	"crypto/md5"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"

	"github.com/hitalos/bina/config"
	"github.com/hitalos/bina/models"
)

func loadFromURL(url string) ([]byte, error) {
	statusCode, body, err := fasthttp.Get(nil, url)

	if statusCode != fiber.StatusOK {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not Found")
	}
	return body, err
}

var logoBuf = []byte{}

// GetLogo return logo image
func GetLogo(url string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		if len(logoBuf) == 0 {
			if logoBuf, err = loadFromURL(url); err != nil {
				return err
			}
		}
		c.Set("Content-Type", "image/png")
		return c.Send(logoBuf)
	}
}

// GetPhoto return contact photos
func GetPhoto(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		contact := c.Params("contact")
		entry := models.Entry{}
		if err := entry.GetByAccount(contact); err != nil {
			return err
		}
		if cfg.EnableGravatar && entry.Emails["mail"] != "" {
			hash := md5.Sum([]byte(entry.Emails["mail"]))
			c.Set("Location", fmt.Sprintf("https://www.gravatar.com/avatar/%x", hash))
			return c.SendStatus(fiber.StatusTemporaryRedirect)
		}
		photoBuf, err := loadFromURL(cfg.PhotosURL + entry.ID + ".jpg")
		if err != nil {
			return err
		}

		c.Set("Content-Type", "image/jpeg")
		return c.Send(photoBuf)
	}
}
