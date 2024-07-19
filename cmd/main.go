package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/hitalos/bina/internal/config"
	"github.com/hitalos/bina/internal/controllers"
)

//go:embed public
var embeds embed.FS

func main() {
	configFilepath := flag.String("c", "config.yml", "Path of config file")
	flag.Parse()
	cfg := config.Load(*configFilepath)

	app := chi.NewRouter()

	setMiddlewares(app)
	setRoutes(app, cfg)
	listen(app, cfg.Port)
}

func setMiddlewares(app *chi.Mux) {
	app.Use(middleware.Compress(6))
	if os.Getenv("DEBUG") == "1" {
		app.Use(middleware.RealIP)
		app.Use(middleware.Logger)
	}
}

func setRoutes(app *chi.Mux, cfg *config.Config) {
	app.Route("/contacts", func(contacts chi.Router) {
		contacts.Get("/all.json", controllers.GetContacts(cfg))
		contacts.Get("/{contact}", controllers.GetCard(cfg))
		contacts.Get("/{contact}", controllers.GetPhoto(cfg))
	})

	app.Get("/images/logo.png", controllers.GetLogo(cfg.LogoURL))

	publicDir, _ := fs.Sub(embeds, "public")
	app.Handle("/*", http.FileServer(http.FS(publicDir)))
}

func listen(app chi.Router, port int) {
	s := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      app,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	fmt.Printf("Listening on port %d\n", port)
	log.Fatalln(s.ListenAndServe())
}
