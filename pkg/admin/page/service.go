package page

import (
	"net/url"

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

func (it *Service) List(r *resource.Resource) (Page[Table], error) {
	page := Page[Table]{
		Name:    r.Table,
		Content: Table{},
		Menu:    it.sidebar,
	}

	fields := r.GetFields('R')
	data, err := it.repo.GetAll(fields, r.Table)
	if err != nil {
		return page, err
	}
	page.Content = Table{
		Fields: fields,
		Data:   data,
	}

	return page, nil
}

func (it *Service) Form(name string, where map[string]string) (Page[Form], error) {
	page := Page[Form]{
		Name:    name,
		Content: Form{},
		Menu:    it.sidebar,
	}

	r, err := it.ResourceService.ResourceByName(name)
	if err != nil {
		return page, err
	}

	updateFields := r.GetFields('U')

	data, err := it.repo.GetById(r.Table, updateFields, where)
	if err != nil {
		return page, err
	}

	page.Content = NewForm(updateFields, data)

	return page, nil
}

func (it *Service) Update(resourceName string, id string, form url.Values) error {
	return nil
}
