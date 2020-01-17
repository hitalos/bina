package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/hitalos/bina/models"
)

var (
	cacheDuration int
	contactsJSON  []byte
	lastCached    time.Time
)

func validCache() bool {
	return lastCached.Add(time.Duration(cacheDuration)*time.Second).Unix() > time.Now().Unix()
}

func init() {
	list, err := models.GetContacts()
	if err != nil {
		fmt.Println(err)
		return
	}
	if contactsJSON, err = json.Marshal(list); err != nil {
		fmt.Println(err)
		return
	}

	lastCached = time.Now()
	cacheDuration, err = strconv.Atoi(os.Getenv("CACHE_DURATION"))
	if err != nil {
		cacheDuration = 300
	}
}

// GetContacts return all contacts in JSON format
func GetContacts(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("If-Modified-Since") != "" {
		browserCacheTime, err := time.Parse(time.RFC1123, r.Header.Get("If-Modified-Since"))
		maxValidCache := lastCached.Add(time.Duration(cacheDuration) * time.Second).Unix()
		if err == nil && browserCacheTime.Unix() < maxValidCache {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}
	w.Header().Set("Content-Type", "aplication/json; charset=utf-8")
	w.Header().Set("Last-Modified", lastCached.Format(time.RFC1123))
	w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", cacheDuration))
	if validCache() {
		w.Write(contactsJSON)
		return
	}

	list, err := models.GetContacts()
	if err != nil {
		errHandler(w, err)
		return
	}
	if contactsJSON, err = json.Marshal(list); err != nil {
		errHandler(w, err)
		return
	}
	lastCached = time.Now()
	w.Write(contactsJSON)
}
