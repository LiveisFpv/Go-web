package httpgin

import (
	"backend/internal/domain"
	"backend/internal/mytype"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Total     int `json:"total"`
	Page      int `json:"page"`
	Page_size int `json:"page_size"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type studentResponse struct {
	Id_num_student      int64  `json:"id_num_student"`
	Name_group          string `json:"name_group"`
	Email_student       string `json:"email_student"`
	Second_name_student string `json:"second_name_student"`
	First_name_student  string `json:"first_name_student"`
	Surname_student     string `json:"surname_student"`
}
type StudentRequest struct {
	Id_num_student      uint64 `json:"id_num_student"`
	Name_group          string `json:"name_group"`
	Email_student       string `json:"email_student"`
	Second_name_student string `json:"second_name_student"`
	First_name_student  string `json:"first_name_student"`
	Surname_student     string `json:"surname_student"`
}
type StudentDeleteRequest struct {
	Ids_num_student []string `json:"ids"`
}
type GroupDeleteRequest struct {
	Group_names []string `json:"ids"`
}

type groupResponse struct {
	Name_group              string          `json:"name_group"`
	Studies_direction_group string          `json:"studies_direction_group"`
	Studies_profile_group   string          `json:"studies_profile_group"`
	Start_date_group        mytype.JsonDate `json:"start_date_group"`
	Studies_period_group    uint8           `json:"studies_period_group"`
}

type GroupRequest struct {
	Name_group              string          `json:"name_group"`
	Studies_direction_group string          `json:"studies_direction_group"`
	Studies_profile_group   string          `json:"studies_profile_group"`
	Start_date_group        mytype.JsonDate `json:"start_date_group"`
	Studies_period_group    uint8           `json:"studies_period_group"`
}
type markResponce struct {
	Id_mark          int64  `json:"id_mark"`
	Id_num_student   int64  `json:"id_num_student"`
	Name_semester    string `json:"name_semester"`
	Lesson_name_mark string `json:"lesson_name_mark"`
	Score_mark       int8   `json:"score_mark"`
	Type_mark        string `json:"type_mark"`
}

type MarkRequest struct {
	Id_mark          int64  `json:"id_mark"`
	Id_num_student   int64  `json:"id_num_student"`
	Name_semester    string `json:"name_semester"`
	Lesson_name_mark string `json:"lesson_name_mark"`
	Score_mark       int8   `json:"score_mark"`
	Type_mark        string `json:"type_mark"`
}

type MarksDeleteRequest struct {
	Ids_mark []string `json:"ids"`
}

// Общий интерфейс ответа
type ResponseData interface{}

// Общая функция для успешного ответа
func SuccessResponse(data ResponseData) *gin.H {
	return &gin.H{
		"data":  data,
		"error": nil,
	}
}

// Общая функция для ошибки
func ErrorResponse(err error) *gin.H {
	return &gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}
func AllSuccessResponse(data ResponseData, pagination Pagination) *gin.H {
	return &gin.H{
		"data":       data,
		"pagination": pagination,
		"error":      nil,
	}
}
func Filter[R any](items []R, filter func(R) bool) []R {
	result := make([]R, 0, len(items))
	for _, item := range items {
		if filter(item) {
			result = append(result, item)
		}
	}
	return result
}

// Универсальная функция для преобразования
func Paginate[T any, R any](items []T, countRow, page int, mapper func(T) R) []R {
	result := make([]R, 0, countRow)
	for _, item := range items {
		result = append(result, mapper(item))
	}
	return result
}

// Преобразование структур
func mapStudentToResponse(student *domain.Student) studentResponse {
	return studentResponse{
		Id_num_student:      student.Id_num_student,
		Name_group:          student.Name_group,
		Email_student:       student.Email_student,
		Second_name_student: student.Second_name_student,
		First_name_student:  student.First_name_student,
		Surname_student:     student.Surname_student,
	}
}

func mapGroupToResponse(group *domain.Group) groupResponse {
	return groupResponse{
		Name_group:              group.Name_group,
		Studies_direction_group: group.Studies_direction_group,
		Studies_profile_group:   group.Studies_profile_group,
		Start_date_group:        group.Start_date_group,
		Studies_period_group:    group.Studies_period_group,
	}
}

func mapMarkToResponse(mark *domain.Mark) markResponce {
	return markResponce{
		Id_mark:          mark.Id_mark,
		Id_num_student:   mark.Id_num_student,
		Name_semester:    mark.Name_semester,
		Lesson_name_mark: mark.Lesson_name_mark,
		Score_mark:       mark.Score_mark,
		Type_mark:        mark.Type_mark,
	}
}

func StudentSuccessResponse(student *domain.Student) *gin.H {
	return SuccessResponse(mapStudentToResponse(student))
}

func AllStudentSuccessResponse(students []*domain.Student, countRow, count, page int) *gin.H {
	data := Paginate(students, countRow, page, mapStudentToResponse)
	return AllSuccessResponse(data, Pagination{
		Total:     count,
		Page:      page,
		Page_size: countRow,
	})
}

func GroupSuccessResponse(group *domain.Group) *gin.H {
	return SuccessResponse(mapGroupToResponse(group))
}

func AllGroupSuccessResponse(groups []*domain.Group, countRow, count, page int) *gin.H {
	data := Paginate(groups, countRow, page, mapGroupToResponse)
	return AllSuccessResponse(data, Pagination{
		Total:     count,
		Page:      page,
		Page_size: countRow,
	})
}

func MarkSuccessResponse(mark *domain.Mark) *gin.H {
	return SuccessResponse(mapMarkToResponse(mark))
}
func LoginSuccessResponse(token *domain.Token) *gin.H {
	return SuccessResponse(token)
}

func AllMarkSuccessResponse(marks []*domain.Mark, countRow, count, page int) *gin.H {
	data := Paginate(marks, countRow, page, mapMarkToResponse)
	return AllSuccessResponse(data, Pagination{
		Total:     count,
		Page:      page,
		Page_size: countRow,
	})
}
