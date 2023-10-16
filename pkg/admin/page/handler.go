package page

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taverok/lazyadmin/pkg/admin/config"
	"github.com/taverok/lazyadmin/pkg/admin/resource"
	"github.com/taverok/lazyadmin/pkg/admin/static"
)

type Handler struct {
	Router *mux.Router
	Config config.Config

	StaticService *static.Service
	Service       *Service
}

func (it *Handler) Register() {
	for _, p := range it.Service.ResourceService.GetResources() {
		// List
		route := fmt.Sprintf("%s/%s", it.Config.UrlPrefix, p.Table)
		slog.Info(fmt.Sprintf("GET endpoint for %s is %s", p.Table, route))

		it.Router.HandleFunc(route, it.Get(p)).Methods(http.MethodGet)
		it.Router.HandleFunc(route, it.Create(p)).Methods(http.MethodPost)
		it.Router.HandleFunc(route, it.Update(p)).Methods(http.MethodPut)
	}
}

func (it *Handler) Get(p *resource.Resource) func(w http.ResponseWriter, r *http.Request) {
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
	return handler
}

func (it *Handler) Create(p *resource.Resource) func(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {

	}
	return handler
}

func (it *Handler) Update(p *resource.Resource) func(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {

	}
	return handler
}
