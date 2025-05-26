package handlers

import (
	"backend/internal/ports/httpgin/presenters"
)

// Import all types and functions from presenters
type (
	StudentRequest            = presenters.StudentRequest
	StudentDeleteRequest      = presenters.StudentDeleteRequest
	GroupRequest              = presenters.GroupRequest
	GroupDeleteRequest        = presenters.GroupDeleteRequest
	MarkRequest               = presenters.MarkRequest
	MarksDeleteRequest        = presenters.MarksDeleteRequest
	AuthRequest               = presenters.AuthRequest
	RegisterRequest           = presenters.RegisterRequest
	ScholarshipRequest        = presenters.ScholarshipRequest
	ScholarshipsDeleteRequest = presenters.ScholarshipsDeleteRequest
	SemesterRequest           = presenters.SemesterRequest
	SemesterDeleteRequest     = presenters.SemesterDeleteRequest
	AchievementRequest        = presenters.AchievementRequest
	AchievementsDeleteRequest = presenters.AchievementsDeleteRequest
)

// Import all response functions
var (
	ErrorResponse                 = presenters.ErrorResponse
	SuccessResponse               = presenters.SuccessResponse
	AllSuccessResponse            = presenters.AllSuccessResponse
	StudentSuccessResponse        = presenters.StudentSuccessResponse
	AllStudentSuccessResponse     = presenters.AllStudentSuccessResponse
	GroupSuccessResponse          = presenters.GroupSuccessResponse
	AllGroupSuccessResponse       = presenters.AllGroupSuccessResponse
	MarkSuccessResponse           = presenters.MarkSuccessResponse
	AllMarkSuccessResponse        = presenters.AllMarkSuccessResponse
	LoginSuccessResponse          = presenters.LoginSuccessResponse
	ScholarshipSuccessResponse    = presenters.ScholarshipSuccessResponse
	AllScholarshipSuccessResponse = presenters.AllScholarshipSuccessResponse
	SemesterSuccessResponse       = presenters.SemesterSuccessResponse
	AllSemesterSuccessResponse    = presenters.AllSemesterSuccessResponse
	AchievementSuccessResponse    = presenters.AchievementSuccessResponse
	AllAchievementSuccessResponse = presenters.AllAchievementSuccessResponse
)
