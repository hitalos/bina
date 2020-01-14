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

func validCache() bool {
	cacheDuration, err := strconv.Atoi(os.Getenv("CACHE_DURATION"))
	if err != nil {
		cacheDuration = 300
	}
	return lastCached.Add(time.Duration(cacheDuration)*time.Second).Unix() > time.Now().Unix()
}

var contactsJSON []byte
var lastCached time.Time

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
}

// GetContacts return all contacts in JSON format
func GetContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	if validCache() {
		w.Write(contactsJSON)
		return
	}

	list, err := models.GetContacts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if contactsJSON, err = json.Marshal(list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	lastCached = time.Now()
	w.Write(contactsJSON)
}
