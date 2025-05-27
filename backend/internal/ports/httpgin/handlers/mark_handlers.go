package handlers

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"

	"backend/internal/app"
	"backend/internal/domain"
)

func CreateMark(c *gin.Context, a *app.App) {
	var reqBody MarkRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	mark, err := a.CreateMark(c,
		reqBody.Id_mark,
		reqBody.Id_num_student,
		reqBody.Name_semester,
		reqBody.Lesson_name_mark,
		reqBody.Score_mark,
		reqBody.Type_mark,
		reqBody.Type_exam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, MarkSuccessResponse(mark))
}

func UpdateMarkbyID(c *gin.Context, a *app.App) {
	var reqBody MarkRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	mark, err := a.UpdateMarkByID(c,
		reqBody.Id_mark,
		reqBody.Id_num_student,
		reqBody.Name_semester,
		reqBody.Lesson_name_mark,
		reqBody.Score_mark,
		reqBody.Type_mark,
		reqBody.Type_exam)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, MarkSuccessResponse(mark))
}

func DeleteMarkbyID(c *gin.Context, a *app.App) {
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

func DeleteMarks(c *gin.Context, a *app.App) {
	var reqBody MarksDeleteRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	for _, id := range reqBody.Ids_mark {
		id_mark, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}
		err = a.DeleteMarkByID(c, int64(id_mark))
		if err != nil {
			c.JSON(http.StatusForbidden, ErrorResponse(err))
			return
		}
	}
	c.JSON(http.StatusOK, MarkSuccessResponse(&domain.Mark{}))
}

func GetAllMark(c *gin.Context, a *app.App) {
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
	marks, count, err := a.GetAllMark(c, filters, rowCount, page, search)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AllMarkSuccessResponse(marks, rowCount, count, page))
}
