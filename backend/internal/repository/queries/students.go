package queries

import (
	"backend/internal/domain"
	"context"
	"fmt"
)

const getStudentByIDQuery = `SELECT * FROM student WHERE id_num_student = $1`

func (q *Queries) FindStudentByID(ctx context.Context, id uint64) (*domain.Student, error) {
	row := q.pool.QueryRow(ctx, getStudentByIDQuery, id)

	student := &domain.Student{}
	err := row.Scan(&student.Id_num_student,
		&student.Name_group,
		&student.Email_student,
		&student.First_name_student,
		&student.Second_name_student,
		&student.Surname_student)
	if err != nil {
		return nil, fmt.Errorf("can't scan students: %w", err)
	}

	return student, nil
}

const getAllStudent = `SELECT * FROM student`

func (q *Queries) GetAllStudent(ctx context.Context) ([]*domain.Student, error) {
	rows, err := q.pool.Query(ctx, getAllStudent)
	if err != nil {
		return nil, fmt.Errorf("can't query students: %w", err)
	}
	defer rows.Close()
	students := []*domain.Student{}
	for rows.Next() {
		student := &domain.Student{}
		err := rows.Scan(
			&student.Id_num_student,
			&student.Name_group,
			&student.Email_student,
			&student.First_name_student,
			&student.Second_name_student,
			&student.Surname_student)
		if err != nil {
			return nil, fmt.Errorf("can't scan students: %w", err)
		}
		students = append(students, student)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return students, nil
}
func (q *Queries) CreateStudent(ctx context.Context, id_num_student uint64, name_group, email_student, second_name_student,
	first_name_student, surname_student string) (*domain.Student, error) {
	sqlStatement := `INSERT INTO student (id_num_student,name_group, email_student, second_name_student, first_name_student, surname_student) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_num_student`
	err := q.pool.QueryRow(ctx, sqlStatement, id_num_student, name_group, email_student, second_name_student, first_name_student, surname_student).Scan(&id_num_student)
	if err != nil {
		return nil, fmt.Errorf("can't create student: %w", err)
	}
	student, err := q.FindStudentByID(ctx, id_num_student)
	if err != nil {
		return nil, fmt.Errorf("can't find student: %w", err)
	}
	return student, err
}
func (q *Queries) UpdateStudentbyID(ctx context.Context, id_num_student uint64, name_group, email_student, second_name_student,
	first_name_student, surname_student string) (*domain.Student, error) {
	sqlStatement := `UPDATE student SET name_group=$1, email_student=$2, second_name_student=$3, first_name_student=$4, surname_student=$5 WHERE id_num_student=$6 RETURNING id_num_student`
	err := q.pool.QueryRow(ctx, sqlStatement, name_group, email_student, second_name_student, first_name_student, surname_student, id_num_student).Scan(&id_num_student)
	if err != nil {
		return nil, fmt.Errorf("can't update student: %w", err)
	}
	student, err := q.FindStudentByID(ctx, id_num_student)
	if err != nil {
		return nil, fmt.Errorf("can't find student: %w", err)
	}
	return student, err
}
func (q *Queries) DeleteStudentbyID(ctx context.Context, id_num_student uint64) error {
	sqlStatement := `DELETE FROM student WHERE id_num_student=$1`
	_, err := q.pool.Exec(ctx, sqlStatement, id_num_student)
	if err != nil {
		return fmt.Errorf("can't delete student: %w", err)
	}
	return nil
}