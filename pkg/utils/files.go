package utils

import (
	"fmt"
	"html/template"
	"os"

	"github.com/Masterminds/sprig/v3"
	"gopkg.in/yaml.v3"
)

func MapYml(file string, target any) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(content, target)
}

func HtmlTemplate(name string) *template.Template {
	current := fmt.Sprintf("resources/templates/lte/%s.gohtml", name)
	base := fmt.Sprintf("resources/templates/lte/base.gohtml")
	return template.Must(
		template.New("base").
			Funcs(sprig.GenericFuncMap()).
			ParseFiles(base, current),
	)
}
