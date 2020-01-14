package controllers

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-chi/chi"

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
func GetLogo(w http.ResponseWriter, r *http.Request) {
	var err error
	if len(logoBuf) == 0 {
		if logoBuf, err = loadFromURL(os.Getenv("LOGO_URL")); err != nil {
			errHandler(w, err)
			return
		}
	}
	w.Header().Set("Content-Type", "image/png")
	if _, err = w.Write(logoBuf); err != nil {
		errHandler(w, err)
	}
}

// GetPhoto return contact photos
func GetPhoto(w http.ResponseWriter, r *http.Request) {
	contact := chi.URLParam(r, "contact")
	entry := models.Entry{}
	if err := entry.GetByAccount(contact); err != nil {
		errHandler(w, err)
		return
	}
	if os.Getenv("ENABLE_GRAVATAR") == "true" && entry.Emails["mail"] != "" {
		hash := md5.Sum([]byte(entry.Emails["mail"]))
		w.Header().Set("Location", fmt.Sprintf("https://www.gravatar.com/avatar/%x", hash))
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	photoBuf, err := loadFromURL(os.Getenv("PHOTOS_URL") + entry.ID + ".jpg")
	if err != nil {
		errHandler(w, err)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	if _, err = w.Write(photoBuf); err != nil {
		errHandler(w, err)
	}
}
