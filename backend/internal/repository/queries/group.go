package queries

import (
	"backend/internal/domain"
	"backend/internal/mytype"
	"context"
	"fmt"
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

func (q *Queries) CreateGroup(ctx context.Context, group_name, Studies_direction_group,
	Studies_profile_group string, Start_date_group mytype.JsonDate, Studies_period_group uint8) (*domain.Group, error) {
	sqlStatement := `INSERT INTO "group" (name_group, studies_direction_group, studies_profile_group, start_date_group, studies_period_group) VALUES ($1, $2, $3, $4, $5) RETURNING name_group`
	err := q.pool.QueryRow(ctx, sqlStatement, group_name, Studies_direction_group, Studies_profile_group, Start_date_group, Studies_period_group).Scan(&group_name)
	if err != nil {
		return nil, fmt.Errorf("can't create group: %w", err)
	}
	group, err := q.FindGroupByName(ctx, group_name)
	if err != nil {
		return nil, fmt.Errorf("can't find group: %w", err)
	}
	return group, nil
}

func (q *Queries) UpdateGroupbyName(ctx context.Context, group_name, Studies_direction_group,
	Studies_profile_group string, Start_date_group mytype.JsonDate, Studies_period_group uint8) (*domain.Group, error) {
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

func (q *Queries) DeleteGroupByName(ctx context.Context, name string) error {
	sqlStatement := `DELETE FROM "group" WHERE name_group=$1`
	_, err := q.pool.Exec(ctx, sqlStatement, name)
	if err != nil {
		return fmt.Errorf("can't delete group: %w", err)
	}
	return nil
}

func (q *Queries) GetAllGroup(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Group, int, error) {
	getAll := `SELECT * FROM "group" WHERE 1=1`
	countQuery := `SELECT COUNT(*) FROM "group" WHERE 1=1`

	var args []interface{}
	argCount := 1

	// Handle global search if provided
	if search != "" {
		getAll += fmt.Sprintf(` AND (
			name_group ILIKE $%d OR 
			studies_direction_group ILIKE $%d OR 
			studies_profile_group ILIKE $%d
		)`, argCount, argCount, argCount)
		countQuery += fmt.Sprintf(` AND (
			name_group ILIKE $%d OR 
			studies_direction_group ILIKE $%d OR 
			studies_profile_group ILIKE $%d
		)`, argCount, argCount, argCount)
		args = append(args, "%"+search+"%")
		argCount++
	}

	// Extract sorting parameters
	sortField := filters["sort"]
	sortOrder := filters["order"]
	delete(filters, "sort")
	delete(filters, "order")

	// Handle specific filters
	for key, value := range filters {
		if value != "" {
			getAll += fmt.Sprintf(" AND %s ILIKE $%d", key, argCount)
			countQuery += fmt.Sprintf(" AND %s ILIKE $%d", key, argCount)
			args = append(args, "%"+value+"%")
			argCount++
		}
	}

	// Handle sorting
	if sortField != "" {
		// Validate sort field to prevent SQL injection
		validFields := map[string]string{
			"name_group":              "name_group",
			"studies_direction_group": "studies_direction_group",
			"studies_profile_group":   "studies_profile_group",
			"start_date_group":        "start_date_group",
			"studies_period_group":    "studies_period_group",
		}
		if field, ok := validFields[sortField]; ok {
			order := "ASC"
			if sortOrder == "desc" {
				order = "DESC"
			}
			getAll += fmt.Sprintf(" ORDER BY %s %s", field, order)
		}
	}

	// Add pagination
	offset := (page - 1) * rowCount
	getAll += fmt.Sprintf(" LIMIT %d OFFSET %d", rowCount, offset)

	rows, err := q.pool.Query(ctx, getAll, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("can't query groups: %w", err)
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
			return nil, 0, fmt.Errorf("can't scan groups: %w", err)
		}
		groups = append(groups, group)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	// Get total count without pagination
	var count int
	err = q.pool.QueryRow(ctx, countQuery, args...).Scan(&count)
	if err != nil {
		return nil, 0, fmt.Errorf("can't count groups: %w", err)
	}

	return groups, count, nil
}
