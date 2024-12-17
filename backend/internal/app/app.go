package app

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"context"
)

type App struct {
	repo repository.Repository
}

// Конструктор приложения, который принимает репозиторий
func NewApp(repo repository.Repository) *App {
	return &App{repo: repo}
}

func (a *App) GetStudentbyID(ctx context.Context, student_id uint64) (*domain.Student, error) {
	student, err := a.repo.FindStudentByID(ctx, student_id)
	return student, err
}
func (a *App) GetAllStudent(ctx context.Context) ([]*domain.Student, error) {
	students, err := a.repo.GetAllStudent(ctx)
	return students, err
}

func (a *App) CreateStudent(ctx context.Context, id_num_student uint64, name_group, email_student,
	second_name_student, first_name_student, surname_student string) (*domain.Student, error) {
	student, err := a.repo.CreateStudent(ctx, id_num_student, name_group, email_student, second_name_student,
		first_name_student, surname_student)
	return student, err
}
func (a *App) UpdateStudentbyID(ctx context.Context, id_num_student uint64, name_group, email_student,
	second_name_student, first_name_student, surname_student string) (*domain.Student, error) {
	student, err := a.repo.UpdateStudentbyID(ctx, id_num_student, name_group, email_student, second_name_student,
		first_name_student, surname_student)
	return student, err
}
func (a *App) DeleteStudentbyID(ctx context.Context, id_num_student uint64) error {
	err := a.repo.DeleteStudentbyID(ctx, id_num_student)
	return err
}
