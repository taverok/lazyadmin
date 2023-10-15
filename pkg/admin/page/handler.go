package page

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taverok/lazyadmin/pkg/admin/config"
	"github.com/taverok/lazyadmin/pkg/admin/static"
)

type Handler struct {
	Router *mux.Router
	Config config.Config

	StaticService *static.Service
	Service       *Service
}

func (it *Handler) Register() {
	for _, p := range it.Service.Resources {
		// List
		route := fmt.Sprintf("%s/%s", it.Config.UrlPrefix, p.Name)
		slog.Info(fmt.Sprintf("GET endpoint for %s is %s", p.Name, route))

		handler := func(w http.ResponseWriter, r *http.Request) {
			page, err := it.Service.List(p)
			if err != nil {
				slog.Error(err.Error())
			}

			err = it.StaticService.Template("table").Execute(w, page)
			if err != nil {
				slog.Error(err.Error())
			}
		}
		it.Router.HandleFunc(route, handler).Methods(http.MethodGet)

		// Update

		// Create

	}
}
