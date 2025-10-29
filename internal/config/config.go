package config

import (
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config represents the configuration of this application
type Config struct {
	Port           int          `yaml:"port"`
	CacheDuration  int          `yaml:"cacheDuration"`
	EnableGravatar bool         `yaml:"enableGravatar"`
	PhotosURL      string       `yaml:"photosURL"`
	LogoURL        string       `yaml:"logoURL"`
	Providers      []Provider   `yaml:"providers"`
	CustomStyles   template.CSS `yaml:"customStyles,omitempty"`
}

// Provider represents the configuration of one information provider
type Provider struct {
	Type                  string            `yaml:"type"`
	IgnoreSSLVerification bool              `yaml:"ignoreSSLVerification"`
	Params                map[string]string `yaml:"params"`
	Fields                Fields            `yaml:"fields"`
}

// Fields defines the fields used in the provider
type Fields struct {
	Identifier string   `yaml:"identifier"`
	FullName   string   `yaml:"fullName"`
	Phones     []string `yaml:"phones"`
	Emails     []string `yaml:"emails"`
	Others     []string `yaml:"others"`
	Photo      string   `yaml:"photo"`
}

func (c *Config) setDefaultsOnEmpty() {
	if c.Port == 0 {
		c.Port = 8000
	}

	if c.CacheDuration == 0 {
		c.CacheDuration = 300
	}

	if len(c.Providers) > 0 {
		for i := range c.Providers {
			if _, ok := c.Providers[i].Params["schema"]; !ok {
				c.Providers[i].Params["schema"] = "ldap"
			}

			if _, ok := c.Providers[i].Params["port"]; !ok {
				c.Providers[i].Params["port"] = "389"
			}
		}
	}
}

// Load loads the configuration from file sistem
func Load(configFilepath string) *Config {
	c := new(Config)

	f, err := os.Open(filepath.Clean(configFilepath))
	if err != nil {
		if os.IsNotExist(err) {
			slog.Error(fmt.Sprintf("Crie um arquivo %q no formato do exemplo do projeto", configFilepath))
			os.Exit(1)
		}
		slog.Error(err.Error())
	}

	if err = yaml.NewDecoder(f).Decode(c); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if err := f.Close(); err != nil {
		slog.Error(err.Error())
	}

	c.setDefaultsOnEmpty()

	return c
}
