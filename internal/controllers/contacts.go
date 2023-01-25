package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hitalos/bina/internal/config"
	"github.com/hitalos/bina/internal/models"
)

var (
	contactsJSON []byte
	lastCached   time.Time
)

func validCache(duration int) bool {
	return lastCached.Add(time.Duration(duration)*time.Second).Unix() > time.Now().Unix()
}

func loadContacts(p []config.Provider) error {
	list, err := models.GetContacts(p)
	if err != nil {
		return err
	}
	if contactsJSON, err = json.Marshal(list); err != nil {
		return err
	}

	lastCached = time.Now()

	return nil
}

// GetContacts return all contacts in JSON format
func GetContacts(cfg *config.Config) http.HandlerFunc {
	if err := loadContacts(cfg.Providers); err != nil {
		log.Println(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		since := r.Header.Get("If-Modified-Since")
		if len(contactsJSON) != 0 && since != "" {
			browserCacheTime, err := time.Parse(time.RFC1123, since)
			maxValidCache := lastCached.Add(time.Duration(cfg.CacheDuration) * time.Second).Unix()
			if err == nil && browserCacheTime.Unix() < maxValidCache {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Header().Add("Last-Modified", lastCached.Format(time.RFC1123))
		w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d", cfg.CacheDuration))
		if len(contactsJSON) != 0 && validCache(cfg.CacheDuration) {
			_, _ = w.Write(contactsJSON)
			return
		}

		if err := loadContacts(cfg.Providers); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("Serviço indisponível"))

			return
		}

		_, _ = w.Write(contactsJSON)
	}
}
