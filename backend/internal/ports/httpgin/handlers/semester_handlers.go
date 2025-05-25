package handlers

import (
	"backend/internal/app"
	"backend/internal/domain"
	"backend/internal/ports/httpgin/presenters"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllSemester(c *gin.Context, a *app.App) {
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
	semesters, count, err := a.GetAllSemester(c, filters, rowCount, page, search)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.AllSemesterSuccessResponse(semesters, rowCount, count, page))
}

func CreateSemester(c *gin.Context, a *app.App) {
	var reqBody presenters.SemesterRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	semester, err := a.CreateSemester(c,
		reqBody.Name_semester,
		reqBody.Date_start_semester,
		reqBody.Date_end_semester)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, presenters.SemesterSuccessResponse(semester))
}

func UpdateSemesterByName(c *gin.Context, a *app.App) {
	var reqBody presenters.SemesterRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	semester, err := a.UpdateSemesterByName(c,
		reqBody.Name_semester,
		reqBody.Date_start_semester,
		reqBody.Date_end_semester)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.SemesterSuccessResponse(semester))
}

func DeleteSemesterByName(c *gin.Context, a *app.App) {
	var reqBody presenters.SemesterRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	err := a.DeleteSemesterByName(c, reqBody.Name_semester)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.SemesterSuccessResponse(&domain.Semester{}))
}

func DeleteSemesters(c *gin.Context, a *app.App) {
	var reqBody presenters.SemesterDeleteRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	err := a.DeleteSemesters(c, reqBody.Names)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.SemesterSuccessResponse(&domain.Semester{}))
}
