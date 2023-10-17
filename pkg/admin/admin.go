package admin

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/taverok/lazyadmin/pkg/admin/auth"
	"github.com/taverok/lazyadmin/pkg/admin/auth/provider"
	"github.com/taverok/lazyadmin/pkg/admin/config"
	"github.com/taverok/lazyadmin/pkg/admin/page"
	"github.com/taverok/lazyadmin/pkg/admin/resource"
	"github.com/taverok/lazyadmin/pkg/db"
)

type App struct {
	Router       *mux.Router
	Config       *config.Config
	AuthProvider provider.Provider
}

func (it *App) Init() error {
	err := it.initAuthProvider()
	if err != nil {
		return err
	}

	authService := auth.Service{Config: it.Config, Provider: it.AuthProvider}
	authHandler := auth.Handler{
		Prefix:  it.Config.UrlPrefix,
		Service: authService,
	}
	authHandler.Register(it.Router)

	for sourceName, source := range it.Config.Sources {
		db, err := db.NewDb(source)
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
			Prefix:  fmt.Sprintf("%s/%s", it.Config.UrlPrefix, sourceName),
			Service: pageService,
		}
		pageHandler.Register(it.Router)
	}

	return nil
}

func (it *App) initAuthProvider() error {
	if it.AuthProvider != nil {
		return nil
	}

	authResolver := provider.NewResolver()
	var p provider.Provider
	p, err := authResolver.Resolve(it.Config.Auth)
	if err != nil {
		return err
	}

	it.AuthProvider = p

	return nil
}
