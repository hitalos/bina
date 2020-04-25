package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hitalos/bina/config"
	"github.com/hitalos/bina/models"
)

var (
	contactsJSON []byte
	lastCached   time.Time
)

func validCache(duration int) bool {
	return lastCached.Add(time.Duration(duration)*time.Second).Unix() > time.Now().Unix()
}

func load(p []config.Provider) {
	list, err := models.GetContacts(p)
	if err != nil {
		log.Println(err)
		return
	}
	if contactsJSON, err = json.Marshal(list); err != nil {
		log.Println(err)
		return
	}

	lastCached = time.Now()
}

// GetContacts return all contacts in JSON format
func GetContacts(c *config.Config) http.HandlerFunc {
	load(c.Providers)
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("If-Modified-Since") != "" {
			browserCacheTime, err := time.Parse(time.RFC1123, r.Header.Get("If-Modified-Since"))
			maxValidCache := lastCached.Add(time.Duration(c.CacheDuration) * time.Second).Unix()
			if err == nil && browserCacheTime.Unix() < maxValidCache {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Last-Modified", lastCached.Format(time.RFC1123))
		w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", c.CacheDuration))
		if validCache(c.CacheDuration) {
			w.Write(contactsJSON)
			return
		}

		load(c.Providers)
		w.Write(contactsJSON)
	}
}
