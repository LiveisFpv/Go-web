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
	)
	if err != nil {
		return nil, fmt.Errorf("can't query mark: %w", err)
	}
	return mark, nil
}
func (q *Queries) GetAllMark(ctx context.Context, filters map[string]string, rowCount, page int) ([]*domain.Mark, int, error) {
	getAll := `SELECT * FROM mark WHERE 1=1`
	countQuery := `SELECT COUNT(*) FROM mark WHERE 1=1`
	var args []interface{}
	getAll, countQuery, args = UnpackFilter(ctx, getAll, countQuery, filters, rowCount, page)
	rows, err := q.pool.Query(ctx, getAll, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("can't query mark: %w", err)
	}
	defer rows.Close()

	var marks []*domain.Mark

	for rows.Next() {
		mark := &domain.Mark{}
		err := rows.Scan(
			&mark.Id_mark,
			&mark.Id_num_student,
			&mark.Lesson_name_mark,
			&mark.Name_semester,
			&mark.Score_mark,
			&mark.Type_mark,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("can't scan marks: %w", err)
		}
		marks = append(marks, mark)
	}
	var count int
	err = q.pool.QueryRow(ctx, countQuery, args...).Scan(&count)
	if err != nil {
		return nil, 0, fmt.Errorf("can't count marks: %w", err)
	}
	return marks, count, nil
}
func (q *Queries) CreateMark(ctx context.Context, id_mark, id_num_student int64,
	lesson_name_mark, name_semester string, score_mark int8, type_mark string) (*domain.Mark, error) {
	sqlStatement := `INSERT INTO "mark" (id_mark, id_num_student, lesson_name_mark, name_semester, score_mark, type_mark) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_mark`
	row := q.pool.QueryRow(ctx, sqlStatement, id_mark, id_num_student, lesson_name_mark, name_semester, score_mark, type_mark)
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
	lesson_name_mark, name_semester string, score_mark int8, type_mark string) (*domain.Mark, error) {
	sqlStatement := `UPDATE "mark" SET id_num_student=$2, lesson_name_mark=$3, name_semester=$4, score_mark=$5, type_mark=$6 WHERE id_mark=$1 RETURNING id_mark`
	row := q.pool.QueryRow(ctx, sqlStatement, id_mark, id_num_student, lesson_name_mark, name_semester, score_mark, type_mark)
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
