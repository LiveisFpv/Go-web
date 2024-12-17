package httpgin

import (
	"fmt"
	"net/http"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"

	"homework8/internal/ads"
	"homework8/internal/app"
	"homework8/internal/user"
	"homework8/internal/validator"
)

// Метод для создания объявления (ad)
func createAd(c *gin.Context, a *app.App) {
	var reqBody createAdRequest
	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		return
	}
	err = validator.Validate(reqBody)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		return
	}
	var ad *ads.Ad
	ad, err = a.CreateAd(c, reqBody.Title, reqBody.Text, reqBody.UserID)

	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AdSuccessResponse(ad))
}
func getAdbyID(c *gin.Context, a *app.App) {
	adId, err := strconv.Atoi(c.Param("ad_id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		return
	}
	ad, err := a.GetAdByID(c, int64(adId))
	if err != nil {
		c.Status(http.StatusNotFound)
		c.JSON(http.StatusNotFound, AdErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AdSuccessResponse(ad))
}

// Метод для получения всех объявлений
func getListAd(c *gin.Context, a *app.App) {

	ads, err := a.GetAds(c)
	if err != nil {
		c.Status(http.StatusTeapot)
		c.JSON(http.StatusTeapot, AdErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, AdListSuccessResponse(ads))
}

// Метод для изменения статуса объявления (опубликовано - Published = true или снято с публикации Published = false)
func changeAdStatus(c *gin.Context, a *app.App) {
	var reqBody changeAdStatusRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		return
	}

	adID, err := strconv.Atoi(c.Param("ad_id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		return
	}
	var ad *ads.Ad
	ad, err = a.ChangeAdStatus(c, int64(adID), reqBody.UserID, reqBody.Published)

	if err != nil {
		c.Status(http.StatusForbidden)
		c.JSON(http.StatusForbidden, AdErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, AdSuccessResponse(ad))
}

func getAdbyName(c *gin.Context, a *app.App) {
	if c.Query("name") != "" {
		ads, err := a.GetAdsByName(c, c.Query("name"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, AdListSuccessResponse(ads))
		return
	}
	c.Status(http.StatusBadRequest)
	c.JSON(http.StatusBadRequest, AdErrorResponse(fmt.Errorf("name parameter is required")))
}

func getListAdfilted(c *gin.Context, a *app.App) {
	filter := c.Query("param")
	if filter == "" {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, AdErrorResponse(fmt.Errorf("param parameter is required")))
		return
	}
	Ads, err := a.GetAds(c)
	if err != nil {
		c.Status(http.StatusForbidden)
		c.JSON(http.StatusForbidden, AdErrorResponse(err))
		return
	}
	filteredAds := make([]ads.Ad, 0)
	if filter == "createdat" {
		for _, ad := range *Ads {
			if ad.CreatedAt.After(time.Now().AddDate(-1, 0, 0)) {
				if !ad.Published {
					continue
				}
				filteredAds = append(filteredAds, ad)
			}
		}
		c.JSON(http.StatusOK, AdListSuccessResponse(&filteredAds))
		return
	} else if filter == "updatedat" {
		for _, ad := range *Ads {
			if ad.UpdatedAt.After(time.Now().AddDate(-1, 0, 0)) {
				if !ad.Published {
					continue
				}
				filteredAds = append(filteredAds, ad)
			}
		}
		c.JSON(http.StatusOK, AdListSuccessResponse(&filteredAds))
		return
	}
	c.JSON(http.StatusNotFound, AdErrorResponse(fmt.Errorf("bad filter")))
}

// Метод для обновления текста(Text) или заголовка(Title) объявления
func updateAd(c *gin.Context, a *app.App) {
	var reqBody updateAdRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		return
	}
	err := validator.Validate(reqBody)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		return
	}
	adID, err := strconv.Atoi(c.Param("ad_id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		return
	}
	var ad *ads.Ad
	ad, err = a.UpdateAd(c, int64(adID), reqBody.Title, reqBody.Text, reqBody.UserID)
	// TODO: метод должен возвращать AdSuccessResponse или ошибку.
	if err != nil {
		c.Status(http.StatusForbidden)
		c.JSON(http.StatusForbidden, AdErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, AdSuccessResponse(ad))
}
func createUser(c *gin.Context, a *app.App) {
	var reqBody createUserRequest
	err := c.ShouldBind(&reqBody)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, UserErrorResponse(err))
		return
	}
	err = validator.Validate(reqBody)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, UserErrorResponse(err))
		return
	}
	var user *user.User
	user, err = a.CreateUser(c, reqBody.Nickname, reqBody.Email)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, UserErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, UserSuccessResponse(user))
}
func updateUser(c *gin.Context, a *app.App) {
	var reqBody updateUserRequest
	err := c.ShouldBind(&reqBody)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	err = validator.Validate(reqBody)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, UserErrorResponse(err))
		return
	}
	Id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, UserErrorResponse(err))
	}
	user, err := a.UpdateUser(c, int64(Id), reqBody.Nickname, reqBody.Email)
	if err != nil {
		c.Status(http.StatusForbidden)
		c.JSON(http.StatusForbidden, UserErrorResponse(err))
	}
	c.JSON(http.StatusOK, UserSuccessResponse(user))
}
func getUserAds(c *gin.Context, a *app.App) {
	Id, err := strconv.Atoi(c.Param("user_id"))
	id := int64(Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserErrorResponse(err))
	}
	_, err = a.GetUser(c, id)
	if err != nil {
		c.JSON(http.StatusForbidden, UserErrorResponse(err))
	}
	ads, err := a.GetUserListAds(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserErrorResponse(err))
	}
	c.JSON(http.StatusAccepted, AdListSuccessResponse(ads))
}
