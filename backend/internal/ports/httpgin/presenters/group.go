package presenters

import (
	"backend/internal/domain"
	"backend/internal/mytype"

	"github.com/gin-gonic/gin"
)

type GroupDeleteRequest struct {
	Group_names []string `json:"ids"`
}

type groupResponse struct {
	Name_group              string          `json:"name_group"`
	Studies_direction_group string          `json:"studies_direction_group"`
	Studies_profile_group   string          `json:"studies_profile_group"`
	Start_date_group        mytype.JsonDate `json:"start_date_group"`
	Studies_period_group    uint8           `json:"studies_period_group"`
}

type GroupRequest struct {
	Name_group              string          `json:"name_group"`
	Studies_direction_group string          `json:"studies_direction_group"`
	Studies_profile_group   string          `json:"studies_profile_group"`
	Start_date_group        mytype.JsonDate `json:"start_date_group"`
	Studies_period_group    uint8           `json:"studies_period_group"`
}

func mapGroupToResponse(group *domain.Group) groupResponse {
	return groupResponse{
		Name_group:              group.Name_group,
		Studies_direction_group: group.Studies_direction_group,
		Studies_profile_group:   group.Studies_profile_group,
		Start_date_group:        group.Start_date_group,
		Studies_period_group:    group.Studies_period_group,
	}
}

func GroupSuccessResponse(group *domain.Group) *gin.H {
	return SuccessResponse(mapGroupToResponse(group))
}

func AllGroupSuccessResponse(groups []*domain.Group, countRow, count, page int) *gin.H {
	data := Paginate(groups, countRow, page, mapGroupToResponse)
	return AllSuccessResponse(data, Pagination{
		Total:     count,
		Page:      page,
		Page_size: countRow,
	})
}
