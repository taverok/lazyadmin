package db

import (
	"database/sql"
	"slices"

	"github.com/jimsmart/schema"
)

type FieldMeta struct {
	Name       string
	Type       string
	IsNullable bool
	IsPK       bool
}

func (it *FieldMeta) IsInt() bool {
	return it.Type == "INT" || it.Type == "BIGINT"
}

func (it *FieldMeta) IsDecimal() bool {
	return it.Type == "DECIMAL"
}

func (it *FieldMeta) IsNum() bool {
	return it.IsInt() || it.IsDecimal()
}

func (it *FieldMeta) IsBool() bool {
	return it.Type == "BOOL"
}

func (it *FieldMeta) IsText() bool {
	return it.Type == "VARCHAR" || it.Type == "TEXT" || it.Type == "NVARCHAR"
}

func (it *FieldMeta) IsDate() bool {
	return it.Type == "DATETIME"
}

func AnalyzeTable(db *sql.DB, table string) (map[string]*FieldMeta, error) {
	types, err := schema.ColumnTypes(db, "", table)
	if err != nil {
		return map[string]*FieldMeta{}, err
	}

	pks, err := schema.PrimaryKey(db, "", table)
	if err != nil {
		return map[string]*FieldMeta{}, err
	}

	result := map[string]*FieldMeta{}
	for _, t := range types {
		nullable, _ := t.Nullable()

		result[t.Name()] = &FieldMeta{
			Name:       t.Name(),
			Type:       t.DatabaseTypeName(),
			IsNullable: nullable,
			IsPK:       slices.Contains(pks, t.Name()),
		}
	}

	return result, nil
}
