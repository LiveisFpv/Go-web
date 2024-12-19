package repository

import (
	"backend/internal/domain"
	"backend/internal/mytype"
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
	GetCountRows(ctx context.Context, TableName string) (int, error)
	FindStudentByID(ctx context.Context, id uint64) (*domain.Student, error)
	GetAllStudent(ctx context.Context) ([]*domain.Student, error)
	CreateStudent(ctx context.Context, id_num_student uint64, name_group, email_student, second_name_student,
		first_name_student, surname_student string) (*domain.Student, error)
	UpdateStudentbyID(ctx context.Context, id_num_student uint64, name_group, email_student, second_name_student,
		first_name_student, surname_student string) (*domain.Student, error)
	DeleteStudentbyID(ctx context.Context, id_num_student uint64) error
	GetAllGroup(ctx context.Context) ([]*domain.Group, error)
	UpdateGroupbyName(ctx context.Context, group_name, Studies_direction_group,
		Studies_profile_group string, Start_date_group mytype.JsonDate, Studies_period_group uint8) (*domain.Group, error)
	CreateGroup(ctx context.Context, group_name, Studies_direction_group,
		Studies_profile_group string, Start_date_group mytype.JsonDate, Studies_period_group uint8) (*domain.Group, error)
	DeleteGroupByName(ctx context.Context, group_name string) error
	GetAllMark(ctx context.Context) ([]*domain.Mark, error)
	CreateMark(ctx context.Context, id_mark, id_num_student int64,
		lesson_name_mark, name_semester string, score_mark int8, type_mark string) (*domain.Mark, error)
	UpdateMarkByID(ctx context.Context, id_mark, id_num_student int64,
		lesson_name_mark, name_semester string, score_mark int8, type_mark string) (*domain.Mark, error)
	DeleteMarkByID(ctx context.Context, id_mark int64) error
}
