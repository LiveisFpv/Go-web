package httpgin

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"

	"backend/internal/app"
)

func getStudentbyID(c *gin.Context, a *app.App) {
	studentId, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, StudentErrorResponse(err))
		return
	}
	student, err := a.GetStudentbyID(c, uint64(studentId))
	if err != nil {
		c.Status(http.StatusNotFound)
		c.JSON(http.StatusNotFound, StudentErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, StudentSuccessResponse(student))
}

func createStudent(c *gin.Context, a *app.App) {
	var reqBody createStudentRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, StudentErrorResponse(err))
		return
	}
	// err := validator.Validate(reqBody)
	// if err!= nil {
	//     c.Status(http.StatusBadRequest)
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
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, StudentErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, StudentSuccessResponse(student))
}
func updateStudentbyID(c *gin.Context, a *app.App) {
	studentId, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, StudentErrorResponse(err))
		return
	}
	var reqBody updateStudentRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, StudentErrorResponse(err))
		return
	}
	student, err := a.UpdateStudentbyID(c,
		uint64(studentId),
		reqBody.Name_group,
		reqBody.Email_student,
		reqBody.Second_name_student,
		reqBody.First_name_student,
		reqBody.Surname_student)
	if err != nil {
		c.Status(http.StatusForbidden)
		c.JSON(http.StatusForbidden, StudentErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, StudentSuccessResponse(student))
}

func deleteStudentbyID(c *gin.Context, a *app.App) {
	studentId, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, StudentErrorResponse(err))
		return
	}
	err = a.DeleteStudentbyID(c, uint64(studentId))
	if err != nil {
		c.Status(http.StatusForbidden)
		c.JSON(http.StatusForbidden, StudentErrorResponse(err))
		return
	}
	c.Status(http.StatusOK)
	c.JSON(http.StatusOK, StudentSuccessResponse(nil))
}
func getAllStudent(c *gin.Context, a *app.App) {
	students, err := a.GetAllStudent(c)
	if err != nil {
		c.Status(http.StatusForbidden)
		c.JSON(http.StatusForbidden, StudentErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AllStudentSuccessResponse(students))
}

func updateGroupbyName(c *gin.Context, a *app.App) {
	groupName := c.Param("group_name")
	var reqBody updateGroupRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, GroupErrorResponse(err))
		return
	}
	group, err := a.UpdateGroupbyName(c,
		groupName,
		reqBody.Studies_direction_group,
		reqBody.Studies_profile_group,
		reqBody.Start_date_group,
		reqBody.Studies_period_group,
	)
	if err != nil {
		c.Status(http.StatusForbidden)
		c.JSON(http.StatusForbidden, GroupErrorResponse(err))
		return
	}
	c.Status(http.StatusOK)
	c.JSON(http.StatusOK, GroupSuccessResponse(group))
}
func getAllGroup(c *gin.Context, a *app.App) {
	groups, err := a.GetAllGroup(c)
	if err != nil {
		c.Status(http.StatusForbidden)
		c.JSON(http.StatusForbidden, GroupErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AllGroupSuccessResponse(groups))
}
