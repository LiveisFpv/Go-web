package handlers

import (
	"backend/internal/app"
	"backend/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStudentbyID(c *gin.Context, a *app.App) {
	studentId, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	student, err := a.GetStudentbyID(c, uint64(studentId))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, StudentSuccessResponse(student))
}

func CreateStudent(c *gin.Context, a *app.App) {
	var reqBody StudentRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	student, err := a.CreateStudent(c,
		reqBody.Id_num_student,
		reqBody.Name_group,
		reqBody.Email_student,
		reqBody.Second_name_student,
		reqBody.First_name_student,
		reqBody.Surname_student)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, StudentSuccessResponse(student))
}

func UpdateStudentbyID(c *gin.Context, a *app.App) {
	var reqBody StudentRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	student, err := a.UpdateStudentbyID(c,
		reqBody.Id_num_student,
		reqBody.Name_group,
		reqBody.Email_student,
		reqBody.Second_name_student,
		reqBody.First_name_student,
		reqBody.Surname_student)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, StudentSuccessResponse(student))
}

func DeleteStudentbyID(c *gin.Context, a *app.App) {
	var reqBody StudentRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	err := a.DeleteStudentbyID(c, reqBody.Id_num_student)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, StudentSuccessResponse(&domain.Student{}))
}

func DeleteStudents(c *gin.Context, a *app.App) {
	var reqBody StudentDeleteRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	for _, id := range reqBody.Ids_num_student {
		id_num_student, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}
		err = a.DeleteStudentbyID(c, uint64(id_num_student))
		if err != nil {
			c.JSON(http.StatusForbidden, ErrorResponse(err))
			return
		}
	}
	c.JSON(http.StatusOK, StudentSuccessResponse(&domain.Student{}))
}

func GetAllStudent(c *gin.Context, a *app.App) {
	param := c.Query("page")
	page, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	param = c.Query("limit")
	rowCount, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	search := c.Query("search")
	// Собираем фильтры (все параметры, кроме page)
	filters := make(map[string]string)
	for key, value := range c.Request.URL.Query() {
		if key != "page" && key != "search" && key != "limit" {
			if len(value[0]) > 0 {
				filters[key] = value[0] // берем первое значение фильтра
			}
		}
	}
	students, count, err := a.GetAllStudent(c, filters, rowCount, page, search)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AllStudentSuccessResponse(students, rowCount, count, page))
}
