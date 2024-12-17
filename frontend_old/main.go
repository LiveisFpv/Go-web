package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Table struct {
	Table_name []map[string]interface{} `json:"table_name"`
	Name       string                   `json:"name"`
	Columns    []string                 `json:"columns"`
	Rows       []map[string]interface{} `json:"rows"`
}

var tableData = Table{
	Table_name: []map[string]interface{}{
		{"link": "#", "name": "table"},
		{"link": "#", "name": "table"},
		{"link": "#", "name": "table"},
		{"link": "#", "name": "table"},
		{"link": "#", "name": "table"},
	},
	Name:    "Dynamic Table",
	Columns: []string{"ID", "Name", "Priority", "Value"},
	Rows: []map[string]interface{}{
		{"ID": 1, "Name": "Project A", "Priority": "High", "Value": 12345},
		{"ID": 2, "Name": "Project B", "Priority": "Low", "Value": 54321},
		{"ID": 3, "Name": "Project C", "Priority": "Medium", "Value": 67890},
	},
}

func main() {
	router := gin.Default()

	// Статика
	router.Static("/static", "./static")

	// Роуты
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "table.html", tableData)
	})
	router.POST("/update", updateHandler)

	// Шаблоны
	router.LoadHTMLGlob("templates/*")

	// Запуск сервера
	router.Run(":8080")
}

func updateHandler(c *gin.Context) {
	var updatedRow map[string]interface{}
	if err := c.ShouldBindJSON(&updatedRow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(updatedRow["ID"].(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Обновление данных
	for i, row := range tableData.Rows {

		if row["ID"] == id {
			tableData.Rows[i] = updatedRow
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record updated successfully", "data": updatedRow})
}
