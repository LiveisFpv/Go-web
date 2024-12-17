package httpgin

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
type createStudentRequest struct {
	Id_num_student      uint64 `json:"id_num_student"`
	Name_group          string `json:"name_group"`
	Email_student       string `json:"email_student"`
	Second_name_student string `json:"second_name_student"`
	First_name_student  string `json:"first_name_student"`
	Surname_student     string `json:"surname_student"`
}
type updateStudentRequest struct {
	Name_group          string `json:"name_group"`
	Email_student       string `json:"email_student"`
	Second_name_student string `json:"second_name_student"`
	First_name_student  string `json:"first_name_student"`
	Surname_student     string `json:"surname_student"`
}

func StudentSuccessResponse(student *domain.Student) *gin.H {
	return &gin.H{
		"data": studentResponse{
			Id_num_student:      student.Id_num_student,
			Name_group:          student.Name_group,
			Email_student:       student.Email_student,
			Second_name_student: student.Second_name_student,
			First_name_student:  student.First_name_student,
			Surname_student:     student.Surname_student,
		},
		"error": nil,
	}
}
func AllStudentSuccessResponse(students []*domain.Student) *gin.H {
	data := []studentResponse{}
	for _, student := range students {
		data = append(data, studentResponse{
			Id_num_student:      student.Id_num_student,
			Name_group:          student.Name_group,
			Email_student:       student.Email_student,
			Second_name_student: student.Second_name_student,
			First_name_student:  student.First_name_student,
			Surname_student:     student.Surname_student,
		})
	}
	return &gin.H{
		"data":  data,
		"error": nil,
	}
}
func StudentErrorResponse(err error) *gin.H {
	return &gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}
