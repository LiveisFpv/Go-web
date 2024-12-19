package httpgin

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"

	"backend/internal/app"
	"backend/internal/domain"
)

const rowCount = 10

func getStudentbyID(c *gin.Context, a *app.App) {
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

func createStudent(c *gin.Context, a *app.App) {
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

func updateStudentbyID(c *gin.Context, a *app.App) {
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

func deleteStudentbyID(c *gin.Context, a *app.App) {
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

func getAllStudent(c *gin.Context, a *app.App) {
	param := c.Query("page")
	page, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	students, err := a.GetAllStudent(c)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AllStudentSuccessResponse(students, rowCount, page))
}

func createGroup(c *gin.Context, a *app.App) {
	var reqBody GroupRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	group, err := a.CreateGroup(c,
		reqBody.Name_group,
		reqBody.Studies_direction_group,
		reqBody.Studies_profile_group,
		reqBody.Start_date_group,
		reqBody.Studies_period_group)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, GroupSuccessResponse(group))
}

func updateGroupbyName(c *gin.Context, a *app.App) {
	var reqBody GroupRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
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
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, GroupSuccessResponse(group))
}

func deleteGroupbyName(c *gin.Context, a *app.App) {
	var reqBody GroupRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	err := a.DeleteGroupByName(c, reqBody.Name_group)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, GroupSuccessResponse(&domain.Group{}))
}

func getAllGroup(c *gin.Context, a *app.App) {
	param := c.Query("page")
	page, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	groups, err := a.GetAllGroup(c)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AllGroupSuccessResponse(groups, rowCount, page))
}

func createMark(c *gin.Context, a *app.App) {
	var reqBody MarkRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	mark, err := a.CreateMark(c,
		reqBody.Id_mark,
		reqBody.Id_num_student,
		reqBody.Lesson_name_mark,
		reqBody.Name_semester,
		reqBody.Score_mark,
		reqBody.Type_mark)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, MarkSuccessResponse(mark))
}

func updateMarkbyID(c *gin.Context, a *app.App) {
	var reqBody MarkRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	mark, err := a.UpdateMarkByID(c,
		reqBody.Id_mark,
		reqBody.Id_num_student,
		reqBody.Lesson_name_mark,
		reqBody.Name_semester,
		reqBody.Score_mark,
		reqBody.Type_mark)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, MarkSuccessResponse(mark))
}

func deleteMarkbyID(c *gin.Context, a *app.App) {
	var reqBody MarkRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	err := a.DeleteMarkByID(c, reqBody.Id_mark)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, MarkSuccessResponse(&domain.Mark{}))
}

func getAllMark(c *gin.Context, a *app.App) {
	param := c.Query("page")
	page, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	marks, err := a.GetAllMark(c)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AllMarkSuccessResponse(marks, rowCount, page))
}
