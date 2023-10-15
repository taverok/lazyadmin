package config

import database "github.com/taverok/lazyadmin/pkg/db"

type Config struct {
	UrlPrefix    string
	Port         int
	ResourcePath string
	Db           database.DataSource
	TemplateName string
}

func DefaultConfig() Config {
	// TODO: вынести в файл

	ds := database.DataSource{
		Host:    "localhost",
		Port:    3309,
		User:    "test",
		Name:    "test",
		SSLMode: "disable",
		Pass:    "test",
	}

	return Config{
		UrlPrefix:    "/admin",
		Port:         8080,
		ResourcePath: "resources",
		TemplateName: "lte",
		Db:           ds,
	}
}
