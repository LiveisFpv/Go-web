package httpgin

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"

	"backend/internal/app"
	"backend/internal/domain"
)

func getStudentbyID(c *gin.Context, a *app.App) {
	studentId, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, StudentErrorResponse(err))
		return
	}
	student, err := a.GetStudentbyID(c, uint64(studentId))
	if err != nil {
		c.JSON(http.StatusNotFound, StudentErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, StudentSuccessResponse(student))
}

func createStudent(c *gin.Context, a *app.App) {
	var reqBody StudentRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, StudentErrorResponse(err))
		return
	}
	// err := validator.Validate(reqBody)
	// if err!= nil {
	//     c.JSON(http.StatusBadRequest, StudentErrorResponse(err))
	//     return
	// }
	student, err := a.CreateStudent(c,
		reqBody.Id_num_student,
		reqBody.Name_group,
		reqBody.Email_student,
		reqBody.Second_name_student,
		reqBody.First_name_student,
		reqBody.Surname_student)
	if err != nil {
		c.JSON(http.StatusBadRequest, StudentErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, StudentSuccessResponse(student))
}

func updateStudentbyID(c *gin.Context, a *app.App) {
	var reqBody StudentRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, StudentErrorResponse(err))
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
		c.JSON(http.StatusForbidden, StudentErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, StudentSuccessResponse(student))
}

func deleteStudentbyID(c *gin.Context, a *app.App) {
	var reqBody StudentRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, StudentErrorResponse(err))
		return
	}
	err := a.DeleteStudentbyID(c, reqBody.Id_num_student)
	if err != nil {
		c.JSON(http.StatusForbidden, StudentErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, StudentSuccessResponse(&domain.Student{}))
}

func getAllStudent(c *gin.Context, a *app.App) {
	students, err := a.GetAllStudent(c)
	if err != nil {
		c.JSON(http.StatusForbidden, StudentErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AllStudentSuccessResponse(students))
}

func createGroup(c *gin.Context, a *app.App) {
	var reqBody GroupRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, GroupErrorResponse(err))
		return
	}
	group, err := a.CreateGroup(c,
		reqBody.Name_group,
		reqBody.Studies_direction_group,
		reqBody.Studies_profile_group,
		reqBody.Start_date_group,
		reqBody.Studies_period_group)
	if err != nil {
		c.JSON(http.StatusBadRequest, GroupErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, GroupSuccessResponse(group))
}

func updateGroupbyName(c *gin.Context, a *app.App) {
	var reqBody GroupRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, GroupErrorResponse(err))
		return
	}
	group, err := a.UpdateGroupbyName(c,
		reqBody.Name_group,
		reqBody.Studies_direction_group,
		reqBody.Studies_profile_group,
		reqBody.Start_date_group,
		reqBody.Studies_period_group,
	)
	if err != nil {
		c.JSON(http.StatusForbidden, GroupErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, GroupSuccessResponse(group))
}

func deleteGroupbyName(c *gin.Context, a *app.App) {
	var reqBody GroupRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, GroupErrorResponse(err))
		return
	}
	err := a.DeleteGroupByName(c, reqBody.Name_group)
	if err != nil {
		c.JSON(http.StatusForbidden, GroupErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, GroupSuccessResponse(&domain.Group{}))
}

func getAllGroup(c *gin.Context, a *app.App) {
	groups, err := a.GetAllGroup(c)
	if err != nil {
		c.JSON(http.StatusForbidden, GroupErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AllGroupSuccessResponse(groups))
}
