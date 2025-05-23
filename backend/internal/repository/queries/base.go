package queries

import (
	"context"
	"fmt"
)

func (q *Queries) GetCountRows(ctx context.Context, TableName string) (int, error) {
	sqlStatements := `SELECT count(*) FROM "` + TableName + `"`
	var num int
	err := q.pool.QueryRow(ctx, sqlStatements).Scan(&num)
	if err != nil {
		return 0, fmt.Errorf("Cannot get count from table %s: %w", TableName, err)
	}
	return num, nil
}
func UnpackFilter(ctx context.Context, getAll, countQuery string, filters map[string]string, rowCount, page int) (string, string, []any) {
	var args []interface{}
	for key, value := range filters {
		getAll += fmt.Sprintf(" AND %s ILIKE $%d", key, len(args)+1)
		countQuery += fmt.Sprintf(" AND %s ILIKE $%d", key, len(args)+1)
		args = append(args, "%"+value+"%")
	}
	offset := (page - 1) * rowCount
	getAll += fmt.Sprintf(" LIMIT %d OFFSET %d", rowCount, offset)
	return getAll, countQuery, args
}
