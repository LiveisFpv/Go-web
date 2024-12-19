package httpgin

import (
	"backend/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a *app.App) {
	//Информация передоваемя на фронт для динамического поддержания работы
	r.GET("/tables", func(c *gin.Context) { getTables(c) })
	r.GET("/student/metadata", func(c *gin.Context) { getStudentMeta(c, a) })
	r.GET("/group/metadata", func(c *gin.Context) { getGroupMeta(c, a) })

	//Запросы на получение и редактирование таблиц
	r.GET("/student/:student_id", func(c *gin.Context) { getStudentbyID(c, a) })
	r.POST("/student/", func(c *gin.Context) { createStudent(c, a) })
	r.PUT("/student/", func(c *gin.Context) { updateStudentbyID(c, a) })
	r.DELETE("/student/", func(c *gin.Context) { deleteStudentbyID(c, a) })
	r.GET("/student/", func(c *gin.Context) { getAllStudent(c, a) })

	// r.GET("/group/:group_name", func(c *gin.Context) { getGroupbyID(c, a) })
	r.POST("/group/", func(c *gin.Context) { createGroup(c, a) })
	r.PUT("/group/", func(c *gin.Context) { updateGroupbyName(c, a) })
	r.DELETE("/group/", func(c *gin.Context) { deleteGroupbyName(c, a) })
	r.GET("/group/", func(c *gin.Context) { getAllGroup(c, a) })

}
