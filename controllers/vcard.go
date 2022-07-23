package controllers

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/hitalos/bina/config"
	"github.com/hitalos/bina/models"
)

var (
	//go:embed vcard.tmpl
	cardTemplate string

	tmpl *template.Template
)

// GetCard handle request to vCard files
func GetCard(cfg *config.Config) http.HandlerFunc {
	if tmpl == nil {
		tmpl = template.Must(template.New("vcard").Parse(fmt.Sprintf(cardTemplate, cfg.LogoURL)))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		contact := chi.URLParam(r, "contact")
		entry := models.Entry{}
		if err := entry.GetByAccount(contact); err != nil {
			log.Println(err)
			return
		}

		if err := entry.AttachPhoto(cfg.PhotosURL + entry.ID + ".jpg"); err != nil {
			log.Println(err)
			return
		}

		created := time.Now().In(time.UTC).Format(time.RFC3339)
		data := struct {
			Contact models.Entry
			Host    string
			Created string
		}{entry, r.Host, created}

		w.Header().Set("Content-Type", "text/vcard; charset=utf-8")
		w.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s.vcf"`, entry.FullName))

		_ = tmpl.Execute(w, data)
	}
}
