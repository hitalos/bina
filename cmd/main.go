package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/hitalos/bina/config"
	"github.com/hitalos/bina/controllers"
)

//go:embed public
var embeds embed.FS

func main() {
	configFilepath := flag.String("c", "config.yml", "Path of config file")
	flag.Parse()
	cfg := config.Load(*configFilepath)

	app := chi.NewRouter()
	app.Use(middleware.Compress(6))
	if os.Getenv("DEBUG") == "1" {
		app.Use(middleware.Logger)
	}

	app.Route("/contacts", func(contacts chi.Router) {
		contacts.Get("/all.json", controllers.GetContacts(cfg))
		contacts.Get("/{contact}", controllers.GetCard(cfg))
		contacts.Get("/{contact}", controllers.GetPhoto(cfg))
	})

	app.Get("/images/logo.png", controllers.GetLogo(cfg.LogoURL))

	publicDir, _ := fs.Sub(embeds, "public")
	app.Handle("/*", http.FileServer(http.FS(publicDir)))

	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), app))
}
