package httpgin

import (
	"backend/internal/domain"
	"time"

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
func AllStudentSuccessResponse(students []*domain.Student, countRow, page int) *gin.H {
	data := []studentResponse{}
	for i, student := range students {
		if i >= (page)*countRow {
			break
		} else if i >= (page-1)*countRow {
			data = append(data, studentResponse{
				Id_num_student:      student.Id_num_student,
				Name_group:          student.Name_group,
				Email_student:       student.Email_student,
				Second_name_student: student.Second_name_student,
				First_name_student:  student.First_name_student,
				Surname_student:     student.Surname_student,
			})
		}
	}
	return &gin.H{
		"data":  data,
		"page":  page,
		"error": nil,
	}
}
func StudentErrorResponse(err error) *gin.H {
	return &gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}

type groupResponse struct {
	Name_group              string    `json:"name_group"`
	Studies_direction_group string    `json:"studies_direction_group"`
	Studies_profile_group   string    `json:"studies_profile_group"`
	Start_date_group        time.Time `json:"start_date_group"`
	Studies_period_group    uint8     `json:"studies_period_group"`
}

type GroupRequest struct {
	Name_group              string    `json:"name_group"`
	Studies_direction_group string    `json:"studies_direction_group"`
	Studies_profile_group   string    `json:"studies_profile_group"`
	Start_date_group        time.Time `json:"start_date_group"`
	Studies_period_group    uint8     `json:"studies_period_group"`
}

func GroupSuccessResponse(group *domain.Group) *gin.H {
	return &gin.H{
		"data": groupResponse{
			Name_group:              group.Name_group,
			Studies_direction_group: group.Studies_direction_group,
			Studies_profile_group:   group.Studies_profile_group,
			Start_date_group:        group.Start_date_group,
			Studies_period_group:    group.Studies_period_group,
		},
		"error": nil,
	}
}
func AllGroupSuccessResponse(group []*domain.Group, countRow, page int) *gin.H {
	data := []groupResponse{}
	for i, group := range group {
		if i >= (page)*countRow {
			break
		} else if i >= (page-1)*countRow {
			data = append(data, groupResponse{
				Name_group:              group.Name_group,
				Studies_direction_group: group.Studies_direction_group,
				Studies_profile_group:   group.Studies_profile_group,
				Start_date_group:        group.Start_date_group,
				Studies_period_group:    group.Studies_period_group,
			})
		}
	}
	return &gin.H{
		"data":  data,
		"error": nil,
	}
}
func GroupErrorResponse(err error) *gin.H {
	return &gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}
