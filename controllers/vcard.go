package controllers

import (
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/go-chi/chi"

	"github.com/hitalos/bina/config"
	"github.com/hitalos/bina/models"
)

// GetCard handle request to vCard files
func GetCard(c *config.Config) http.HandlerFunc {
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
	return func(w http.ResponseWriter, r *http.Request) {
		contact := chi.URLParam(r, "contact")
		entry := models.Entry{}
		if err := entry.GetByAccount(contact); err != nil {
			errHandler(w, err)
			return
		}

		if err := entry.AttachPhoto(os.Getenv("PHOTOS_URL") + entry.ID + ".jpg"); err != nil {
			errHandler(w, err)
			return
		}

		created := time.Now().In(time.UTC).Format(time.RFC3339)
		data := struct {
			Contact models.Entry
			Host    string
			Created string
		}{
			entry,
			r.Host,
			created}
		w.Header().Set("Content-Type", "text/vcard; charset=utf-8")
		w.Header().Set("Content-Disposition", "inline; filename=\""+entry.FullName+".vcf\"")
		if err := tmpl.Execute(w, data); err != nil {
			errHandler(w, err)
		}
	}
}
