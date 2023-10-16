package admin

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/taverok/lazyadmin/pkg/admin/config"
	"github.com/taverok/lazyadmin/pkg/admin/page"
	"github.com/taverok/lazyadmin/pkg/admin/resource"
	database "github.com/taverok/lazyadmin/pkg/db"
)

type App struct {
	Router *mux.Router
	Config *config.Config
}

func (it *App) Init() error {
	for sourceName, source := range it.Config.Sources {
		db, err := database.NewDb(source)
		if err != nil {
			return err
		}

		// repos
		pageRepo := &page.MysqlRepo{DB: db}

		// services
		resourceService := resource.NewService(db, sourceName)
		pageService := page.NewService(pageRepo, resourceService)

		// handlers
		pageHandler := page.Handler{
			Router:  it.Router,
			Prefix:  fmt.Sprintf("%s/%s", it.Config.UrlPrefix, sourceName),
			Service: pageService,
		}
		pageHandler.Register()
	}

	return nil
}
