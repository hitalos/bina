package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	apm "go.elastic.co/apm/module/apmchi"

	"github.com/hitalos/bina/config"
	"github.com/hitalos/bina/controllers"
)

func main() {
	configFilepath := flag.String("c", "config.yml", "Path of config file")
	flag.Parse()
	c := config.Load(*configFilepath)
	r := chi.NewRouter()
	r.Use(middleware.RealIP, apm.Middleware(), middleware.Compress(6))
	if os.Getenv("DEBUG") == "1" {
		r.Use(middleware.Logger)
	}

	r.Route("/contacts", func(r chi.Router) {
		r.Get("/all.json", controllers.GetContacts(c))
		r.Get("/{contact:[a-z]+}.vcf", controllers.GetCard(c))
		r.Get("/{contact:[a-z]+}.jpg", controllers.GetPhoto(c))
	})
	r.Get("/images/logo.png", controllers.GetLogo(c.LogoURL))

	r.Handle("/*", http.FileServer(rice.MustFindBox("../public").HTTPBox()))

	fmt.Printf("Listening on port :%d\n", c.Port)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", c.Port), r))
}
