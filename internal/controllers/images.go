package controllers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/hitalos/bina/internal/config"
	"github.com/hitalos/bina/internal/models"
)

var (
	ErrNonOKHTTPStatus = errors.New("non-OK HTTP status")
)

func loadFromURL(url string, ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %s", ErrNonOKHTTPStatus, http.StatusText(resp.StatusCode))
	}

	return io.ReadAll(resp.Body)
}

var logoBuf = []byte{}

// GetLogo return logo image
func GetLogo(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		if len(logoBuf) == 0 {
			if logoBuf, err = loadFromURL(url, r.Context()); err != nil {
				slog.Error(err.Error())

				return
			}
		}

		w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(url)))
		_, _ = w.Write(logoBuf)
	}
}

// GetPhoto return contact photos
func GetPhoto(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contact := chi.URLParam(r, "contact")
		entry := models.Entry{}
		if err := entry.GetByAccount(contact); err != nil {
			slog.Error(err.Error())

			return
		}
		if cfg.EnableGravatar && entry.Emails["mail"] != "" {
			hash := sha256.Sum256([]byte(strings.TrimSpace(entry.Emails["mail"])))
			w.Header().Set("Location", "https://www.gravatar.com/avatar/"+hex.EncodeToString(hash[:]))
			w.WriteHeader(http.StatusTemporaryRedirect)

			return
		}
		photoBuf, err := loadFromURL(cfg.PhotosURL+entry.ID+".jpg", r.Context())
		if err != nil {
			if errors.Is(err, ErrNonOKHTTPStatus) {
				http.Redirect(w, r, "/images/default.png", http.StatusTemporaryRedirect)
			}

			return
		}

		if len(photoBuf) == 0 {
			http.Redirect(w, r, "/images/default.png", http.StatusTemporaryRedirect)

			return
		}

		w.Header().Set("Content-Type", "image/jpeg")
		_, _ = w.Write(photoBuf)
	}
}
