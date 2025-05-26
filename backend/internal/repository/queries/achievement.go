package queries

import (
	"backend/internal/domain"
	"context"
	"fmt"
)

func (q *Queries) GetAchievementByID(ctx context.Context, id int64) (*domain.Achievement, error) {
	sqlStatement := `SELECT id_achivment, id_num_student, id_category, name_achivement, 
		TO_CHAR(date_achivment, 'YYYY-MM-DD') as date_achivment 
		FROM "achievement" WHERE id_achivment = $1`
	achievement := &domain.Achievement{}
	err := q.pool.QueryRow(ctx, sqlStatement, id).Scan(
		&achievement.Id_achivment,
		&achievement.Id_num_student,
		&achievement.Id_category,
		&achievement.Name_achivement,
		&achievement.Date_achivment,
	)
	if err != nil {
		return nil, fmt.Errorf("can't query achievement: %w", err)
	}
	return achievement, nil
}

func (q *Queries) GetAllAchievement(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Achievement, int, error) {
	getAll := `SELECT a.id_achivment, a.id_num_student, a.id_category, a.name_achivement, 
		TO_CHAR(a.date_achivment, 'YYYY-MM-DD') as date_achivment,
		st.surname_student, st.first_name_student, st.second_name_student, 
		c.achivments_type_category, st.name_group 
		FROM achievement a 
		LEFT JOIN student st ON a.id_num_student = st.id_num_student 
		LEFT JOIN category c ON a.id_category = c.id_category
		WHERE 1=1`
	countQuery := `SELECT COUNT(*) 
		FROM achievement a 
		LEFT JOIN student st ON a.id_num_student = st.id_num_student 
		LEFT JOIN category c ON a.id_category = c.id_category
		WHERE 1=1`

	var args []interface{}
	argCount := 1

	// Handle global search if provided
	if search != "" {
		getAll += fmt.Sprintf(` AND (
			CAST(a.id_achivment AS TEXT) ILIKE $%d OR 
			CAST(a.id_num_student AS TEXT) ILIKE $%d OR 
			CAST(a.id_category AS TEXT) ILIKE $%d OR 
			a.name_achivement ILIKE $%d OR 
			CAST(a.date_achivment AS TEXT) ILIKE $%d OR
			st.surname_student ILIKE $%d OR
			st.first_name_student ILIKE $%d OR
			st.second_name_student ILIKE $%d OR
			c.achivments_type_category ILIKE $%d OR
			st.name_group ILIKE $%d
		)`, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount)
		countQuery += fmt.Sprintf(` AND (
			CAST(a.id_achivment AS TEXT) ILIKE $%d OR 
			CAST(a.id_num_student AS TEXT) ILIKE $%d OR 
			CAST(a.id_category AS TEXT) ILIKE $%d OR 
			a.name_achivement ILIKE $%d OR 
			CAST(a.date_achivment AS TEXT) ILIKE $%d OR
			st.surname_student ILIKE $%d OR
			st.first_name_student ILIKE $%d OR
			st.second_name_student ILIKE $%d OR
			c.achivments_type_category ILIKE $%d OR
			st.name_group ILIKE $%d
		)`, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount)
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
			// Map filter keys to their correct table references
			fieldMap := map[string]string{
				"id_achivment":             "a.id_achivment",
				"id_num_student":           "a.id_num_student",
				"id_category":              "a.id_category",
				"name_achivement":          "a.name_achivement",
				"date_achivment":           "a.date_achivment",
				"surname_student":          "st.surname_student",
				"first_name_student":       "st.first_name_student",
				"second_name_student":      "st.second_name_student",
				"achivments_type_category": "c.achivments_type_category",
				"name_group":               "st.name_group",
			}

			if field, ok := fieldMap[key]; ok {
				// Special handling for date filtering
				if key == "date_achivment" {
					getAll += fmt.Sprintf(" AND %s >= $%d", field, argCount)
					countQuery += fmt.Sprintf(" AND %s >= $%d", field, argCount)
					args = append(args, value)
					argCount++
				} else if key == "id_achivment" || key == "id_num_student" || key == "id_category" {
					getAll += fmt.Sprintf(" AND CAST(%s AS TEXT) ILIKE $%d", field, argCount)
					countQuery += fmt.Sprintf(" AND CAST(%s AS TEXT) ILIKE $%d", field, argCount)
					args = append(args, "%"+value+"%")
					argCount++
				} else {
					getAll += fmt.Sprintf(" AND %s ILIKE $%d", field, argCount)
					countQuery += fmt.Sprintf(" AND %s ILIKE $%d", field, argCount)
					args = append(args, "%"+value+"%")
					argCount++
				}
			}
		}
	}

	// Handle sorting
	if sortField != "" {
		// Validate sort field to prevent SQL injection
		validFields := map[string]string{
			"id_achivment":             "a.id_achivment",
			"id_num_student":           "a.id_num_student",
			"id_category":              "a.id_category",
			"name_achivement":          "a.name_achivement",
			"date_achivment":           "a.date_achivment",
			"surname_student":          "st.surname_student",
			"first_name_student":       "st.first_name_student",
			"second_name_student":      "st.second_name_student",
			"achivments_type_category": "c.achivments_type_category",
			"name_group":               "st.name_group",
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
		return nil, 0, fmt.Errorf("can't query achievements: %w", err)
	}
	defer rows.Close()

	var achievements []*domain.Achievement
	for rows.Next() {
		achievement := &domain.Achievement{}
		err := rows.Scan(
			&achievement.Id_achivment,
			&achievement.Id_num_student,
			&achievement.Id_category,
			&achievement.Name_achivement,
			&achievement.Date_achivment,
			&achievement.Student_surname,
			&achievement.Student_name,
			&achievement.Student_second_name,
			&achievement.Category_type,
			&achievement.Name_group,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("can't scan achievements: %w", err)
		}
		achievements = append(achievements, achievement)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	var count int
	err = q.pool.QueryRow(ctx, countQuery, args...).Scan(&count)
	if err != nil {
		return nil, 0, fmt.Errorf("can't count achievements: %w", err)
	}

	return achievements, count, nil
}

func (q *Queries) CreateAchievement(ctx context.Context, id_num_student int64,
	id_category int64, name_achivement string, date_achivment string) (*domain.Achievement, error) {
	sqlStatement := `INSERT INTO "achievement" (id_num_student, id_category, name_achivement, date_achivment) 
		VALUES ($1, $2, $3, $4) RETURNING id_achivment`
	row := q.pool.QueryRow(ctx, sqlStatement, id_num_student, id_category, name_achivement, date_achivment)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("can't create achievement: %w", err)
	}
	achievement, err := q.GetAchievementByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't find achievement: %w", err)
	}
	return achievement, nil
}

