package queries

import (
	"backend/internal/domain"
	"context"
	"fmt"
)

func (q *Queries) GetAllSemester(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Semester, int, error) {
	getAllSemester := `SELECT 
		name_semester,
		date_start_semester::text,
		date_end_semester::text 
	FROM semester WHERE 1=1`
	countQuery := `SELECT COUNT(*) FROM semester WHERE 1=1`

	var args []interface{}
	argCount := 1

	// Handle global search if provided
	if search != "" {
		getAllSemester += fmt.Sprintf(` AND (
			name_semester ILIKE $%d OR 
			date_start_semester::text ILIKE $%d OR 
			date_end_semester::text ILIKE $%d
		)`, argCount, argCount, argCount)
		countQuery += fmt.Sprintf(` AND (
			name_semester ILIKE $%d OR 
			date_start_semester::text ILIKE $%d OR 
			date_end_semester::text ILIKE $%d
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
			getAllSemester += fmt.Sprintf(" AND %s ILIKE $%d", key, argCount)
			countQuery += fmt.Sprintf(" AND %s ILIKE $%d", key, argCount)
			args = append(args, "%"+value+"%")
			argCount++
		}
	}

	// Handle sorting
	if sortField != "" {
		// Validate sort field to prevent SQL injection
		validFields := map[string]string{
			"name_semester":       "name_semester",
			"date_start_semester": "date_start_semester",
			"date_end_semester":   "date_end_semester",
		}
		if field, ok := validFields[sortField]; ok {
			order := "ASC"
			if sortOrder == "desc" {
				order = "DESC"
			}
			getAllSemester += fmt.Sprintf(" ORDER BY %s %s", field, order)
		}
	}

	// Add pagination
	offset := (page - 1) * rowCount
	getAllSemester += fmt.Sprintf(" LIMIT %d OFFSET %d", rowCount, offset)

	rows, err := q.pool.Query(ctx, getAllSemester, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("can't query semesters: %w", err)
	}
	defer rows.Close()

	semesters := []*domain.Semester{}
	for rows.Next() {
		semester := &domain.Semester{}
		err := rows.Scan(
			&semester.Name_semester,
			&semester.Date_start_semester,
			&semester.Date_end_semester)
		if err != nil {
			return nil, 0, fmt.Errorf("can't scan semesters: %w", err)
		}
		semesters = append(semesters, semester)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	// Получаем общее количество строк без пагинации
	var count int
	err = q.pool.QueryRow(ctx, countQuery, args...).Scan(&count)
	if err != nil {
		return nil, 0, fmt.Errorf("can't count semesters: %w", err)
	}

	return semesters, count, nil
}

func (q *Queries) CreateSemester(ctx context.Context, name_semester string, date_start_semester, date_end_semester string) (*domain.Semester, error) {
	sqlStatement := `INSERT INTO semester (name_semester, date_start_semester, date_end_semester) VALUES ($1, $2, $3) RETURNING name_semester`
	err := q.pool.QueryRow(ctx, sqlStatement, name_semester, date_start_semester, date_end_semester).Scan(&name_semester)
	if err != nil {
		return nil, fmt.Errorf("can't create semester: %w", err)
	}
	semester, err := q.FindSemesterByName(ctx, name_semester)
	if err != nil {
		return nil, fmt.Errorf("can't find semester: %w", err)
	}
	return semester, err
}

func (q *Queries) FindSemesterByName(ctx context.Context, name string) (*domain.Semester, error) {
	row := q.pool.QueryRow(ctx, `SELECT 
		name_semester,
		date_start_semester::text,
		date_end_semester::text 
	FROM semester WHERE name_semester = $1`, name)

	semester := &domain.Semester{}
	err := row.Scan(&semester.Name_semester,
		&semester.Date_start_semester,
		&semester.Date_end_semester)
	if err != nil {
		return nil, fmt.Errorf("can't scan semester: %w", err)
	}

	return semester, nil
}

func (q *Queries) UpdateSemesterByName(ctx context.Context, name_semester string, date_start_semester, date_end_semester string) (*domain.Semester, error) {
	sqlStatement := `UPDATE semester SET date_start_semester=$1, date_end_semester=$2 WHERE name_semester=$3 RETURNING name_semester`
	err := q.pool.QueryRow(ctx, sqlStatement, date_start_semester, date_end_semester, name_semester).Scan(&name_semester)
	if err != nil {
		return nil, fmt.Errorf("can't update semester: %w", err)
	}
	semester, err := q.FindSemesterByName(ctx, name_semester)
	if err != nil {
		return nil, fmt.Errorf("can't find semester: %w", err)
	}
	return semester, err
}

func (q *Queries) DeleteSemesterByName(ctx context.Context, name_semester string) error {
	sqlStatement := `DELETE FROM semester WHERE name_semester=$1`
	_, err := q.pool.Exec(ctx, sqlStatement, name_semester)
	if err != nil {
		return fmt.Errorf("can't delete semester: %w", err)
	}
	return nil
}

func (q *Queries) DeleteSemesters(ctx context.Context, names []string) error {
	sqlStatement := `DELETE FROM semester WHERE name_semester = ANY($1)`
	_, err := q.pool.Exec(ctx, sqlStatement, names)
	if err != nil {
		return fmt.Errorf("can't delete semesters: %w", err)
	}
	return nil
}
