package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
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

func loadContacts(p []config.Provider) {
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
func GetContacts(cfg *config.Config) fiber.Handler {
	loadContacts(cfg.Providers)
	return func(c *fiber.Ctx) error {
		if len(contactsJSON) != 0 && c.Get("If-Modified-Since") != "" {
			browserCacheTime, err := time.Parse(time.RFC1123, c.Get("If-Modified-Since"))
			maxValidCache := lastCached.Add(time.Duration(cfg.CacheDuration) * time.Second).Unix()
			if err == nil && browserCacheTime.Unix() < maxValidCache {
				return c.SendStatus(fiber.StatusNotModified)
			}
		}
		c.Set("Content-Type", fiber.MIMEApplicationJSONCharsetUTF8)
		c.Set("Last-Modified", lastCached.Format(time.RFC1123))
		c.Set("Cache-Control", fmt.Sprintf("max-age=%d", cfg.CacheDuration))
		if len(contactsJSON) != 0 && validCache(cfg.CacheDuration) {
			return c.Send(contactsJSON)
		}

		loadContacts(cfg.Providers)
		return c.Send(contactsJSON)
	}
}
