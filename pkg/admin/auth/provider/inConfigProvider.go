package provider

import (
	"strings"

	"github.com/taverok/lazyadmin/pkg/rest"
	"gopkg.in/yaml.v3"
)

type InConfigProvider struct {
	jwtSecret string
	users     []User
}

func (it *InConfigProvider) Name() string {
	return "inConfig"
}

func (it *InConfigProvider) SetCredentials(config Config) error {
	raw, err := yaml.Marshal(config.RawDetails)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(raw, &it.users)
}

func (it *InConfigProvider) Authenticate(name, pass string) (Principal, error) {
	for _, u := range it.users {
		if strings.ToLower(u.Name) == strings.ToLower(name) && u.Pass == pass {
			return Principal{
				Name: u.Name,
				Role: "admin",
			}, nil
		}
	}

	return Principal{}, rest.ErrNotFound
}

type User struct {
	Name string `yaml:"name"`
	Pass string `yaml:"pass"`
}
