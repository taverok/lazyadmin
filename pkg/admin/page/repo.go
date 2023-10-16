package page

import (
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
	var fieldNames []string

	for _, f := range fields {
		fieldNames = append(fieldNames, f.Name)
	}

	sql, _, err := sq.Select(fieldNames...).From(table).ToSql()
	if err != nil {
		return result, err
	}

	rows, err := it.DB.Queryx(sql)
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
