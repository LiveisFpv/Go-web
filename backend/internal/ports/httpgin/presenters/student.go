package presenters

import (
	"backend/internal/domain"

	"github.com/gin-gonic/gin"
)

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
