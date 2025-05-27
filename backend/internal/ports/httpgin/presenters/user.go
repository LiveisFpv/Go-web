package presenters

import (
	"backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type userResponse struct {
	Id         int64   `json:"user_id"`
	Login      string  `json:"user_login"`
	Email      string  `json:"user_email"`
	Student_id *int    `json:"user_student_id"`
	Role       *string `json:"user_role"`
}

type UserRequest struct {
	Id         int64   `json:"user_id"`
	Login      string  `json:"user_login"`
	Email      string  `json:"user_email"`
	Password   string  `json:"user_password"`
	Student_id *int    `json:"user_student_id"`
	Role       *string `json:"user_role"`
}

type UserDeleteRequest struct {
	Ids []int64 `json:"ids"`
}

// Transform structures
func mapUserToResponse(user *domain.User) userResponse {
	return userResponse{
		Id:         user.Id,
		Login:      user.Login,
		Email:      user.Email,
		Student_id: user.Student_id,
		Role:       user.Role,
	}
}

func UserSuccessResponse(user *domain.User) *gin.H {
	return SuccessResponse(mapUserToResponse(user))
}

func AllUserSuccessResponse(users []*domain.User, countRow, count, page int) *gin.H {
	data := Paginate(users, countRow, page, mapUserToResponse)
	return AllSuccessResponse(data, Pagination{
		Total:     count,
		Page:      page,
		Page_size: countRow,
	})
}