func (q *Queries) UpdateAchievementByID(ctx context.Context, id_achivment, id_num_student int64,
	id_category int64, name_achivement string, date_achivment string) (*domain.Achievement, error) {
	sqlStatement := `UPDATE "achievement" 
		SET id_num_student=$2, id_category=$3, name_achivement=$4, date_achivment=$5 
		WHERE id_achivment=$1 RETURNING id_achivment`
	row := q.pool.QueryRow(ctx, sqlStatement, id_achivment, id_num_student, id_category, name_achivement, date_achivment)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("can't update achievement: %w", err)
	}
	achievement, err := q.GetAchievementByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't find achievement: %w", err)
	}
	return achievement, nil
}

func (q *Queries) DeleteAchievementByID(ctx context.Context, id_achivment int64) error {
	sqlStatement := `DELETE FROM "achievement" WHERE id_achivment=$1`
	_, err := q.pool.Exec(ctx, sqlStatement, id_achivment)
	if err != nil {
		return fmt.Errorf("can't delete achievement: %w", err)
	}
	return nil
}

func (q *Queries) DeleteAchievements(ctx context.Context, ids []int64) error {
	sqlStatement := `DELETE FROM "achievement" WHERE id_achivment = ANY($1)`
	_, err := q.pool.Exec(ctx, sqlStatement, ids)
	if err != nil {
		return fmt.Errorf("can't delete achievements: %w", err)
	}
	return nil
}
