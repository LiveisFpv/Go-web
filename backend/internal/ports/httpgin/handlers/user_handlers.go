package handlers

import (
	"backend/internal/app"
	"backend/internal/domain"
	"backend/internal/ports/httpgin/presenters"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context, a *app.App) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	user, err := a.GetUserByID(c, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.UserSuccessResponse(user))
}

func CreateUser(c *gin.Context, a *app.App) {
	var reqBody presenters.UserRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	user, err := a.CreateUser(c,
		reqBody.Email,
		reqBody.Login,
		reqBody.Password,
		reqBody.Student_id,
		*reqBody.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, presenters.UserSuccessResponse(user))
}

func UpdateUser(c *gin.Context, a *app.App) {
	var reqBody presenters.UserRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	user := &domain.User{
		Id:         reqBody.Id,
		Email:      reqBody.Email,
		Login:      reqBody.Login,
		Password:   reqBody.Password,
		Student_id: reqBody.Student_id,
		Role:       reqBody.Role,
	}
	err := a.UpdateUser(c, user)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.UserSuccessResponse(user))
}

func DeleteUser(c *gin.Context, a *app.App) {
	var reqBody presenters.UserRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	err := a.DeleteUser(c, reqBody.Id)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.UserSuccessResponse(&domain.User{}))
}

func DeleteUsers(c *gin.Context, a *app.App) {
	var reqBody presenters.UserDeleteRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	err := a.DeleteUsers(c, reqBody.Ids)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.UserSuccessResponse(&domain.User{}))
}

func GetAllUsers(c *gin.Context, a *app.App) {
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
	// Collect filters (all parameters except page)
	filters := make(map[string]string)
	for key, value := range c.Request.URL.Query() {
		if key != "page" && key != "search" && key != "limit" {
			if len(value[0]) > 0 {
				filters[key] = value[0]
			}
		}
	}
	users, count, err := a.GetAllUsers(c, filters, rowCount, page, search)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.AllUserSuccessResponse(users, rowCount, count, page))
}

func GetUserByEmail(c *gin.Context, a *app.App) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse(fmt.Errorf("email parameter is required")))
		return
	}
	user, err := a.GetUserByEmail(c, email)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, presenters.UserSuccessResponse(user))
}
