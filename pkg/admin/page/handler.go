package page

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taverok/lazyadmin/pkg/admin/resource"
	"github.com/taverok/lazyadmin/pkg/utils"
)

type Handler struct {
	Prefix  string
	Service *Service
}

func (it *Handler) Register(router *mux.Router) {
	for _, r := range it.Service.ResourceService.GetResources() {
		// List
		route := fmt.Sprintf("/%s/%s", it.Prefix, r.Table)
		slog.Info(fmt.Sprintf("GET endpoint for %s is %s", r.Table, route))
		router.HandleFunc(route, it.Get(r)).Methods(http.MethodGet)

		router.HandleFunc(route, it.Create(r)).Methods(http.MethodPost)
		router.HandleFunc(route, it.Update(r)).Methods(http.MethodPut)
	}

	route := fmt.Sprintf("%s/form/{resource}/{id}", it.Prefix)

	router.HandleFunc(route, it.UpdateForm()).Methods(http.MethodGet)
}

func (it *Handler) Get(p *resource.Resource) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		page := Page[Table]{}
		defer utils.HtmlTemplate("table").Execute(w, &page)

		page, err := it.Service.List(p)
		if err != nil {
			page.Error = err.Error()
			return
		}
	}
}

func (it *Handler) UpdateForm() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		page := Page[Form]{}
		defer utils.HtmlTemplate("form").Execute(w, &page)

		vars := mux.Vars(r)
		resourceName := vars["resource"]
		id := vars["id"]
		err := r.ParseForm()
		if err != nil {
			page.Error = err.Error()
			return
		}
		page, err = it.Service.Form(resourceName, map[string]string{"id": id})
		if err != nil {
			page.Error = err.Error()
			return
		}
	}
}

func (it *Handler) Update(p *resource.Resource) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		resourceName := vars["resource"]
		id := vars["id"]
		err := r.ParseForm()
		if err != nil {
			slog.Error(err.Error())
			return
		}
		it.Service.Update(resourceName, id, r.Form)
	}
}

func (it *Handler) Create(p *resource.Resource) func(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {

	}
	return handler
}
