package queries

import (
	"backend/internal/domain"
	"context"
	"fmt"
)

func (q *Queries) GetMarkByID(ctx context.Context, id int64) (*domain.Mark, error) {
	sqlStatement := `SELECT * FROM "mark" WHERE id_mark = $1`
	mark := &domain.Mark{}
	err := q.pool.QueryRow(ctx, sqlStatement, id).Scan(
		&mark.Id_mark,
		&mark.Id_num_student,
		&mark.Lesson_name_mark,
		&mark.Name_semester,
		&mark.Score_mark,
		&mark.Type_mark,
		&mark.Type_exam,
	)
	if err != nil {
		return nil, fmt.Errorf("can't query mark: %w", err)
	}
	return mark, nil
}
func (q *Queries) GetAllMark(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Mark, int, error) {
	getAll := `SELECT m.*, s.surname_student, s.first_name_student, s.second_name_student, s.name_group 
		FROM mark m 
		LEFT JOIN student s ON m.id_num_student = s.id_num_student 
		WHERE 1=1`
	countQuery := `SELECT COUNT(*) 
		FROM mark m 
		LEFT JOIN student s ON m.id_num_student = s.id_num_student 
		WHERE 1=1`

	var args []interface{}
	argCount := 1

	// Handle global search if provided
	if search != "" {
		getAll += fmt.Sprintf(` AND (
			CAST(m.id_mark AS TEXT) ILIKE $%d OR 
			CAST(m.id_num_student AS TEXT) ILIKE $%d OR 
			m.lesson_name_mark ILIKE $%d OR 
			m.name_semester ILIKE $%d OR 
			CAST(m.score_mark AS TEXT) ILIKE $%d OR 
			m.type_mark ILIKE $%d OR
			m.type_exam ILIKE $%d OR
			s.surname_student ILIKE $%d OR
			s.first_name_student ILIKE $%d OR
			s.second_name_student ILIKE $%d OR
			s.name_group ILIKE $%d
		)`, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount)
		countQuery += fmt.Sprintf(` AND (
			CAST(m.id_mark AS TEXT) ILIKE $%d OR 
			CAST(m.id_num_student AS TEXT) ILIKE $%d OR 
			m.lesson_name_mark ILIKE $%d OR 
			m.name_semester ILIKE $%d OR 
			CAST(m.score_mark AS TEXT) ILIKE $%d OR 
			m.type_mark ILIKE $%d OR
			m.type_exam ILIKE $%d OR
			s.surname_student ILIKE $%d OR
			s.first_name_student ILIKE $%d OR
			s.second_name_student ILIKE $%d OR
			s.name_group ILIKE $%d
		)`, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount, argCount)
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
				"id_mark":             "m.id_mark",
				"id_num_student":      "m.id_num_student",
				"lesson_name_mark":    "m.lesson_name_mark",
				"name_semester":       "m.name_semester",
				"score_mark":          "m.score_mark",
				"type_mark":           "m.type_mark",
				"type_exam":           "m.type_exam",
				"surname_student":     "s.surname_student",
				"first_name_student":  "s.first_name_student",
				"second_name_student": "s.second_name_student",
				"name_group":          "s.name_group",
			}

			if field, ok := fieldMap[key]; ok {
				// Handle numeric fields differently
				if key == "id_mark" || key == "id_num_student" || key == "score_mark" {
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
			"id_mark":             "m.id_mark",
			"id_num_student":      "m.id_num_student",
			"lesson_name_mark":    "m.lesson_name_mark",
			"name_semester":       "m.name_semester",
			"score_mark":          "m.score_mark",
			"type_mark":           "m.type_mark",
			"type_exam":           "m.type_exam",
			"surname_student":     "s.surname_student",
			"first_name_student":  "s.first_name_student",
			"second_name_student": "s.second_name_student",
			"name_group":          "s.name_group",
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
		return nil, 0, fmt.Errorf("can't query marks: %w", err)
	}
	defer rows.Close()

	var marks []*domain.Mark
	for rows.Next() {
		mark := &domain.Mark{}
		err := rows.Scan(
			&mark.Id_mark,
			&mark.Id_num_student,
			&mark.Name_semester,
			&mark.Lesson_name_mark,
			&mark.Score_mark,
			&mark.Type_mark,
			&mark.Type_exam,
			&mark.Student_surname,
			&mark.Student_name,
			&mark.Student_second_name,
			&mark.Student_group,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("can't scan marks: %w", err)
		}
		marks = append(marks, mark)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	var count int
	err = q.pool.QueryRow(ctx, countQuery, args...).Scan(&count)
	if err != nil {
		return nil, 0, fmt.Errorf("can't count marks: %w", err)
	}
	fmt.Println(getAll)
	return marks, count, nil
}
func (q *Queries) CreateMark(ctx context.Context, id_mark, id_num_student int64,
	name_semester, lesson_name_mark string, score_mark int8, type_mark, type_exam string) (*domain.Mark, error) {
	sqlStatement := `INSERT INTO "mark" (id_num_student, name_semester, lesson_name_mark, score_mark, type_mark, type_exam) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_mark`
	row := q.pool.QueryRow(ctx, sqlStatement, id_num_student, name_semester, lesson_name_mark, score_mark, type_mark, type_exam)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("can't create mark: %w", err)
	}
	mark, err := q.GetMarkByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't find mark: %w", err)
	}
	return mark, nil
}
func (q *Queries) UpdateMarkByID(ctx context.Context, id_mark, id_num_student int64,
	name_semester, lesson_name_mark string, score_mark int8, type_mark, type_exam string) (*domain.Mark, error) {
	sqlStatement := `UPDATE "mark" SET id_num_student=$2, name_semester=$3, lesson_name_mark=$4, score_mark=$5, type_mark=$6, type_exam=$7 
		WHERE id_mark=$1 RETURNING id_mark`
	row := q.pool.QueryRow(ctx, sqlStatement, id_mark, id_num_student, name_semester, lesson_name_mark, score_mark, type_mark, type_exam)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("can't update mark: %w", err)
	}
	mark, err := q.GetMarkByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't find mark: %w", err)
	}
	return mark, nil
}
func (q *Queries) DeleteMarkByID(ctx context.Context, id_mark int64) error {
	sqlStatement := `DELETE FROM "mark" WHERE id_mark=$1`
	_, err := q.pool.Exec(ctx, sqlStatement, id_mark)
	if err != nil {
		return fmt.Errorf("can't delete mark: %w", err)
	}
	return nil
}
