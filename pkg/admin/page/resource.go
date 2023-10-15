package page

import (
	"strings"
)

type Resource struct {
	Name   string   `yaml:"name"`
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
	//Roles []string `yaml:"roles"`
	Name  string    `yaml:"name"`
	Width string    `yaml:"width"`
	Ops   string    `yaml:"ops"`
	Type  FieldType `yaml:"type"`
}

func (it *Field) SetDefaults() {
	if it.Ops == "" {
		it.Ops = "CRU"
	}
	if it.Type == "" {
		it.Type = "text"
	}
	if it.Width == "" {
		it.Width = "150px"
	}
}

type FieldType string
