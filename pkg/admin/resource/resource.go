package resource

import (
	"sort"
	"strings"
	"unicode"

	"github.com/taverok/lazyadmin/pkg/db"
)

type Resource struct {
	Table  string   `yaml:"table"`
	Source string   `yaml:"source"`
	Fields []*Field `yaml:"fields"`
	// TODO: pagination
	// TODO: filters
}

func (it *Resource) SetDefaults(metas map[string]*db.FieldMeta) {
	if it.Source == "" {
		it.Source = "main"
	}

	for _, f := range it.Fields {
		meta := metas[f.Name]
		f.SetDefaults(meta)
	}

	it.setDefaultFields(metas)
}

func (it *Resource) setDefaultFields(metas map[string]*db.FieldMeta) {
	if len(it.Fields) > 0 {
		return
	}

	for _, meta := range metas {
		ops := "CRU"
		if meta.IsPK {
			ops = "R"
		}
		it.Fields = append(it.Fields, &Field{
			Name: meta.Name,
			Ops:  ops,
			Meta: *meta,
		})
	}

	sort.Slice(it.Fields, func(i, j int) bool {
		first := it.Fields[i]
		second := it.Fields[j]

		if first.Meta.IsPK && !second.Meta.IsPK {
			return true
		}
		if !first.Meta.IsDate() && second.Meta.IsDate() {
			return true
		}

		return strings.Compare(first.Name, second.Name) > 0
	})
}

func (it *Resource) GetFields(op rune) []*Field {
	var fields []*Field
	for _, f := range it.Fields {
		if strings.ContainsRune(f.Ops, unicode.ToUpper(op)) {
			fields = append(fields, f)
		}
	}

	return fields
}
