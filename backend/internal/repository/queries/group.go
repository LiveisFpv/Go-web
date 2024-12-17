package queries

import (
	"backend/internal/domain"
	"context"
	"fmt"
)

func (q *Queries) GetAllGroup(ctx context.Context) ([]*domain.Group, error) {
	sqlStatement := `SELECT * FROM "group"`
	rows, err := q.pool.Query(ctx, sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("can't query groups: %w", err)
	}
	defer rows.Close()

	var groups []*domain.Group

	for rows.Next() {
		group := &domain.Group{}
		err := rows.Scan(
			&group.Name_group,
			&group.Studies_direction_group,
			&group.Studies_profile_group,
			&group.Start_date_group,
			&group.Studies_period_group,
			&group.Id_group,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan groups: %w", err)
		}
		groups = append(groups, group)
	}
	return groups, nil
}
