package queries

import (
	"backend/internal/domain"
	"context"
	"fmt"
)

const getStudentByIDQuery = `SELECT id, name, birthday FROM users WHERE id = $1`

func (q *Queries) FindStudentByID(ctx context.Context, id int) (*domain.Student, error) {
	row := q.pool.QueryRow(ctx, getStudentByIDQuery, id)

	student := &domain.Student{}
	if err := row.Scan(&student.ID, &student.Name, &student.Birthday); err != nil {
		return nil, fmt.Errorf("can't scan students: %w", err)
	}

	return student, nil
}
