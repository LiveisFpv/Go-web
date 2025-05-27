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
	Login(ctx context.Context, login, password string) (*domain.User, error)
	Register(ctx context.Context, email, login, password string) error
	GetCountRows(ctx context.Context, TableName string) (int, error)

	// User CRUD operations
	GetUserByID(ctx context.Context, userId int64) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetAllUsers(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.User, int, error)
	CreateUser(ctx context.Context, email, login, password string, studentId *int, role string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, userId int64) error
	DeleteUsers(ctx context.Context, ids []int64) error

	FindStudentByID(ctx context.Context, id uint64) (*domain.Student, error)
	GetAllStudent(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Student, int, error)
	CreateStudent(ctx context.Context, id_num_student uint64, name_group, email_student, second_name_student,
		first_name_student, surname_student string) (*domain.Student, error)
	UpdateStudentbyID(ctx context.Context, id_num_student uint64, name_group, email_student, second_name_student,
		first_name_student, surname_student string) (*domain.Student, error)
	DeleteStudentbyID(ctx context.Context, id_num_student uint64) error

	GetAllGroup(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Group, int, error)
	UpdateGroupbyName(ctx context.Context, group_name, Studies_direction_group,
		Studies_profile_group string, Start_date_group mytype.JsonDate, Studies_period_group uint8) (*domain.Group, error)
	CreateGroup(ctx context.Context, group_name, Studies_direction_group,
		Studies_profile_group string, Start_date_group mytype.JsonDate, Studies_period_group uint8) (*domain.Group, error)
	DeleteGroupByName(ctx context.Context, group_name string) error

	GetAllMark(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Mark, int, error)
	CreateMark(ctx context.Context, id_mark, id_num_student int64,
		name_semester, lesson_name_mark string, score_mark int8, type_mark string, type_exam string) (*domain.Mark, error)
	UpdateMarkByID(ctx context.Context, id_mark, id_num_student int64,
		name_semester, lesson_name_mark string, score_mark int8, type_mark string, type_exam string) (*domain.Mark, error)
	DeleteMarkByID(ctx context.Context, id_mark int64) error

	// Scholarship operations
	GetAllScholarship(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Scholarship, int, error)
	CreateScholarship(ctx context.Context, id_num_student int64,
		name_semester string, size_scholarshp float64, id_budget int64) (*domain.Scholarship, error)
	UpdateScholarshipByID(ctx context.Context, id_scholarship, id_num_student int64,
		name_semester string, size_scholarshp float64, id_budget int64) (*domain.Scholarship, error)
	DeleteScholarshipByID(ctx context.Context, id_scholarship int64) error

	// Semester operations
	GetAllSemester(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Semester, int, error)
	CreateSemester(ctx context.Context, name_semester string, date_start_semester, date_end_semester string) (*domain.Semester, error)
	UpdateSemesterByName(ctx context.Context, name_semester string, date_start_semester, date_end_semester string) (*domain.Semester, error)
	DeleteSemesterByName(ctx context.Context, name_semester string) error
	DeleteSemesters(ctx context.Context, names []string) error

	// Budget operations
	GetAllBudget(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Budget, int, error)
	CreateBudget(ctx context.Context, type_scholarship_budget string, name_semester string, size_budget float64) (*domain.Budget, error)
	UpdateBudgetByID(ctx context.Context, id_budget int, type_scholarship_budget string, name_semester string, size_budget float64) (*domain.Budget, error)
	DeleteBudgetByID(ctx context.Context, id_budget int) error
	DeleteBudgets(ctx context.Context, ids []int) error

	// Achievement Category operations
	GetAllAchievementCategory(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.AchievementCategory, int, error)
	CreateAchievementCategory(ctx context.Context, achivments_type_category string, score_category uint8) (*domain.AchievementCategory, error)
	UpdateAchievementCategoryByID(ctx context.Context, id_category uint64, achivments_type_category string, score_category uint8) (*domain.AchievementCategory, error)
	DeleteAchievementCategoryByID(ctx context.Context, id_category uint64) error
	DeleteAchievementCategories(ctx context.Context, ids []uint64) error

	// Achievement operations
	GetAllAchievement(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Achievement, int, error)
	CreateAchievement(ctx context.Context, id_num_student int64,
		id_category int64, name_achivement string, date_achivment string) (*domain.Achievement, error)
	UpdateAchievementByID(ctx context.Context, id_achivment, id_num_student int64,
		id_category int64, name_achivement string, date_achivment string) (*domain.Achievement, error)
	DeleteAchievementByID(ctx context.Context, id_achivment int64) error
	DeleteAchievements(ctx context.Context, ids []int64) error
}
