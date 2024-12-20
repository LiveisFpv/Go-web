package httpgin

import (
	"backend/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getStudentMeta(c *gin.Context, a *app.App) {
	meta := gin.H{
		"data": []map[string]interface{}{
			{"name": "id_num_student", "type": "number", "required": true, "unique": true, "id": false},
			{"name": "name_group", "type": "text", "required": true, "unique": false, "id": false},
			{"name": "email_student", "type": "email", "required": true, "unique": false, "id": false},
			{"name": "second_name_student", "type": "text", "required": true, "unique": false, "id": false},
			{"name": "first_name_student", "type": "text", "required": true, "unique": false, "id": false},
			{"name": "surname_student", "type": "text", "required": false, "unique": false, "id": false},
		},
		"error": nil,
	}
	c.JSON(http.StatusOK, meta)
}

func getGroupMeta(c *gin.Context, a *app.App) {
	meta := gin.H{
		"data": []map[string]interface{}{
			{"name": "name_group", "type": "text", "required": true, "unique": true, "id": false},
			{"name": "studies_direction_group", "type": "text", "required": true, "unique": false, "id": false},
			{"name": "studies_profile_group", "type": "text", "required": true, "unique": false, "id": false},
			{"name": "start_date_group", "type": "date", "required": true, "unique": false, "id": false},
			{"name": "studies_period_group", "type": "number", "required": true, "unique": false, "id": false},
		},
		"error": nil,
	}
	c.JSON(http.StatusOK, meta)
}

func getMarkMeta(c *gin.Context, a *app.App) {
	meta := gin.H{
		"data": []map[string]interface{}{
			{"name": "id_mark", "type": "number", "required": true, "unique": true, "id": true},
			{"name": "id_num_student", "type": "number", "required": true, "unique": false, "id": false},
			{"name": "name_semester", "type": "text", "required": true, "unique": false, "id": false},
			{"name": "lesson_name_mark", "type": "text", "required": true, "unique": false, "id": false},
			{"name": "score_mark", "type": "number", "required": true, "unique": false, "id": false},
			{"name": "type_mark", "type": "text", "required": true, "unique": false, "id": false},
		},
		"error": nil,
	}
	c.JSON(http.StatusOK, meta)
}

func getTables(c *gin.Context) {
	tables := []string{
		"student",
		"group",
		"mark",
		"semester",
		"scholarship",
		"budget",
		"achievement",
		"category",
	}
	c.JSON(http.StatusOK, gin.H{"data": tables})
}
