package httpgin

import (
	"backend/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a *app.App) {
	r.POST("/ads", func(c *gin.Context) { createAd(c, a) })                    // Метод для создания объявления (ad)
	r.PUT("/ads/:ad_id/status", func(c *gin.Context) { changeAdStatus(c, a) }) // Метод для изменения статуса объявления (опубликовано - Published = true или снято с публикации Published = false)
	r.PUT("/ads/:ad_id", func(c *gin.Context) { updateAd(c, a) })              // Метод для обновления текста(Text) или заголовка(Title) объявления
	r.GET("/ads/:ad_id", func(c *gin.Context) { getAdbyID(c, a) })
	r.GET("/ads", func(c *gin.Context) { getListAd(c, a) }) // Метод для получчения всех объявлений
	r.GET("/ads/filter", func(c *gin.Context) { getListAdfilted(c, a) })
	r.GET("/ads/search", func(c *gin.Context) { getAdbyName(c, a) })
	r.POST("/user", func(c *gin.Context) { createUser(c, a) })
	r.PUT("/user/:user_id", func(c *gin.Context) { updateUser(c, a) })
	r.GET("/user/:user_id/ads", func(c *gin.Context) { getUserAds(c, a) })
}
