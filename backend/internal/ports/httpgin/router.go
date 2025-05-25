package httpgin

import (
	"backend/internal/app"
	"backend/internal/ports/httpgin/handlers"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a *app.App) {
	//Информация передоваемя на фронт для динамического поддержания работы
	r.GET("/student/metadata", func(c *gin.Context) { handlers.GetStudentMeta(c, a) })
	r.GET("/group/metadata", func(c *gin.Context) { handlers.GetGroupMeta(c, a) })
	r.GET("/mark/metadata", func(c *gin.Context) { handlers.GetMarkMeta(c, a) })

	//Запросы на получение и редактирование таблиц
	r.GET("/student/:student_id", func(c *gin.Context) { handlers.GetStudentbyID(c, a) })
	r.POST("/student/", func(c *gin.Context) { handlers.CreateStudent(c, a) })
	r.PUT("/student/", func(c *gin.Context) { handlers.UpdateStudentbyID(c, a) })
	r.DELETE("/student/", func(c *gin.Context) { handlers.DeleteStudentbyID(c, a) })
	r.GET("/student/", func(c *gin.Context) { handlers.GetAllStudent(c, a) })
	r.DELETE("/student/ids", func(c *gin.Context) { handlers.DeleteStudents(c, a) })

	// r.GET("/group/:group_name", func(c *gin.Context) { getGroupbyID(c, a) })
	r.POST("/group/", func(c *gin.Context) { handlers.CreateGroup(c, a) })
	r.PUT("/group/", func(c *gin.Context) { handlers.UpdateGroupbyName(c, a) })
	r.DELETE("/group/", func(c *gin.Context) { handlers.DeleteGroupbyName(c, a) })
	r.GET("/group/", func(c *gin.Context) { handlers.GetAllGroup(c, a) })
	r.DELETE("/group/ids", func(c *gin.Context) { handlers.DeleteGroups(c, a) })

	r.POST("/mark/", func(c *gin.Context) { handlers.CreateMark(c, a) })
	r.PUT("/mark/", func(c *gin.Context) { handlers.UpdateMarkbyID(c, a) })
	r.DELETE("/mark/", func(c *gin.Context) { handlers.DeleteMarkbyID(c, a) })
	r.GET("/mark/", func(c *gin.Context) { handlers.GetAllMark(c, a) })
	r.DELETE("/mark/ids", func(c *gin.Context) { handlers.DeleteMarks(c, a) })

	// Semester operations
	r.GET("/semester/", func(c *gin.Context) { handlers.GetAllSemester(c, a) })
	r.POST("/semester/", func(c *gin.Context) { handlers.CreateSemester(c, a) })
	r.PUT("/semester/", func(c *gin.Context) { handlers.UpdateSemesterByName(c, a) })
	r.DELETE("/semester/", func(c *gin.Context) { handlers.DeleteSemesterByName(c, a) })
	r.DELETE("/semester/ids", func(c *gin.Context) { handlers.DeleteSemesters(c, a) })
}

func OpenRouter(r *gin.RouterGroup, a *app.App) {
	//Аутентификация и регистрация
	r.POST("/auth", func(c *gin.Context) { handlers.Auth(c, a) })
	r.POST("/register", func(c *gin.Context) { handlers.Register(c, a) })
	r.GET("/tables", func(c *gin.Context) { handlers.GetTables(c) })
}
