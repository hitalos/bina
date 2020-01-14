package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	rice "github.com/GeertJohan/go.rice"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	apm "go.elastic.co/apm/module/apmchi"

	"github.com/hitalos/bina/controllers"
)

func main() {
	godotenv.Load()
	r := chi.NewRouter()
	r.Use(middleware.RealIP, apm.Middleware(), middleware.DefaultCompress)
	if os.Getenv("DEBUG") == "1" {
		r.Use(middleware.Logger)
	}

	r.Route("/contacts", func(r chi.Router) {
		r.Get("/all.json", controllers.GetContacts)
		r.Get("/{contact:[a-z]+}.vcf", controllers.GetCard)
		r.Get("/{contact:[a-z]+}.jpg", controllers.GetPhoto)
	})
	r.Get("/images/logo.png", controllers.GetLogo)

	r.Handle("/*", http.FileServer(rice.MustFindBox("../public").HTTPBox()))

	port := os.Getenv("PORT")
	if _, err := strconv.Atoi(port); err != nil {
		port = "8000"
	}
	fmt.Println("Listening on port", port)
	fmt.Println(http.ListenAndServe(":"+port, r))
}
