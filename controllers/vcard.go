package controllers

import (
	_ "embed"
	"fmt"
	"text/template"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/hitalos/bina/config"
	"github.com/hitalos/bina/models"
)

var (
	//go:embed vcard.tmpl
	cardTemplate string

	tmpl *template.Template
)

// GetCard handle request to vCard files
func GetCard(cfg *config.Config) fiber.Handler {
	if tmpl == nil {
		tmpl = template.Must(template.New("vcard").Parse(fmt.Sprintf(cardTemplate, cfg.LogoURL)))
	}

	return func(c *fiber.Ctx) error {
		contact := c.Params("contact")
		entry := models.Entry{}
		if err := entry.GetByAccount(contact); err != nil {
			return err
		}

		if err := entry.AttachPhoto(cfg.PhotosURL + entry.ID + ".jpg"); err != nil {
			return err
		}

		created := time.Now().In(time.UTC).Format(time.RFC3339)
		data := struct {
			Contact models.Entry
			Host    string
			Created string
		}{entry, c.Hostname(), created}

		c.Set("Content-Type", "text/vcard; charset=utf-8")
		c.Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s.vcf"`, entry.FullName))

		return tmpl.Execute(c, data)
	}
}
