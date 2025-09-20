package controllers

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"

	"github.com/hitalos/bina/internal/config"
	"github.com/hitalos/bina/internal/models"
)

func loadFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return io.ReadAll(resp.Body)
	}

	return nil, errors.New(http.StatusText(resp.StatusCode))
}

var logoBuf = []byte{}

// GetLogo return logo image
func GetLogo(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		if len(logoBuf) == 0 {
			if logoBuf, err = loadFromURL(url); err != nil {
				log.Println(err)
				return
			}
		}

		w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(url)))
		_, _ = w.Write(logoBuf)
	}
}

// GetPhoto return contact photos
func GetPhoto(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contact := chi.URLParam(r, "contact")
		entry := models.Entry{}
		if err := entry.GetByAccount(contact); err != nil {
			log.Println(err)
			return
		}
		if cfg.EnableGravatar && entry.Emails["mail"] != "" {
			hash := md5.Sum([]byte(entry.Emails["mail"]))
			w.Header().Set("Location", fmt.Sprintf("https://www.gravatar.com/avatar/%x", hash))
			w.WriteHeader(http.StatusTemporaryRedirect)
			return
		}
		photoBuf, err := loadFromURL(cfg.PhotosURL + entry.ID + ".jpg")
		if err != nil {
			if err.Error() == http.StatusText(http.StatusNotFound) {
				http.Redirect(w, r, "/images/default.png", http.StatusTemporaryRedirect)
			}
			return
		}

		if len(photoBuf) == 0 {
			http.Redirect(w, r, "/images/default.png", http.StatusTemporaryRedirect)
			return
		}

		w.Header().Set("Content-Type", "image/jpeg")
		_, _ = w.Write(photoBuf)
	}
}
