package presenters

import (
	"backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type markResponce struct {
	Id_mark          int64  `json:"id_mark"`
	Id_num_student   int64  `json:"id_num_student"`
	Name_semester    string `json:"name_semester"`
	Lesson_name_mark string `json:"lesson_name_mark"`
	Score_mark       int8   `json:"score_mark"`
	Type_mark        string `json:"type_mark"`
	Type_exam        string `json:"type_exam"`
	// Student information
	Student_surname     string `json:"surname_student"`
	Student_name        string `json:"first_name_student"`
	Student_second_name string `json:"second_name_student"`
	Student_group       string `json:"name_group"`
}

type MarkRequest struct {
	Id_mark          int64  `json:"id_mark"`
	Id_num_student   int64  `json:"id_num_student"`
	Name_semester    string `json:"name_semester"`
	Lesson_name_mark string `json:"lesson_name_mark"`
	Score_mark       int8   `json:"score_mark"`
	Type_mark        string `json:"type_mark"`
	Type_exam        string `json:"type_exam"`
}

type MarksDeleteRequest struct {
	Ids_mark []string `json:"ids"`
}

func mapMarkToResponse(mark *domain.Mark) markResponce {
	return markResponce{
		Id_mark:             mark.Id_mark,
		Id_num_student:      mark.Id_num_student,
		Name_semester:       mark.Name_semester,
		Lesson_name_mark:    mark.Lesson_name_mark,
		Score_mark:          mark.Score_mark,
		Type_mark:           mark.Type_mark,
		Type_exam:           mark.Type_exam,
		Student_surname:     mark.Student_surname,
		Student_name:        mark.Student_name,
		Student_second_name: mark.Student_second_name,
		Student_group:       mark.Student_group,
	}
}

func MarkSuccessResponse(mark *domain.Mark) *gin.H {
	return SuccessResponse(mapMarkToResponse(mark))
}

func AllMarkSuccessResponse(marks []*domain.Mark, countRow, count, page int) *gin.H {
	data := Paginate(marks, countRow, page, mapMarkToResponse)
	return AllSuccessResponse(data, Pagination{
		Total:     count,
		Page:      page,
		Page_size: countRow,
	})
}
