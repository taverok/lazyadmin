package config

import (
	"os"

	"github.com/taverok/lazyadmin/pkg/admin/auth/provider"
	"github.com/taverok/lazyadmin/pkg/db"
	"gopkg.in/yaml.v3"
)

type Config struct {
	UrlPrefix string                   `yaml:"urlPrefix"`
	Port      int                      `yaml:"port"`
	Sources   map[string]db.DataSource `yaml:"sources"`
	Auth      provider.Config          `yaml:"auth"`
}

func NewConfig(path string) (*Config, error) {
	c := &Config{}
	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(configFile, c)
	if err != nil {
		return nil, err
	}

	return c, err
}
