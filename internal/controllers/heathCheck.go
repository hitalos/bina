package controllers

import (
	"net/http"

	"github.com/hitalos/bina/internal/config"
	"github.com/hitalos/bina/internal/services/ldap"
)

func HealthCheck(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, p := range cfg.Providers {
			conn, err := ldap.NewLdapConnection(p)
			if err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				_, _ = w.Write([]byte("unavailable"))

				return
			}
			_ = conn.Close()
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}
}
