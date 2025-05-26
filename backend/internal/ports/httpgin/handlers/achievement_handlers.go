package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend/internal/app"
	"backend/internal/domain"
)

func CreateAchievement(c *gin.Context, a *app.App) {
	var reqBody AchievementRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	achievement, err := a.CreateAchievement(c,
		reqBody.Id_num_student,
		reqBody.Id_category,
		reqBody.Name_achivement,
		reqBody.Date_achivment)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, AchievementSuccessResponse(achievement))
}

func UpdateAchievementByID(c *gin.Context, a *app.App) {
	var reqBody AchievementRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	achievement, err := a.UpdateAchievementByID(c,
		reqBody.Id_achivment,
		reqBody.Id_num_student,
		reqBody.Id_category,
		reqBody.Name_achivement,
		reqBody.Date_achivment)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AchievementSuccessResponse(achievement))
}

func DeleteAchievementByID(c *gin.Context, a *app.App) {
	var reqBody AchievementRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	err := a.DeleteAchievementByID(c, reqBody.Id_achivment)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AchievementSuccessResponse(&domain.Achievement{}))
}

func DeleteAchievements(c *gin.Context, a *app.App) {
	var reqBody AchievementsDeleteRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	ids := make([]int64, 0, len(reqBody.Ids_achievement))
	for _, id := range reqBody.Ids_achievement {
		id_achievement, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}
		ids = append(ids, id_achievement)
	}

	err := a.DeleteAchievements(c, ids)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AchievementSuccessResponse(&domain.Achievement{}))
}

func GetAllAchievement(c *gin.Context, a *app.App) {
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
	achievements, count, err := a.GetAllAchievement(c, filters, rowCount, page, search)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AllAchievementSuccessResponse(achievements, rowCount, count, page))
}
