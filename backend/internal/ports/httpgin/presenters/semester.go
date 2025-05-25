package presenters

import (
	"backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type semesterResponse struct {
	Name_semester       string `json:"name_semester"`
	Date_start_semester string `json:"date_start_semester"`
	Date_end_semester   string `json:"date_end_semester"`
}

type SemesterRequest struct {
	Name_semester       string `json:"name_semester"`
	Date_start_semester string `json:"date_start_semester"`
	Date_end_semester   string `json:"date_end_semester"`
}

type SemesterDeleteRequest struct {
	Names []string `json:"names"`
}

// Преобразование структур
func mapSemesterToResponse(semester *domain.Semester) semesterResponse {
	return semesterResponse{
		Name_semester:       semester.Name_semester,
		Date_start_semester: semester.Date_start_semester,
		Date_end_semester:   semester.Date_end_semester,
	}
}

func SemesterSuccessResponse(semester *domain.Semester) *gin.H {
	return SuccessResponse(mapSemesterToResponse(semester))
}

func AllSemesterSuccessResponse(semesters []*domain.Semester, countRow, count, page int) *gin.H {
	data := Paginate(semesters, countRow, page, mapSemesterToResponse)
	return AllSuccessResponse(data, Pagination{
		Total:     count,
		Page:      page,
		Page_size: countRow,
	})
}
