package handlers

import (
	"backend/internal/ports/httpgin/presenters"
)

// Import all types and functions from presenters
type (
	StudentRequest       = presenters.StudentRequest
	StudentDeleteRequest = presenters.StudentDeleteRequest
	GroupRequest         = presenters.GroupRequest
	GroupDeleteRequest   = presenters.GroupDeleteRequest
	MarkRequest          = presenters.MarkRequest
	MarksDeleteRequest   = presenters.MarksDeleteRequest
	AuthRequest          = presenters.AuthRequest
	RegisterRequest      = presenters.RegisterRequest
)

// Import all response functions
var (
	ErrorResponse             = presenters.ErrorResponse
	SuccessResponse           = presenters.SuccessResponse
	AllSuccessResponse        = presenters.AllSuccessResponse
	StudentSuccessResponse    = presenters.StudentSuccessResponse
	AllStudentSuccessResponse = presenters.AllStudentSuccessResponse
	GroupSuccessResponse      = presenters.GroupSuccessResponse
	AllGroupSuccessResponse   = presenters.AllGroupSuccessResponse
	MarkSuccessResponse       = presenters.MarkSuccessResponse
	AllMarkSuccessResponse    = presenters.AllMarkSuccessResponse
	LoginSuccessResponse      = presenters.LoginSuccessResponse
)
