package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend/internal/app"
	"backend/internal/domain"
)

func CreateScholarship(c *gin.Context, a *app.App) {
	var reqBody ScholarshipRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	scholarship, err := a.CreateScholarship(c,
		reqBody.Id_num_student,
		reqBody.Name_semester,
		reqBody.Size_scholarshp,
		reqBody.Id_budget)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, ScholarshipSuccessResponse(scholarship))
}

func UpdateScholarshipByID(c *gin.Context, a *app.App) {
	var reqBody ScholarshipRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	scholarship, err := a.UpdateScholarshipByID(c,
		reqBody.Id_scholarship,
		reqBody.Id_num_student,
		reqBody.Name_semester,
		reqBody.Size_scholarshp,
		reqBody.Id_budget)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, ScholarshipSuccessResponse(scholarship))
}

func DeleteScholarshipByID(c *gin.Context, a *app.App) {
	var reqBody ScholarshipRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	err := a.DeleteScholarshipByID(c, reqBody.Id_scholarship)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, ScholarshipSuccessResponse(&domain.Scholarship{}))
}

func DeleteScholarships(c *gin.Context, a *app.App) {
	var reqBody ScholarshipsDeleteRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	for _, id := range reqBody.Ids_scholarship {
		id_scholarship, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}
		err = a.DeleteScholarshipByID(c, int64(id_scholarship))
		if err != nil {
			c.JSON(http.StatusForbidden, ErrorResponse(err))
			return
		}
	}
	c.JSON(http.StatusOK, ScholarshipSuccessResponse(&domain.Scholarship{}))
}

func GetAllScholarship(c *gin.Context, a *app.App) {
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
	scholarships, count, err := a.GetAllScholarship(c, filters, rowCount, page, search)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AllScholarshipSuccessResponse(scholarships, rowCount, count, page))
}
