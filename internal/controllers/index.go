package controllers

import (
	"html/template"
	"io"
	"net/http"

	"github.com/hitalos/bina/cmd/public"
	"github.com/hitalos/bina/internal/config"
)

func Index(cfg *config.Config) http.HandlerFunc {
	f, _ := public.FS.Open("index.html")
	bs, _ := io.ReadAll(f)
	idxTmpl, _ := template.New("").Parse(string(bs))

	return func(w http.ResponseWriter, _ *http.Request) {
		data := struct {
			CustomStyles template.CSS
		}{cfg.CustomStyles}
		_ = idxTmpl.Execute(w, data)
	}
}
