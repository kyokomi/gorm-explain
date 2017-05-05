package gorm_explain

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/bndr/gotabulate"
	"github.com/jinzhu/gorm"
)

// Callback output explain callback
func Callback(scope *gorm.Scope) {
	if !strings.HasPrefix(strings.ToUpper(scope.SQL), "SELECT") {
		return
	}

	rows, err := scope.SQLDB().Query("EXPLAIN "+scope.SQL, scope.SQLVars...)
	if scope.Err(err) != nil {
		return
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	results, _ := convertToResult(rows)

	// Create an object from 2D interface array
	table := gotabulate.Create(results)
	table.SetHeaders(columns)
	table.SetEmptyString("None")
	table.SetAlign("right")
	fmt.Println(table.Render("grid"))
}

func convertToResult(rows *sql.Rows) (res [][]interface{}, err error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	args := make([]interface{}, len(values))

	for i := range values {
		args[i] = &values[i]
	}

	for rows.Next() {
		if err := rows.Scan(args...); err != nil {
			return nil, err
		}
		row := []interface{}{}
		for _, col := range values {
			row = append(row, string(col))
		}
		res = append(res, row)
	}

	return res, nil
}
