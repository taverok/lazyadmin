package resource

import "github.com/taverok/lazyadmin/pkg/db"

type Field struct {
	Name string `yaml:"name"`
	Px   int    `yaml:"px"`
	Ops  string `yaml:"ops"`
	Meta db.FieldMeta
}

func (it *Field) SetDefaults(meta *db.FieldMeta) {
	it.Meta = *meta

	if it.Ops == "" {
		it.Ops = "CRU"
	}
}

func (it Field) GetPxWidth() int {
	if it.Px > 0 {
		return it.Px
	}

	if it.Meta.IsPK && it.Meta.IsNum() {
		return 20
	}

	return 100
}
