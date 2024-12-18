package app

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"context"
	"time"
)

type App struct {
	repo repository.Repository
}

// Конструктор приложения, который принимает репозиторий
func NewApp(repo repository.Repository) *App {
	return &App{repo: repo}
}

// Получение студента по ID
func (a *App) GetStudentbyID(ctx context.Context, student_id uint64) (*domain.Student, error) {
	student, err := a.repo.FindStudentByID(ctx, student_id)
	return student, err
}

// Получение всех студентов из таблицы
func (a *App) GetAllStudent(ctx context.Context) ([]*domain.Student, error) {
	students, err := a.repo.GetAllStudent(ctx)
	return students, err
}

// Добавление нового студента в таблицу
func (a *App) CreateStudent(ctx context.Context, id_num_student uint64, name_group, email_student,
	second_name_student, first_name_student, surname_student string) (*domain.Student, error) {
	student, err := a.repo.CreateStudent(ctx, id_num_student, name_group, email_student, second_name_student,
		first_name_student, surname_student)
	return student, err
}

// Обновление студента по ID
func (a *App) UpdateStudentbyID(ctx context.Context, id_num_student uint64, name_group, email_student,
	second_name_student, first_name_student, surname_student string) (*domain.Student, error) {
	student, err := a.repo.UpdateStudentbyID(ctx, id_num_student, name_group, email_student, second_name_student,
		first_name_student, surname_student)
	return student, err
}

// Удаление студента по ID
func (a *App) DeleteStudentbyID(ctx context.Context, id_num_student uint64) error {
	err := a.repo.DeleteStudentbyID(ctx, id_num_student)
	return err
}

// Получение всех групп из таблицы
func (a *App) GetAllGroup(ctx context.Context) ([]*domain.Group, error) {
	groups, err := a.repo.GetAllGroup(ctx)
	return groups, err
}

// Обновление группы по имени
func (a *App) UpdateGroupbyName(ctx context.Context, group_name, Studies_direction_group,
	Studies_profile_group string, Start_date_group time.Time, Studies_period_group uint8) (*domain.Group, error) {
	group, err := a.repo.UpdateGroupbyName(ctx,
		group_name,
		Studies_direction_group,
		Studies_profile_group,
		Start_date_group,
		Studies_period_group,
	)
	return group, err
}
