package presenters

import (
	"backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type budgetResponse struct {
	Type_scholarship_budget string  `json:"type_scholarship_budget"`
	Name_semester           string  `json:"name_semester"`
	Size_budget             float64 `json:"size_budget"`
	Id_budget               int     `json:"id_budget"`
}

type BudgetRequest struct {
	Id_budget               int     `json:"id_budget"`
	Type_scholarship_budget string  `json:"type_scholarship_budget"`
	Name_semester           string  `json:"name_semester"`
	Size_budget             float64 `json:"size_budget"`
}

type BudgetDeleteRequest struct {
	Ids []int `json:"ids"`
}

// Преобразование структур
func mapBudgetToResponse(budget *domain.Budget) budgetResponse {
	return budgetResponse{
		Type_scholarship_budget: budget.Type_scholarship_budget,
		Name_semester:           budget.Name_semester,
		Size_budget:             budget.Size_budget,
		Id_budget:               budget.Id_budget,
	}
}

func BudgetSuccessResponse(budget *domain.Budget) *gin.H {
	return SuccessResponse(mapBudgetToResponse(budget))
}

func AllBudgetSuccessResponse(budgets []*domain.Budget, countRow, count, page int) *gin.H {
	data := Paginate(budgets, countRow, page, mapBudgetToResponse)
	return AllSuccessResponse(data, Pagination{
		Total:     count,
		Page:      page,
		Page_size: countRow,
	})
}
