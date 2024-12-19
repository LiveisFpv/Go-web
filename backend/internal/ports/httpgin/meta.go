package httpgin

import (
	"backend/internal/app"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

const countrows = 10

func getStudentMeta(c *gin.Context, a *app.App) {
	count, err := a.GetCountRows(c, "student")
	if err != nil {
		meta := gin.H{
			"data":  nil,
			"error": err,
		}
		c.JSON(http.StatusForbidden, meta)
		return
	}
	meta := gin.H{
		"data": []map[string]interface{}{
			{"name": "id_num_student", "type": "number", "required": true, "unique": true},
			{"name": "name_group", "type": "text", "required": true, "unique": false},
			{"name": "email_student", "type": "email", "required": true, "unique": false},
			{"name": "second_name_student", "type": "text", "required": true, "unique": false},
			{"name": "first_name_student", "type": "text", "required": true, "unique": false},
			{"name": "surname_student", "type": "text", "required": false, "unique": false},
		},
		"pages": math.Ceil(float64(count) / countrows),
		"error": err,
	}
	c.JSON(http.StatusOK, meta)
}

func getGroupMeta(c *gin.Context, a *app.App) {
	count, err := a.GetCountRows(c, "group")
	if err != nil {
		meta := gin.H{
			"data":  nil,
			"error": err,
		}
		c.JSON(http.StatusForbidden, meta)
		return
	}
	meta := gin.H{
		"data": []map[string]interface{}{
			{"name": "name_group", "type": "text", "required": true, "unique": true},
			{"name": "studies_direction_group", "type": "text", "required": true, "unique": false},
			{"name": "studies_profile_group", "type": "text", "required": true, "unique": false},
			{"name": "start_date_group", "type": "text", "required": true, "unique": false},
			{"name": "studies_period_group", "type": "number", "required": true, "unique": false},
		},
		"pages": math.Ceil(float64(count) / countrows),
		"error": err,
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
