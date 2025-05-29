package queries

import (
	"backend/internal/domain"
	"context"
	"fmt"
)

func (q *Queries) GetScholarshipByID(ctx context.Context, id int64) (*domain.Scholarship, error) {
	sqlStatement := `SELECT * FROM "scholarship" WHERE id_scholarship = $1`
	scholarship := &domain.Scholarship{}
	err := q.pool.QueryRow(ctx, sqlStatement, id).Scan(
		&scholarship.Id_scholarship,
		&scholarship.Id_num_student,
		&scholarship.Name_semester,
		&scholarship.Size_scholarshp,
		&scholarship.Id_budget,
	)
	if err != nil {
		return nil, fmt.Errorf("can't query scholarship: %w", err)
	}
	return scholarship, nil
}

func (q *Queries) GetAllScholarship(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Scholarship, int, error) {
	getAll := `SELECT s.*, st.surname_student, st.first_name_student, st.second_name_student, st.name_group, b.type_scholarship_budget 
		FROM scholarship s 
		LEFT JOIN student st ON s.id_num_student = st.id_num_student 
		LEFT JOIN budget b ON s.id_budget = b.id_budget
		WHERE 1=1`
	countQuery := `SELECT COUNT(*) 
		FROM scholarship s 
		LEFT JOIN student st ON s.id_num_student = st.id_num_student 
		LEFT JOIN budget b ON s.id_budget = b.id_budget
		WHERE 1=1`

	var args []interface{}
	argCount := 1

	// Handle global search if provided
	if search != "" {
		getAll += fmt.Sprintf(` AND (
			CAST(s.id_scholarship AS TEXT) ILIKE $%d OR 
			CAST(s.id_num_student AS TEXT) ILIKE $%d OR 
			s.name_semester ILIKE $%d OR 
			CAST(s.size_scholarshp AS TEXT) ILIKE $%d OR 
			CAST(s.id_budget AS TEXT) ILIKE $%d OR
			st.surname_student ILIKE $%d OR
			st.first_name_student ILIKE $%d OR
			st.second_name_student ILIKE $%d OR
			st.name_group ILIKE $%d OR
			b.type_scholarship_budget ILIKE $%d
		)`, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount)
		countQuery += fmt.Sprintf(` AND (
			CAST(s.id_scholarship AS TEXT) ILIKE $%d OR 
			CAST(s.id_num_student AS TEXT) ILIKE $%d OR 
			s.name_semester ILIKE $%d OR 
			CAST(s.size_scholarshp AS TEXT) ILIKE $%d OR 
			CAST(s.id_budget AS TEXT) ILIKE $%d OR
			st.surname_student ILIKE $%d OR
			st.first_name_student ILIKE $%d OR
			st.second_name_student ILIKE $%d OR
			st.name_group ILIKE $%d OR
			b.type_scholarship_budget ILIKE $%d
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
				"id_scholarship":          "s.id_scholarship",
				"id_num_student":          "s.id_num_student",
				"name_semester":           "s.name_semester",
				"size_scholarshp":         "s.size_scholarshp",
				"id_budget":               "s.id_budget",
				"surname_student":         "st.surname_student",
				"first_name_student":      "st.first_name_student",
				"second_name_student":     "st.second_name_student",
				"name_group":              "st.name_group",
				"type_scholarship_budget": "b.type_scholarship_budget",
			}

			if field, ok := fieldMap[key]; ok {
				// Handle numeric fields differently
				if key == "id_scholarship" || key == "id_num_student" || key == "id_budget" || key == "size_scholarshp" {
					getAll += fmt.Sprintf(" AND CAST(%s AS TEXT) ILIKE $%d", field, argCount)
					countQuery += fmt.Sprintf(" AND CAST(%s AS TEXT) ILIKE $%d", field, argCount)
				} else {
					getAll += fmt.Sprintf(" AND %s ILIKE $%d", field, argCount)
					countQuery += fmt.Sprintf(" AND %s ILIKE $%d", field, argCount)
				}
				args = append(args, "%"+value+"%")
				argCount++
			}
		}
	}

	// Handle sorting
	if sortField != "" {
		// Validate sort field to prevent SQL injection
		validFields := map[string]string{
			"id_scholarship":          "s.id_scholarship",
			"id_num_student":          "s.id_num_student",
			"name_semester":           "s.name_semester",
			"size_scholarshp":         "s.size_scholarshp",
			"id_budget":               "s.id_budget",
			"surname_student":         "st.surname_student",
			"first_name_student":      "st.first_name_student",
			"second_name_student":     "st.second_name_student",
			"name_group":              "st.name_group",
			"type_scholarship_budget": "b.type_scholarship_budget",
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
		return nil, 0, fmt.Errorf("can't query scholarships: %w", err)
	}
	defer rows.Close()

	var scholarships []*domain.Scholarship
	for rows.Next() {
		scholarship := &domain.Scholarship{}
		err := rows.Scan(
			&scholarship.Id_scholarship,
			&scholarship.Id_num_student,
			&scholarship.Name_semester,
			&scholarship.Size_scholarshp,
			&scholarship.Id_budget,
			&scholarship.Student_surname,
			&scholarship.Student_name,
			&scholarship.Student_second_name,
			&scholarship.Student_group,
			&scholarship.Type_scholarship_budget,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("can't scan scholarships: %w", err)
		}
		scholarships = append(scholarships, scholarship)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	var count int
	err = q.pool.QueryRow(ctx, countQuery, args...).Scan(&count)
	if err != nil {
		return nil, 0, fmt.Errorf("can't count scholarships: %w", err)
	}

	return scholarships, count, nil
}

func (q *Queries) CreateScholarship(ctx context.Context, id_num_student int64,
	name_semester string, size_scholarshp float64, id_budget int64) (*domain.Scholarship, error) {
	sqlStatement := `INSERT INTO "scholarship" (id_num_student, name_semester, size_scholarshp, id_budget) 
		VALUES ($1, $2, $3, $4) RETURNING id_scholarship`
	row := q.pool.QueryRow(ctx, sqlStatement, id_num_student, name_semester, size_scholarshp, id_budget)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("can't create scholarship: %w", err)
	}
	scholarship, err := q.GetScholarshipByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't find scholarship: %w", err)
	}
	return scholarship, nil
}

func (q *Queries) UpdateScholarshipByID(ctx context.Context, id_scholarship, id_num_student int64,
	name_semester string, size_scholarshp float64, id_budget int64) (*domain.Scholarship, error) {
	sqlStatement := `UPDATE "scholarship" 
		SET id_num_student=$2, name_semester=$3, size_scholarshp=$4, id_budget=$5 
		WHERE id_scholarship=$1 RETURNING id_scholarship`
	row := q.pool.QueryRow(ctx, sqlStatement, id_scholarship, id_num_student, name_semester, size_scholarshp, id_budget)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("can't update scholarship: %w", err)
	}
	scholarship, err := q.GetScholarshipByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't find scholarship: %w", err)
	}
	return scholarship, nil
}

func (q *Queries) DeleteScholarshipByID(ctx context.Context, id_scholarship int64) error {
	sqlStatement := `DELETE FROM "scholarship" WHERE id_scholarship=$1`
	_, err := q.pool.Exec(ctx, sqlStatement, id_scholarship)
	if err != nil {
		return fmt.Errorf("can't delete scholarship: %w", err)
	}
	return nil
}

func (q *Queries) AssignScholarships(ctx context.Context, current_semester, budget_type string) error {
	sqlStatement := `CALL public.assign_scholarships($1, $2)`
	_, err := q.pool.Exec(ctx, sqlStatement, current_semester, budget_type)
	if err != nil {
		return fmt.Errorf("can't assign scholarships: %w", err)
	}
	return nil
}
