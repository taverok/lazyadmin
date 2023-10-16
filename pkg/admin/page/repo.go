package page

import (
	"database/sql"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/taverok/lazyadmin/pkg/admin/resource"
)

type MysqlRepo struct {
	DB *sqlx.DB
}

func (it *MysqlRepo) GetAll(fields []*resource.Field, table string) ([][]any, error) {
	var result [][]any
	fieldNames := resource.FieldNames(fields)

	query, _, err := sq.Select(fieldNames...).From(table).ToSql()
	if err != nil {
		return result, err
	}

	rows, err := it.DB.Queryx(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		cols, err := rows.SliceScan()
		if err != nil {
			slog.Error(err.Error())
			continue
		}
		result = append(result, cols)
	}

	return result, nil
}

func (it *MysqlRepo) GetById(table string, fields []*resource.Field, where map[string]string) ([]any, error) {
	var result []any
	fieldNames := resource.FieldNames(fields)

	criteria := sq.Select(fieldNames...).From(table)
	for k, v := range where {
		criteria.Where(sq.Eq{k: v})
	}
	query, _, err := criteria.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := it.DB.Queryx(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		return rows.SliceScan()
	}

	return nil, sql.ErrNoRows
}
