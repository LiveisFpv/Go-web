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
func getTables(c *gin.Context) {
	tables := []string{
		"student",
		"group",
	}
	c.JSON(http.StatusOK, gin.H{"data": tables})
}
