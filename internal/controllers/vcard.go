package controllers

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"text/template"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/hitalos/bina/internal/config"
	"github.com/hitalos/bina/internal/models"
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
			slog.Error(err.Error())

			return
		}

		if entry.Photo == "" {
			if err := entry.AttachPhoto(cfg.PhotosURL+entry.ID+".jpg", r.Context()); err != nil {
				slog.Error(err.Error())

				return
			}
		}

		bs, _ := base64.StdEncoding.DecodeString(entry.Photo)
		imgType := "JPG"
		switch http.DetectContentType(bs[0:512]) {
		case "image/png":
			imgType = "PNG"
		case "image/gif":
			imgType = "GIF"
		case "image/webp":
			imgType = "WEBP"
		}

		created := time.Now().In(time.UTC).Format(time.RFC3339)
		data := struct {
			Contact models.Entry
			ImgType string
			Host    string
			Created string
		}{entry, imgType, r.Host, created}

		w.Header().Set("Content-Type", "text/vcard; charset=utf-8")
		w.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s.vcf"`, url.PathEscape(entry.FullName)))

		_ = tmpl.Execute(w, data)
	}
}
