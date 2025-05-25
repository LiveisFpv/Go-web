package presenters

import (
	"backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type scholarshipResponse struct {
	Id_scholarship  int64   `json:"id_scholarship"`
	Id_num_student  int64   `json:"id_num_student"`
	Name_semester   string  `json:"name_semester"`
	Size_scholarshp float64 `json:"size_scholarshp"`
	Id_budget       int64   `json:"id_budget"`
	// Student information
	Student_surname         string `json:"surname_student"`
	Student_name            string `json:"first_name_student"`
	Student_second_name     string `json:"second_name_student"`
	Student_group           string `json:"name_group"`
	Type_scholarship_budget string `json:"type_scholarship_budget"`
}

type ScholarshipRequest struct {
	Id_scholarship  int64   `json:"id_scholarship"`
	Id_num_student  int64   `json:"id_num_student"`
	Name_semester   string  `json:"name_semester"`
	Size_scholarshp float64 `json:"size_scholarshp"`
	Id_budget       int64   `json:"id_budget"`
}

type ScholarshipsDeleteRequest struct {
	Ids_scholarship []string `json:"ids"`
}

func mapScholarshipToResponse(scholarship *domain.Scholarship) scholarshipResponse {
	return scholarshipResponse{
		Id_scholarship:          scholarship.Id_scholarship,
		Id_num_student:          scholarship.Id_num_student,
		Name_semester:           scholarship.Name_semester,
		Size_scholarshp:         scholarship.Size_scholarshp,
		Id_budget:               scholarship.Id_budget,
		Student_surname:         scholarship.Student_surname,
		Student_name:            scholarship.Student_name,
		Student_second_name:     scholarship.Student_second_name,
		Student_group:           scholarship.Student_group,
		Type_scholarship_budget: scholarship.Type_scholarship_budget,
	}
}

func ScholarshipSuccessResponse(scholarship *domain.Scholarship) *gin.H {
	return SuccessResponse(mapScholarshipToResponse(scholarship))
}

func AllScholarshipSuccessResponse(scholarships []*domain.Scholarship, countRow, count, page int) *gin.H {
	data := Paginate(scholarships, countRow, page, mapScholarshipToResponse)
	return AllSuccessResponse(data, Pagination{
		Total:     count,
		Page:      page,
		Page_size: countRow,
	})
}
