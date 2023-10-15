package static

import (
	"fmt"
	"html/template"
	"os"

	"github.com/Masterminds/sprig/v3"
	"github.com/taverok/lazyadmin/pkg/admin/config"
	"gopkg.in/yaml.v3"
)

type Service struct {
	Config config.Config
}

func (it Service) GetBytes(name string) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("%s/%s", it.Config.ResourcePath, name))
}

func (it Service) MapYml(file string, target any) error {
	content, err := it.GetBytes(fmt.Sprintf("%s.yml", file))
	if err != nil {
		return err
	}

	return yaml.Unmarshal(content, target)
}

func (it Service) Template(name string) *template.Template {
	current := fmt.Sprintf("%s/templates/%s/%s.gohtml", it.Config.ResourcePath, it.Config.TemplateName, name)
	base := fmt.Sprintf("%s/templates/%s/base.gohtml", it.Config.ResourcePath, it.Config.TemplateName)
	return template.Must(
		template.New("base").
			Funcs(sprig.GenericFuncMap()).
			ParseFiles(base, current),
	)
}
