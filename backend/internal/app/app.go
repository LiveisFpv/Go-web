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
	name_semester, lesson_name_mark string, score_mark int8, type_mark string, type_exam string) (*domain.Mark, error) {
	mark, err := a.repo.CreateMark(ctx,
		id_mark,
		id_num_student,
		name_semester,
		lesson_name_mark,
		score_mark,
		type_mark,
		type_exam,
	)
	return mark, err
}
func (a *App) UpdateMarkByID(ctx context.Context, id_mark, id_num_student int64,
	name_semester, lesson_name_mark string, score_mark int8, type_mark string, type_exam string) (*domain.Mark, error) {
	mark, err := a.repo.UpdateMarkByID(ctx,
		id_mark,
		id_num_student,
		name_semester,
		lesson_name_mark,
		score_mark,
		type_mark,
		type_exam,
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

// Budget operations
func (a *App) GetAllBudget(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Budget, int, error) {
	budgets, count, err := a.repo.GetAllBudget(ctx, filters, rowCount, page, search)
	return budgets, count, err
}

func (a *App) CreateBudget(ctx context.Context, type_scholarship_budget string, name_semester string, size_budget float64) (*domain.Budget, error) {
	budget, err := a.repo.CreateBudget(ctx, type_scholarship_budget, name_semester, size_budget)
	return budget, err
}

func (a *App) UpdateBudgetByID(ctx context.Context, id_budget int, type_scholarship_budget string, name_semester string, size_budget float64) (*domain.Budget, error) {
	budget, err := a.repo.UpdateBudgetByID(ctx, id_budget, type_scholarship_budget, name_semester, size_budget)
	return budget, err
}

func (a *App) DeleteBudgetByID(ctx context.Context, id_budget int) error {
	err := a.repo.DeleteBudgetByID(ctx, id_budget)
	return err
}

func (a *App) DeleteBudgets(ctx context.Context, ids []int) error {
	err := a.repo.DeleteBudgets(ctx, ids)
	return err
}

// Achievement Category operations
func (a *App) GetAllAchievementCategory(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.AchievementCategory, int, error) {
	categories, count, err := a.repo.GetAllAchievementCategory(ctx, filters, rowCount, page, search)
	return categories, count, err
}

func (a *App) CreateAchievementCategory(ctx context.Context, achivments_type_category string, score_category uint8) (*domain.AchievementCategory, error) {
	category, err := a.repo.CreateAchievementCategory(ctx, achivments_type_category, score_category)
	return category, err
}

func (a *App) UpdateAchievementCategoryByID(ctx context.Context, id_category uint64, achivments_type_category string, score_category uint8) (*domain.AchievementCategory, error) {
	category, err := a.repo.UpdateAchievementCategoryByID(ctx, id_category, achivments_type_category, score_category)
	return category, err
}

func (a *App) DeleteAchievementCategoryByID(ctx context.Context, id_category uint64) error {
	err := a.repo.DeleteAchievementCategoryByID(ctx, id_category)
	return err
}

func (a *App) DeleteAchievementCategories(ctx context.Context, ids []uint64) error {
	err := a.repo.DeleteAchievementCategories(ctx, ids)
	return err
}

// Achievement operations
func (a *App) GetAllAchievement(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Achievement, int, error) {
	achievements, count, err := a.repo.GetAllAchievement(ctx, filters, rowCount, page, search)
	return achievements, count, err
}

func (a *App) CreateAchievement(ctx context.Context, id_num_student int64,
	id_category int64, name_achivement string, date_achivment string) (*domain.Achievement, error) {
	achievement, err := a.repo.CreateAchievement(ctx,
		id_num_student,
		id_category,
		name_achivement,
		date_achivment,
	)
	return achievement, err
}

func (a *App) UpdateAchievementByID(ctx context.Context, id_achivment, id_num_student int64,
	id_category int64, name_achivement string, date_achivment string) (*domain.Achievement, error) {
	achievement, err := a.repo.UpdateAchievementByID(ctx,
		id_achivment,
		id_num_student,
		id_category,
		name_achivement,
		date_achivment,
	)
	return achievement, err
}

func (a *App) DeleteAchievementByID(ctx context.Context, id_achivment int64) error {
	err := a.repo.DeleteAchievementByID(ctx, id_achivment)
	return err
}

func (a *App) DeleteAchievements(ctx context.Context, ids []int64) error {
	err := a.repo.DeleteAchievements(ctx, ids)
	return err
}

// User CRUD operations
func (a *App) GetUserByID(ctx context.Context, userId int64) (*domain.User, error) {
	user, err := a.repo.GetUserByID(ctx, userId)
	return user, err
}

func (a *App) GetAllUsers(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.User, int, error) {
	users, count, err := a.repo.GetAllUsers(ctx, filters, rowCount, page, search)
	return users, count, err
}

func (a *App) CreateUser(ctx context.Context, email, login, password string, studentId *int, role string) (*domain.User, error) {
	hashedPassword, err := crypt.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user, err := a.repo.CreateUser(ctx, email, login, hashedPassword, studentId, role)
	return user, err
}

func (a *App) DeleteUser(ctx context.Context, userId int64) error {
	err := a.repo.DeleteUser(ctx, userId)
	return err
}

func (a *App) DeleteUsers(ctx context.Context, ids []int64) error {
	err := a.repo.DeleteUsers(ctx, ids)
	return err
}

func (a *App) UpdateUser(ctx context.Context, user *domain.User) error {
	err := a.repo.UpdateUser(ctx, user)
	return err
}

func (a *App) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := a.repo.GetUserByEmail(ctx, email)
	return user, err
}
