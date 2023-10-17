package page

import "github.com/taverok/lazyadmin/pkg/admin/resource"

type Page[T any] struct {
	Name    string
	Content T
	Menu    []string
	Error   string
}

func (it Page[T]) HasError() bool {
	return it.Error != ""
}

type Empty struct {
}

type Table struct {
	Fields []*resource.Field
	Data   [][]any //TODO swap for FieldValue
}

type Form struct {
	Fields []FieldValue
}

func NewForm(fields []*resource.Field, data []any) Form {
	var values []FieldValue
	for i, f := range fields {
		value := data[i]
		values = append(values, NewFieldValue(f, value))
	}

	return Form{
		Fields: values,
	}
}

type FieldValue struct {
	*resource.Field
	Value any
}

func NewFieldValue(f *resource.Field, v any) FieldValue {
	return FieldValue{
		Field: f,
		Value: v,
	}
}

type Login struct {
	User string
	Pass string
}
