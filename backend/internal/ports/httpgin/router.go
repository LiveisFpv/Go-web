package httpgin

import (
	"backend/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a *app.App) {
	r.GET("/student/:student_id", func(c *gin.Context) { getStudentbyID(c, a) })
	r.POST("/student", func(c *gin.Context) { createStudent(c, a) })
	r.PUT("/student/:student_id", func(c *gin.Context) { updateStudentbyID(c, a) })
	r.DELETE("/student/:student_id", func(c *gin.Context) { deleteStudentbyID(c, a) })
	r.GET("/student", func(c *gin.Context) { getAllStudent(c, a) })
	r.GET("/tables", func(c *gin.Context) { getTables(c) })
}
