package resource

import (
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
	database "github.com/taverok/lazyadmin/pkg/db"
	"github.com/taverok/lazyadmin/pkg/rest"
	"github.com/taverok/lazyadmin/pkg/utils"
)

type Service struct {
	db             *sqlx.DB
	resources      []*Resource
	nameToResource map[string]*Resource
}

func NewService(db *sqlx.DB, sourceName string) *Service {
	service := &Service{
		db:             db,
		nameToResource: map[string]*Resource{},
	}

	err := service.init(sourceName)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return service
}

func (it *Service) init(sourceName string) error {
	err := utils.MapYml("config/resources.yml", &it.resources)
	if err != nil {
		return err
	}

	for _, r := range it.resources {
		if r.Source != sourceName {
			continue
		}
		metas, err := database.AnalyzeTable(it.db.DB, r.Table)
		if err != nil {
			return err
		}

		for _, f := range r.Fields {
			meta := metas[f.Name]
			f.SetDefaults(meta)
		}

		it.nameToResource[r.Table] = r
	}

	return nil
}

func (it *Service) ResourceByName(name string) (*Resource, error) {
	resource, ok := it.nameToResource[name]
	if !ok {
		return nil, rest.ErrNotFound
	}

	return resource, nil
}

func (it *Service) GetResources() []*Resource {
	return it.resources
}

func FieldNames(fields []*Field) []string {
	var names []string
	for _, f := range fields {
		names = append(names, f.Name)
	}

	return names
}
