package presenters

import (
	"backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type achievementCategoryResponse struct {
	Id_category              uint64 `json:"id_category"`
	Achivments_type_category string `json:"achivments_type_category"`
	Score_category           uint8  `json:"score_category"`
}

type AchievementCategoryRequest struct {
	Id_category              uint64 `json:"id_category"`
	Achivments_type_category string `json:"achivments_type_category"`
	Score_category           uint8  `json:"score_category"`
}

type AchievementCategoryDeleteRequest struct {
	Ids []uint64 `json:"ids"`
}

// Преобразование структур
func mapAchievementCategoryToResponse(category *domain.AchievementCategory) achievementCategoryResponse {
	return achievementCategoryResponse{
		Id_category:              category.Id_category,
		Achivments_type_category: category.Achivments_type_category,
		Score_category:           category.Score_category,
	}
}

func AchievementCategorySuccessResponse(category *domain.AchievementCategory) *gin.H {
	return SuccessResponse(mapAchievementCategoryToResponse(category))
}

func AllAchievementCategorySuccessResponse(categories []*domain.AchievementCategory, countRow, count, page int) *gin.H {
	data := Paginate(categories, countRow, page, mapAchievementCategoryToResponse)
	return AllSuccessResponse(data, Pagination{
		Total:     count,
		Page:      page,
		Page_size: countRow,
	})
}
