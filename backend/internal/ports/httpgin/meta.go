package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getStudentMeta(c *gin.Context) {
	meta := gin.H{
		"data": []map[string]interface{}{
			{"name": "id_num_student", "type": "number", "required": true},
			{"name": "name_group", "type": "text", "required": true},
			{"name": "email_student", "type": "email", "required": true},
			{"name": "second_name_student", "type": "text", "required": true},
			{"name": "first_name_student", "type": "text", "required": true},
			{"name": "surname_student", "type": "text", "required": false},
		},
	}
	c.JSON(http.StatusOK, meta)
}

func getGroupMeta(c *gin.Context) {
	meta := gin.H{
		"data": []map[string]interface{}{
			{"name": "name_group", "type": "text", "required": true},
			{"name": "studies_direction_group", "type": "text", "required": true},
			{"name": "studies_profile_group", "type": "text", "required": true},
			{"name": "start_date_group", "type": "text", "required": true},
			{"name": "studies_period_group", "type": "number", "required": true},
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
