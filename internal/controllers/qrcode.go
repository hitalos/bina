package controllers

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log/slog"
	"net/http"
	"strings"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/go-chi/chi/v5"

	"github.com/hitalos/bina/internal/config"
	"github.com/hitalos/bina/internal/models"
)

// GetQRCode handle request to QR Code images
func GetQRCode(cfg *config.Config) http.HandlerFunc {
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

		entry.Photo = ""

		data := struct {
			Contact models.Entry
			ImgType string
			Host    string
			Created string
		}{entry, "", "", ""}

		buf := new(bytes.Buffer)
		if err := tmpl.Execute(buf, data); err != nil {
			slog.Error(err.Error())

			return
		}

		qrcode, err := qr.Encode(strings.TrimSpace(buf.String()), qr.L, qr.Auto)
		if err != nil {
			slog.Error(err.Error())

			return
		}

		blockSize := 4
		size := qrcode.Bounds().Max.X * blockSize
		if qrcode, err = barcode.Scale(qrcode, size, size); err != nil {
			slog.Error(err.Error())

			return
		}

		bColor := color.RGBA{255, 255, 255, 255}
		borderedImage := image.NewRGBA(image.Rect(0, 0, qrcode.Bounds().Dx()+4*blockSize, qrcode.Bounds().Dy()+4*blockSize))
		draw.Draw(
			borderedImage,
			borderedImage.Bounds(),
			&image.Uniform{C: bColor},
			image.Point{},
			draw.Src)
		draw.Draw(
			borderedImage,
			image.Rect(2*blockSize, 2*blockSize, qrcode.Bounds().Dx()+2*blockSize, qrcode.Bounds().Dy()+2*blockSize),
			qrcode,
			image.Point{},
			draw.Src)

		if err = png.Encode(w, borderedImage); err != nil {
			slog.Error(err.Error())

			return
		}
	}
}
