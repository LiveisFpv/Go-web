package httpgin

import (
	"homework8/internal/ads"
	"homework8/internal/user"

	"github.com/gin-gonic/gin"
)

type createAdRequest struct {
	Title  string `json:"title" validate:"max:100 min:1"`
	Text   string `json:"text" validate:"max:500 min:1"`
	UserID int64  `json:"user_id"`
}

type adResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	AuthorID  int64  `json:"author_id"`
	Published bool   `json:"published"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type changeAdStatusRequest struct {
	Published bool  `json:"published"`
	UserID    int64 `json:"user_id"`
}

type updateAdRequest struct {
	Title  string `json:"title" validate:"max:100 min:1"`
	Text   string `json:"text" validate:"max:500 min:1"`
	UserID int64  `json:"user_id"`
}
type userResponse struct {
	Id       int64
	Nickname string
	Email    string
}
type createUserRequest struct {
	Nickname string `validate:"max:100 min:1"`
	Email    string `validate:"max:100 min:1"`
}
type updateUserRequest struct {
	Nickname string `validate:"max:100 min:1"`
	Email    string `validate:"max:100 min:1"`
}

func UserSuccessResponse(user *user.User) *gin.H {
	return &gin.H{
		"data": userResponse{
			Id:       user.Id,
			Nickname: user.Nickname,
			Email:    user.Email,
		},
		"error": nil,
	}
}

func AdSuccessResponse(ad *ads.Ad) *gin.H {
	return &gin.H{
		"data": adResponse{
			ID:        ad.ID,
			Title:     ad.Title,
			Text:      ad.Text,
			AuthorID:  ad.AuthorID,
			Published: ad.Published,
			CreatedAt: ad.CreatedAt.String(),
			UpdatedAt: ad.UpdatedAt.String(),
		},
		"error": nil,
	}
}
func UserErrorResponse(err error) *gin.H {
	return &gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}

func AdListSuccessResponse(ad *[]ads.Ad) *gin.H {
	var res []adResponse
	res = make([]adResponse, 0)
	for _, a := range *ad {
		res = append(res, adResponse{
			ID:        a.ID,
			Title:     a.Title,
			Text:      a.Text,
			AuthorID:  a.AuthorID,
			Published: a.Published,
			CreatedAt: a.CreatedAt.String(),
			UpdatedAt: a.UpdatedAt.String(),
		})
	}
	return &gin.H{
		"data":  res,
		"error": nil,
	}
}

func AdErrorResponse(err error) *gin.H {
	return &gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}
