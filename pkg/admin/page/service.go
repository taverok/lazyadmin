package page

import (
	"github.com/gorilla/mux"
)

type Service struct {
	repo      *MysqlRepo
	Resources []*Resource

	sidebar    []string
	nameToType []string
}

func NewService(handler *mux.Router, repo *MysqlRepo, resources []*Resource) *Service {
	service := Service{
		repo:      repo,
		Resources: resources,
	}

	for _, r := range resources {
		service.sidebar = append(service.sidebar, r.Name)
	}

	return &service
}

func (it *Service) List(p *Resource) (Page, error) {
	fields := p.GetFields('R')
	data, err := it.repo.GetAll(fields, p.Name)
	if err != nil {
		return Page{}, err
	}

	return it.toTablePage(p, data, fields), nil
}

func (it *Service) toTablePage(p *Resource, data [][]any, fields []*Field) Page {

	return Page{
		Name: p.Name,
		Content: &Table{
			Fields: fields,
			Data:   data,
		},
		Menu: it.sidebar,
	}
}
