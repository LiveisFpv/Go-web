package presenters

import (
	"backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginSuccessResponse(token *domain.Token) *gin.H {
	return SuccessResponse(token)
}
