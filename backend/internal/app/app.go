package app

import (
	"backend/internal/crypt"
	"backend/internal/domain"
	"backend/internal/mytype"
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

func (a *App) GetCountRows(ctx context.Context, tableName string) (int, error) {
	count, err := a.repo.GetCountRows(ctx, tableName)
	return count, err
}
func (a *App) Login(ctx context.Context, login, password string) (bool, error) {
	user, err := a.repo.Login(ctx, login, password)
	if err != nil {
		return false, err
	}
	return crypt.CheckPassword(user.Password, password), nil
}
func (a *App) Register(ctx context.Context, email, login, password string) error {
	hashedPassword, err := crypt.HashPassword(password)
	if err != nil {
		return err
	}
	err = a.repo.Register(ctx, email, login, hashedPassword)
	return err
}

// Получение студента по ID
func (a *App) GetStudentbyID(ctx context.Context, student_id uint64) (*domain.Student, error) {
	student, err := a.repo.FindStudentByID(ctx, student_id)
	return student, err
}

// Получение всех студентов из таблицы
func (a *App) GetAllStudent(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Student, int, error) {
	students, count, err := a.repo.GetAllStudent(ctx, filters, rowCount, page, search)
	return students, count, err
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
func (a *App) GetAllGroup(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Group, int, error) {
	groups, count, err := a.repo.GetAllGroup(ctx, filters, rowCount, page, search)
	return groups, count, err
}

func (a *App) CreateGroup(ctx context.Context, group_name, Studies_direction_group,
	Studies_profile_group string, Start_date_group mytype.JsonDate, Studies_period_group uint8) (*domain.Group, error) {
	group, err := a.repo.CreateGroup(ctx,
		group_name,
		Studies_direction_group,
		Studies_profile_group,
		Start_date_group,
		Studies_period_group)
	return group, err
}

// Обновление группы по имени
func (a *App) UpdateGroupbyName(ctx context.Context, group_name, Studies_direction_group,
	Studies_profile_group string, Start_date_group mytype.JsonDate, Studies_period_group uint8) (*domain.Group, error) {
	group, err := a.repo.UpdateGroupbyName(ctx,
		group_name,
		Studies_direction_group,
		Studies_profile_group,
		Start_date_group,
		Studies_period_group,
	)
	return group, err
}

// Удаление группы по имени
func (a *App) DeleteGroupByName(ctx context.Context, group_name string) error {
	err := a.repo.DeleteGroupByName(ctx, group_name)
	return err
}

func (a *App) GetAllMark(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Mark, int, error) {
	marks, count, err := a.repo.GetAllMark(ctx, filters, rowCount, page, search)
	return marks, count, err
}

func (a *App) CreateMark(ctx context.Context, id_mark, id_num_student int64,
	name_semester, lesson_name_mark string, score_mark int8, type_mark string) (*domain.Mark, error) {
	mark, err := a.repo.CreateMark(ctx,
		id_mark,
		id_num_student,
		name_semester,
		lesson_name_mark,
		score_mark,
		type_mark,
	)
	return mark, err
}
func (a *App) UpdateMarkByID(ctx context.Context, id_mark, id_num_student int64,
	name_semester, lesson_name_mark string, score_mark int8, type_mark string) (*domain.Mark, error) {
	mark, err := a.repo.UpdateMarkByID(ctx,
		id_mark,
		id_num_student,
		name_semester,
		lesson_name_mark,
		score_mark,
		type_mark,
	)
	return mark, err
}

func (a *App) DeleteMarkByID(ctx context.Context, id_mark int64) error {
	err := a.repo.DeleteMarkByID(ctx, id_mark)
	return err
}

// Scholarship operations
func (a *App) GetAllScholarship(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Scholarship, int, error) {
	scholarships, count, err := a.repo.GetAllScholarship(ctx, filters, rowCount, page, search)
	return scholarships, count, err
}

func (a *App) CreateScholarship(ctx context.Context, id_num_student int64,
	name_semester string, size_scholarshp float64, id_budget int64) (*domain.Scholarship, error) {
	scholarship, err := a.repo.CreateScholarship(ctx,
		id_num_student,
		name_semester,
		size_scholarshp,
		id_budget,
	)
	return scholarship, err
}

func (a *App) UpdateScholarshipByID(ctx context.Context, id_scholarship, id_num_student int64,
	name_semester string, size_scholarshp float64, id_budget int64) (*domain.Scholarship, error) {
	scholarship, err := a.repo.UpdateScholarshipByID(ctx,
		id_scholarship,
		id_num_student,
		name_semester,
		size_scholarshp,
		id_budget,
	)
	return scholarship, err
}

func (a *App) DeleteScholarshipByID(ctx context.Context, id_scholarship int64) error {
	err := a.repo.DeleteScholarshipByID(ctx, id_scholarship)
	return err
}

// Semester operations
func (a *App) GetAllSemester(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Semester, int, error) {
	semesters, count, err := a.repo.GetAllSemester(ctx, filters, rowCount, page, search)
	return semesters, count, err
}

func (a *App) CreateSemester(ctx context.Context, name_semester string, date_start_semester, date_end_semester string) (*domain.Semester, error) {
	semester, err := a.repo.CreateSemester(ctx, name_semester, date_start_semester, date_end_semester)
	return semester, err
}

func (a *App) UpdateSemesterByName(ctx context.Context, name_semester string, date_start_semester, date_end_semester string) (*domain.Semester, error) {
	semester, err := a.repo.UpdateSemesterByName(ctx, name_semester, date_start_semester, date_end_semester)
	return semester, err
}

func (a *App) DeleteSemesterByName(ctx context.Context, name_semester string) error {
	err := a.repo.DeleteSemesterByName(ctx, name_semester)
	return err
}

func (a *App) DeleteSemesters(ctx context.Context, names []string) error {
	err := a.repo.DeleteSemesters(ctx, names)
	return err
}
