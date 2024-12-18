package queries

import (
	"backend/internal/domain"
	"context"
	"fmt"
	"time"
)

func (q *Queries) FindGroupByName(ctx context.Context, name string) (*domain.Group, error) {
	sqlStatement := `SELECT * FROM "group" WHERE name_group = $1`
	group := &domain.Group{}
	err := q.pool.QueryRow(ctx, sqlStatement, name).Scan(
		&group.Name_group,
		&group.Studies_direction_group,
		&group.Studies_profile_group,
		&group.Start_date_group,
		&group.Studies_period_group)
	if err != nil {
		return nil, fmt.Errorf("can't find group: %w", err)
	}
	return group, nil
}

func (q *Queries) UpdateGroupbyName(ctx context.Context, group_name, Studies_direction_group,
	Studies_profile_group string, Start_date_group time.Time, Studies_period_group uint8) (*domain.Group, error) {
	sqlStatement := `UPDATE "group" SET studies_direction_group=$1, studies_profile_group=$2, start_date_group=$3, studies_period_group=$4 WHERE name_group=$5 RETURNING name_group`
	err := q.pool.QueryRow(ctx, sqlStatement, Studies_direction_group, Studies_profile_group, Start_date_group, Studies_period_group, group_name).Scan(&group_name)
	if err != nil {
		return nil, fmt.Errorf("can't update group: %w", err)
	}
	group, err := q.FindGroupByName(ctx, group_name)
	if err != nil {
		return nil, fmt.Errorf("can't find group: %w", err)
	}
	return group, nil
}

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
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan groups: %w", err)
		}
		groups = append(groups, group)
	}
	return groups, nil
}
