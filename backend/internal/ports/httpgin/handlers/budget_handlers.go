package handlers

import (
	"backend/internal/app"
	"backend/internal/domain"
	"backend/internal/ports/httpgin/presenters"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBudget(c *gin.Context, a *app.App) {
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
	budgets, count, err := a.GetAllBudget(c, filters, rowCount, page, search)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.AllBudgetSuccessResponse(budgets, rowCount, count, page))
}

func CreateBudget(c *gin.Context, a *app.App) {
	var reqBody presenters.BudgetRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	budget, err := a.CreateBudget(c,
		reqBody.Type_scholarship_budget,
		reqBody.Name_semester,
		reqBody.Size_budget)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, presenters.BudgetSuccessResponse(budget))
}

func UpdateBudgetByID(c *gin.Context, a *app.App) {
	var reqBody presenters.BudgetRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	budget, err := a.UpdateBudgetByID(c,
		reqBody.Id_budget,
		reqBody.Type_scholarship_budget,
		reqBody.Name_semester,
		reqBody.Size_budget)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.BudgetSuccessResponse(budget))
}

func DeleteBudgetByID(c *gin.Context, a *app.App) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	err = a.DeleteBudgetByID(c, id)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.BudgetSuccessResponse(&domain.Budget{}))
}

func DeleteBudgets(c *gin.Context, a *app.App) {
	var reqBody presenters.BudgetDeleteRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	err := a.DeleteBudgets(c, reqBody.Ids)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.BudgetSuccessResponse(&domain.Budget{}))
}
