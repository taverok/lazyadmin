package resource

import (
	"strings"

	database "github.com/taverok/lazyadmin/pkg/db"
)

type Resource struct {
	Table  string   `yaml:"table"`
	Fields []*Field `yaml:"fields"`
	// TODO: pagination
	// TODO: filters
}

func (it Resource) GetFields(op rune) []*Field {
	var fields []*Field
	for _, f := range it.Fields {
		if strings.ContainsRune(f.Ops, op) {
			fields = append(fields, f)
		}
	}

	return fields
}

type Field struct {
	//TODO: Roles []string `yaml:"roles"`
	Name  string `yaml:"name"`
	Width string `yaml:"width"`
	Ops   string `yaml:"ops"`
	Meta  database.FieldMeta
}

func (it *Field) SetDefaults(meta *database.FieldMeta) {
	it.Meta = *meta

	if it.Ops == "" {
		it.Ops = "CRU"
	}
	if it.Width == "" {
		it.Width = "150px"
	}
}

func (it Field) GetPxWidth() int {
	if it.Meta.IsPK && it.Meta.IsNum() {
		return 20
	}

	return 100
}
