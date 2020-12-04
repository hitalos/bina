package controllers

import (
	"log"
	"os"
	"text/template"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/hitalos/bina/config"
	"github.com/hitalos/bina/models"
)

// GetCard handle request to vCard files
func GetCard(cfg *config.Config) fiber.Handler {
	cardTemplate := `BEGIN:VCARD
VERSION:3.0
FN;CHARSET=UTF-8:{{ .Contact.FullName }}
N;CHARSET=UTF-8:{{ .Contact.LastName }};{{ .Contact.FirstName }};;;
NICKNAME;CHARSET=UTF-8:{{ .Contact.ID }}
EMAIL;CHARSET=UTF-8;type=HOME,INTERNET:{{ index .Contact.Emails "mail" }}
LOGO;TYPE=PNG:` + os.Getenv("LOGO_URL") + `
PHOTO;ENCODING=b;TYPE=JPG:{{ .Contact.Photo }}
TEL;TYPE=CELL:{{ index .Contact.Phones "mobile" }}
TEL;TYPE=WORK,VOICE:{{ index .Contact.Phones "ipPhone" }}
TITLE;CHARSET=UTF-8:{{ .Contact.Title }}
ROLE;CHARSET=UTF-8:{{ .Contact.Title }}
NOTE;CHARSET=UTF-8:{{ .Contact.Department }} - {{ .Contact.PhysicalDeliveryOfficeName }}
SOURCE;CHARSET=UTF-8:http://{{ .Host }}/contacts/{{ .Contact.ID }}.vcf
REV:{{ .Created }}
END:VCARD`

	tmpl, err := template.New("vcard").Parse(cardTemplate)
	if err != nil {
		log.Fatalln(err)
	}
	return func(c *fiber.Ctx) error {
		contact := c.Params("contact")
		entry := models.Entry{}
		if err := entry.GetByAccount(contact); err != nil {
			return err
		}

		if err := entry.AttachPhoto(os.Getenv("PHOTOS_URL") + entry.ID + ".jpg"); err != nil {
			return err
		}

		created := time.Now().In(time.UTC).Format(time.RFC3339)
		data := struct {
			Contact models.Entry
			Host    string
			Created string
		}{
			entry,
			c.Hostname(),
			created}
		c.Set("Content-Type", "text/vcard; charset=utf-8")
		c.Set("Content-Disposition", "inline; filename=\""+entry.FullName+".vcf\"")
		return tmpl.Execute(c, data)
	}
}
