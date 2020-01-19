package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config represents the configuration of this application
type Config struct {
	Port           int        `yaml:"port"`
	CacheDuration  int        `yaml:"cacheDuration"`
	EnableGravatar bool       `yaml:"enableGravatar"`
	PhotosURL      string     `yaml:"photosURL"`
	LogoURL        string     `yaml:"logoURL"`
	Providers      []Provider `yaml:"providers"`
}

// Provider represents the configuration of one information provider
type Provider struct {
	Type   string            `yaml:"type"`
	Params map[string]string `yaml:"params"`
	Fields Fields            `yaml:"fields"`
}

// Fields defines the fields used in the provider
type Fields struct {
	Identifier string   `yaml:"identifier"`
	FullName   string   `yaml:"fullName"`
	Phones     []string `yaml:"phones"`
	Emails     []string `yaml:"emails"`
	Others     []string `yaml:"others"`
}

func (c *Config) setDefaultsOnEmpty() {
	if c.Port == 0 {
		c.Port = 8000
	}
	if c.CacheDuration == 0 {
		c.CacheDuration = 300
	}
}

// Load loads the configuration from file sistem
func Load() *Config {
	c := new(Config)
	f, err := os.Open("config.yml")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalln("Crie um arquivo 'config.yml' no formato do exemplo do projeto")
		}
		log.Fatalln(err)
	}
	defer f.Close()
	yaml.NewDecoder(f).Decode(c)

	c.setDefaultsOnEmpty()
	return c
}
