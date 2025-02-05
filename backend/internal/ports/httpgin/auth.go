package httpgin

import (
	"backend/internal/app"
	"backend/internal/crypt"
	"backend/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func auth(c *gin.Context, a *app.App) {
	var reqBody AuthRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	check, err := a.Login(c, reqBody.Login, reqBody.Password)
	if err != nil || !check {
		c.JSON(http.StatusUnauthorized, ErrorResponse(err))
		return
	}
	token, err := crypt.GenerateJWT(reqBody.Login)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, LoginSuccessResponse(&domain.Token{Token: token}))
}

func register(c *gin.Context, a *app.App) {
	var reqBody RegisterRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	err := a.Register(c, reqBody.Email, reqBody.Login, reqBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	flag, err := a.Login(c, reqBody.Login, reqBody.Password)
	if err != nil || !flag {
		c.JSON(http.StatusUnauthorized, ErrorResponse(err))
		return
	}
	token, err := crypt.GenerateJWT(reqBody.Login)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, LoginSuccessResponse(&domain.Token{Token: token}))
}
