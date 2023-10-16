package resource

import (
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/taverok/lazyadmin/pkg/admin/static"
	database "github.com/taverok/lazyadmin/pkg/db"
)

type Service struct {
	db            *sqlx.DB
	staticService *static.Service
	resources     []*Resource
}

func NewService(db *sqlx.DB, staticService *static.Service) *Service {
	service := &Service{
		db:            db,
		staticService: staticService,
	}

	err := service.init()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return service
}

func (it *Service) init() error {
	err := it.staticService.MapYml("resources", &it.resources)
	if err != nil {
		return err
	}

	for _, r := range it.resources {
		metas, err := database.AnalyzeTable(it.db.DB, r.Table)
		if err != nil {
			return err
		}

		for _, f := range r.Fields {
			meta := metas[f.Name]
			f.SetDefaults(meta)
		}
	}

	return nil
}

func (it *Service) GetResources() []*Resource {
	return it.resources
}
