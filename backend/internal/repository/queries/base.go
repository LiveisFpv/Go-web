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
