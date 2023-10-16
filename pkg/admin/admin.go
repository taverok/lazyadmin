package admin

import (
	"github.com/gorilla/mux"
	"github.com/taverok/lazyadmin/pkg/admin/config"
	"github.com/taverok/lazyadmin/pkg/admin/page"
	"github.com/taverok/lazyadmin/pkg/admin/resource"
	"github.com/taverok/lazyadmin/pkg/admin/static"
	database "github.com/taverok/lazyadmin/pkg/db"
)

type App struct {
	Router *mux.Router
	Config config.Config
}

func (it *App) Init() error {
	db, err := database.NewMysqlDb(it.Config.Db)
	if err != nil {
		return err
	}

	// repos
	pageRepo := &page.MysqlRepo{DB: db}

	// services
	staticService := &static.Service{Config: it.Config}
	resourceService := resource.NewService(db, staticService)
	pageService := page.NewService(pageRepo, resourceService)

	// handlers
	pageHandler := page.Handler{
		Router:        it.Router,
		Config:        it.Config,
		StaticService: staticService,
		Service:       pageService,
	}
	pageHandler.Register()

	return nil
}
