package page

import (
	"github.com/taverok/lazyadmin/pkg/admin/resource"
)

type Service struct {
	repo            *MysqlRepo
	ResourceService *resource.Service

	sidebar    []string
	nameToType []string
}

func NewService(repo *MysqlRepo, resourceService *resource.Service) *Service {
	service := Service{
		repo:            repo,
		ResourceService: resourceService,
	}

	for _, r := range resourceService.GetResources() {
		service.sidebar = append(service.sidebar, r.Table)
	}

	return &service
}

func (it *Service) List(p *resource.Resource) (Page[Table], error) {
	fields := p.GetFields('R')
	data, err := it.repo.GetAll(fields, p.Table)
	if err != nil {
		return Page[Table]{}, err
	}

	return it.toTablePage(p, data, fields), nil
}

func (it *Service) toTablePage(p *resource.Resource, data [][]any, fields []*resource.Field) Page[Table] {
	return Page[Table]{
		Name: p.Table,
		Content: Table{
			Fields: fields,
			Data:   data,
		},
		Menu: it.sidebar,
	}
}
