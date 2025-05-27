package presenters

import (
	"github.com/gin-gonic/gin"
)

// Общий интерфейс ответа
type ResponseData interface{}

// Общая функция для успешного ответа
func SuccessResponse(data ResponseData) *gin.H {
	return &gin.H{
		"data":  data,
		"error": nil,
	}
}

// Общая функция для ошибки
func ErrorResponse(err error) *gin.H {
	return &gin.H{
		"data":  nil,
		"error": err,
	}
}
func AllSuccessResponse(data ResponseData, pagination Pagination) *gin.H {
	return &gin.H{
		"data":       data,
		"pagination": pagination,
		"error":      nil,
	}
}
func Filter[R any](items []R, filter func(R) bool) []R {
	result := make([]R, 0, len(items))
	for _, item := range items {
		if filter(item) {
			result = append(result, item)
		}
	}
	return result
}

// Универсальная функция для преобразования
func Paginate[T any, R any](items []T, countRow, page int, mapper func(T) R) []R {
	result := make([]R, 0, countRow)
	for _, item := range items {
		result = append(result, mapper(item))
	}
	return result
}
