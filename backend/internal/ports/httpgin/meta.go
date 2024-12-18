package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getStudentMeta(c *gin.Context) {
	meta := gin.H{
		"data": []map[string]interface{}{
			{"name": "id_num_student", "type": "number", "required": true, "unique": true},
			{"name": "name_group", "type": "text", "required": true, "unique": false},
			{"name": "email_student", "type": "email", "required": true, "unique": false},
			{"name": "second_name_student", "type": "text", "required": true, "unique": false},
			{"name": "first_name_student", "type": "text", "required": true, "unique": false},
			{"name": "surname_student", "type": "text", "required": false, "unique": false},
		},
	}
	c.JSON(http.StatusOK, meta)
}

func getGroupMeta(c *gin.Context) {
	meta := gin.H{
		"data": []map[string]interface{}{
			{"name": "name_group", "type": "text", "required": true, "unique": true},
			{"name": "studies_direction_group", "type": "text", "required": true, "unique": false},
			{"name": "studies_profile_group", "type": "text", "required": true, "unique": false},
			{"name": "start_date_group", "type": "text", "required": true, "unique": false},
			{"name": "studies_period_group", "type": "number", "required": true, "unique": false},
		},
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
