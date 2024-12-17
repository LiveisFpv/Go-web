package repository

import (
	"backend/internal/domain"
	"backend/internal/repository/queries"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type repo struct {
	*queries.Queries
	pool   *pgxpool.Pool
	logger logrus.FieldLogger
}

func NewRepository(pgxPool *pgxpool.Pool, logger logrus.FieldLogger) Repository {
	return &repo{
		Queries: queries.New(pgxPool),
		pool:    pgxPool,
		logger:  logger,
	}
}

type Repository interface {
	FindStudentByID(ctx context.Context, id uint64) (*domain.Student, error)
	GetAllStudent(ctx context.Context) ([]*domain.Student, error)
	CreateStudent(ctx context.Context, id_num_student uint64, name_group, email_student, second_name_student,
		first_name_student, surname_student string) (*domain.Student, error)
	UpdateStudentbyID(ctx context.Context, id_num_student uint64, name_group, email_student, second_name_student,
		first_name_student, surname_student string) (*domain.Student, error)
	DeleteStudentbyID(ctx context.Context, id_num_student uint64) error
}
