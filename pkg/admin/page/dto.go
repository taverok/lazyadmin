package page

import "github.com/taverok/lazyadmin/pkg/admin/resource"

type Page[T any] struct {
	Name    string
	Content T
	Menu    []string
}

type Table struct {
	Fields []*resource.Field
	Data   [][]any
}
