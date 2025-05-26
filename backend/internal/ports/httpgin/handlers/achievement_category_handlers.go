package handlers

import (
	"backend/internal/app"
	"backend/internal/domain"
	"backend/internal/ports/httpgin/presenters"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllAchievementCategory(c *gin.Context, a *app.App) {
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
	categories, count, err := a.GetAllAchievementCategory(c, filters, rowCount, page, search)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.AllAchievementCategorySuccessResponse(categories, rowCount, count, page))
}

func CreateAchievementCategory(c *gin.Context, a *app.App) {
	var reqBody presenters.AchievementCategoryRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	category, err := a.CreateAchievementCategory(c,
		reqBody.Achivments_type_category,
		reqBody.Score_category)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, presenters.AchievementCategorySuccessResponse(category))
}

func UpdateAchievementCategoryByID(c *gin.Context, a *app.App) {
	var reqBody presenters.AchievementCategoryRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	category, err := a.UpdateAchievementCategoryByID(c,
		reqBody.Id_category,
		reqBody.Achivments_type_category,
		reqBody.Score_category)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.AchievementCategorySuccessResponse(category))
}

func DeleteAchievementCategoryByID(c *gin.Context, a *app.App) {
	var reqBody presenters.AchievementCategoryRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	err := a.DeleteAchievementCategoryByID(c, reqBody.Id_category)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.AchievementCategorySuccessResponse(&domain.AchievementCategory{}))
}

func DeleteAchievementCategories(c *gin.Context, a *app.App) {
	var reqBody presenters.AchievementCategoryDeleteRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	err := a.DeleteAchievementCategories(c, reqBody.Ids)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.AchievementCategorySuccessResponse(&domain.AchievementCategory{}))
}
