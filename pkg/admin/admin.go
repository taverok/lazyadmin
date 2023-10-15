package admin

import (
	"github.com/gorilla/mux"
	"github.com/taverok/lazyadmin/pkg/admin/config"
	"github.com/taverok/lazyadmin/pkg/admin/page"
	"github.com/taverok/lazyadmin/pkg/admin/static"
	database "github.com/taverok/lazyadmin/pkg/db"
)

type App struct {
	Router        *mux.Router
	Config        config.Config
	StaticService *static.Service

	resources []*page.Resource
}

func (it *App) Init() error {
	db, err := database.NewMysqlDb(it.Config.Db)
	if err != nil {
		return err
	}

	it.StaticService = &static.Service{Config: it.Config}

	err = it.initResources()
	if err != nil {
		return err
	}

	// repos
	pageRepo := &page.MysqlRepo{DB: db}

	// services
	pageService := page.NewService(it.Router, pageRepo, it.resources)

	// handlers
	pageHandler := page.Handler{
		Router:        it.Router,
		Config:        it.Config,
		StaticService: it.StaticService,
		Service:       pageService,
	}
	pageHandler.Register()

	return nil
}

func (it *App) initResources() error {
	err := it.StaticService.MapYml("resources", &it.resources)
	if err != nil {
		return err
	}

	for _, resource := range it.resources {
		for _, f := range resource.Fields {
			f.SetDefaults()
		}
	}

	return nil
}
