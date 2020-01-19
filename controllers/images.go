package controllers

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/hitalos/bina/config"
	"github.com/hitalos/bina/models"
)

func loadFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New("404 Not Found")
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

var logoBuf = []byte{}

// GetLogo return logo image
func GetLogo(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		if len(logoBuf) == 0 {
			if logoBuf, err = loadFromURL(url); err != nil {
				errHandler(w, err)
				return
			}
		}
		w.Header().Set("Content-Type", "image/png")
		if _, err = w.Write(logoBuf); err != nil {
			errHandler(w, err)
		}
	}
}

// GetPhoto return contact photos
func GetPhoto(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contact := chi.URLParam(r, "contact")
		entry := models.Entry{}
		if err := entry.GetByAccount(contact); err != nil {
			errHandler(w, err)
			return
		}
		if c.EnableGravatar && entry.Emails["mail"] != "" {
			hash := md5.Sum([]byte(entry.Emails["mail"]))
			w.Header().Set("Location", fmt.Sprintf("https://www.gravatar.com/avatar/%x", hash))
			w.WriteHeader(http.StatusTemporaryRedirect)
			return
		}
		photoBuf, err := loadFromURL(c.PhotosURL + entry.ID + ".jpg")
		if err != nil {
			errHandler(w, err)
			return
		}

		w.Header().Set("Content-Type", "image/jpeg")
		if _, err = w.Write(photoBuf); err != nil {
			errHandler(w, err)
		}
	}
}
