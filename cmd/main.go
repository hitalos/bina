package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/hitalos/bina/config"
	"github.com/hitalos/bina/controllers"
)

func main() {
	configFilepath := flag.String("c", "config.yml", "Path of config file")
	flag.Parse()
	cfg := config.Load(*configFilepath)

	app := fiber.New(fiber.Config{
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	})
	app.Use(compress.New())
	if os.Getenv("DEBUG") == "1" {
		app.Use(logger.New())
	}

	contacts := app.Group("/contacts")

	contacts.Get("/all.json", controllers.GetContacts(cfg))
	contacts.Get("/:contact.vcf", controllers.GetCard(cfg))
	contacts.Get("/:contact.jpg", controllers.GetPhoto(cfg))

	app.Get("/images/logo.png", controllers.GetLogo(cfg.LogoURL))

	app.Use("/", filesystem.New(filesystem.Config{Root: rice.MustFindBox("../public").HTTPBox()}))

	fmt.Println(app.Listen(fmt.Sprintf(":%d", cfg.Port)))
}
