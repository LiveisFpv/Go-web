package httpgin

import (
	"backend/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a *app.App) {
	//Аутентификация и регистрация
	r.POST("/auth", func(c *gin.Context) { auth(c, a) })
	r.POST("/register", func(c *gin.Context) { register(c, a) })

	//Информация передоваемя на фронт для динамического поддержания работы
	r.GET("/tables", func(c *gin.Context) { getTables(c) })
	r.GET("/student/metadata", func(c *gin.Context) { getStudentMeta(c, a) })
	r.GET("/group/metadata", func(c *gin.Context) { getGroupMeta(c, a) })
	r.GET("/mark/metadata", func(c *gin.Context) { getMarkMeta(c, a) })

	//Запросы на получение и редактирование таблиц
	r.GET("/student/:student_id", func(c *gin.Context) { getStudentbyID(c, a) })
	r.POST("/student/", func(c *gin.Context) { createStudent(c, a) })
	r.PUT("/student/", func(c *gin.Context) { updateStudentbyID(c, a) })
	r.DELETE("/student/", func(c *gin.Context) { deleteStudentbyID(c, a) })
	r.GET("/student/", func(c *gin.Context) { getAllStudent(c, a) })
	r.DELETE("/student/ids", func(c *gin.Context) { deleteStudents(c, a) })

	// r.GET("/group/:group_name", func(c *gin.Context) { getGroupbyID(c, a) })
	r.POST("/group/", func(c *gin.Context) { createGroup(c, a) })
	r.PUT("/group/", func(c *gin.Context) { updateGroupbyName(c, a) })
	r.DELETE("/group/", func(c *gin.Context) { deleteGroupbyName(c, a) })
	r.GET("/group/", func(c *gin.Context) { getAllGroup(c, a) })
	r.DELETE("/group/ids", func(c *gin.Context) { deleteGroups(c, a) })

	r.POST("/mark/", func(c *gin.Context) { createMark(c, a) })
	r.PUT("/mark/", func(c *gin.Context) { updateMarkbyID(c, a) })
	r.DELETE("/mark/", func(c *gin.Context) { deleteMarkbyID(c, a) })
	r.GET("/mark/", func(c *gin.Context) { getAllMark(c, a) })
	r.DELETE("/mark/ids", func(c *gin.Context) { deleteMarks(c, a) })

	// r.POST("/mark/", func(c *gin.Context) { createScholarship(c, a) })
	// r.PUT("/mark/", func(c *gin.Context) { updateScholarshipbyID(c, a) })
	// r.DELETE("/mark/", func(c *gin.Context) { deleteScholarshipbyID(c, a) })
	// r.GET("/mark/", func(c *gin.Context) { getAllScholarship(c, a) })
}
