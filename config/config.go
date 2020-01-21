package config

import (
	"fmt"
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
func Load(configFilepath string) *Config {
	c := new(Config)
	f, err := os.Open(configFilepath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Crie um arquivo %q no formato do exemplo do projeto\n", configFilepath)
			os.Exit(1)
		}
		fmt.Println(err)
	}
	defer f.Close()
	if err = yaml.NewDecoder(f).Decode(c); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c.setDefaultsOnEmpty()
	return c
}
