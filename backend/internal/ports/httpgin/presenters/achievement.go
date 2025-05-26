package presenters

import (
	"backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type achievementResponse struct {
	Id_achivment    int64  `json:"id_achivment"`
	Id_num_student  int64  `json:"id_num_student"`
	Id_category     int64  `json:"id_category"`
	Name_achivement string `json:"name_achivement"`
	Date_achivment  string `json:"date_achivement"`
	// Student information
	Student_surname     string `json:"surname_student"`
	Student_name        string `json:"first_name_student"`
	Student_second_name string `json:"second_name_student"`
	// Category information
	Category_type string `json:"achivments_type_category"`
	// Group information
	Name_group string `json:"name_group"`
}

type AchievementRequest struct {
	Id_achivment    int64  `json:"id_achivment"`
	Id_num_student  int64  `json:"id_num_student"`
	Id_category     int64  `json:"id_category"`
	Name_achivement string `json:"name_achivement"`
	Date_achivment  string `json:"date_achivement"`
}

type AchievementsDeleteRequest struct {
	Ids_achievement []string `json:"ids"`
}

func mapAchievementToResponse(achievement *domain.Achievement) achievementResponse {
	return achievementResponse{
		Id_achivment:        achievement.Id_achivment,
		Id_num_student:      achievement.Id_num_student,
		Id_category:         achievement.Id_category,
		Name_achivement:     achievement.Name_achivement,
		Date_achivment:      achievement.Date_achivment,
		Student_surname:     achievement.Student_surname,
		Student_name:        achievement.Student_name,
		Student_second_name: achievement.Student_second_name,
		Category_type:       achievement.Category_type,
		Name_group:          achievement.Name_group,
	}
}

func AchievementSuccessResponse(achievement *domain.Achievement) *gin.H {
	return SuccessResponse(mapAchievementToResponse(achievement))
}

func AllAchievementSuccessResponse(achievements []*domain.Achievement, countRow, count, page int) *gin.H {
	data := Paginate(achievements, countRow, page, mapAchievementToResponse)
	return AllSuccessResponse(data, Pagination{
		Total:     count,
		Page:      page,
		Page_size: countRow,
	})
}
