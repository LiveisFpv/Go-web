package handlers

import (
	"backend/internal/app"
	"backend/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateGroup(c *gin.Context, a *app.App) {
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

func UpdateGroupbyName(c *gin.Context, a *app.App) {
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

func DeleteGroupbyName(c *gin.Context, a *app.App) {
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

func DeleteGroups(c *gin.Context, a *app.App) {
	var reqBody GroupDeleteRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	for _, group_name := range reqBody.Group_names {
		err := a.DeleteGroupByName(c, group_name)
		if err != nil {
			c.JSON(http.StatusForbidden, ErrorResponse(err))
			return
		}
	}
	c.JSON(http.StatusOK, GroupSuccessResponse(&domain.Group{}))
}

func GetAllGroup(c *gin.Context, a *app.App) {
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
	groups, count, err := a.GetAllGroup(c, filters, rowCount, page, search)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AllGroupSuccessResponse(groups, rowCount, count, page))
}
