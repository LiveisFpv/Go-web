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

	// User operations
	r.GET("/user/:user_id", func(c *gin.Context) { handlers.GetUserByID(c, a) })
	r.GET("/user/email/:email", func(c *gin.Context) { handlers.GetUserByEmail(c, a) })
	r.POST("/user/", func(c *gin.Context) { handlers.CreateUser(c, a) })
	r.PUT("/user/", func(c *gin.Context) { handlers.UpdateUser(c, a) })
	r.DELETE("/user/", func(c *gin.Context) { handlers.DeleteUser(c, a) })
	r.GET("/user/", func(c *gin.Context) { handlers.GetAllUsers(c, a) })
	r.DELETE("/user/ids", func(c *gin.Context) { handlers.DeleteUsers(c, a) })

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

	// Scholarship operations
	r.POST("/scholarship/", func(c *gin.Context) { handlers.CreateScholarship(c, a) })
	r.PUT("/scholarship/", func(c *gin.Context) { handlers.UpdateScholarshipByID(c, a) })
	r.DELETE("/scholarship/", func(c *gin.Context) { handlers.DeleteScholarshipByID(c, a) })
	r.GET("/scholarship/", func(c *gin.Context) { handlers.GetAllScholarship(c, a) })
	r.DELETE("/scholarship/ids", func(c *gin.Context) { handlers.DeleteScholarships(c, a) })
	r.POST("/scholarship/assign", func(c *gin.Context) { handlers.AssignScholarships(c, a) })

	// Semester operations
	r.GET("/semester/", func(c *gin.Context) { handlers.GetAllSemester(c, a) })
	r.POST("/semester/", func(c *gin.Context) { handlers.CreateSemester(c, a) })
	r.PUT("/semester/", func(c *gin.Context) { handlers.UpdateSemesterByName(c, a) })
	r.DELETE("/semester/", func(c *gin.Context) { handlers.DeleteSemesterByName(c, a) })
	r.DELETE("/semester/ids", func(c *gin.Context) { handlers.DeleteSemesters(c, a) })

	// Budget operations
	r.GET("/budget/", func(c *gin.Context) { handlers.GetAllBudget(c, a) })
	r.POST("/budget/", func(c *gin.Context) { handlers.CreateBudget(c, a) })
	r.PUT("/budget/", func(c *gin.Context) { handlers.UpdateBudgetByID(c, a) })
	r.DELETE("/budget/:id", func(c *gin.Context) { handlers.DeleteBudgetByID(c, a) })
	r.DELETE("/budget/ids", func(c *gin.Context) { handlers.DeleteBudgets(c, a) })

	// Achievement Category operations
	r.GET("/achievement-category/", func(c *gin.Context) { handlers.GetAllAchievementCategory(c, a) })
	r.POST("/achievement-category/", func(c *gin.Context) { handlers.CreateAchievementCategory(c, a) })
	r.PUT("/achievement-category/", func(c *gin.Context) { handlers.UpdateAchievementCategoryByID(c, a) })
	r.DELETE("/achievement-category/", func(c *gin.Context) { handlers.DeleteAchievementCategoryByID(c, a) })
	r.DELETE("/achievement-category/ids", func(c *gin.Context) { handlers.DeleteAchievementCategories(c, a) })

	// Achievement operations
	r.POST("/achievement/", func(c *gin.Context) { handlers.CreateAchievement(c, a) })
	r.PUT("/achievement/", func(c *gin.Context) { handlers.UpdateAchievementByID(c, a) })
	r.DELETE("/achievement/", func(c *gin.Context) { handlers.DeleteAchievementByID(c, a) })
	r.GET("/achievement/", func(c *gin.Context) { handlers.GetAllAchievement(c, a) })
	r.DELETE("/achievement/ids", func(c *gin.Context) { handlers.DeleteAchievements(c, a) })
}

func OpenRouter(r *gin.RouterGroup, a *app.App) {
	//Аутентификация и регистрация
	r.POST("/auth", func(c *gin.Context) { handlers.Auth(c, a) })
	r.POST("/register", func(c *gin.Context) { handlers.Register(c, a) })
	r.GET("/tables", func(c *gin.Context) { handlers.GetTables(c) })
}
