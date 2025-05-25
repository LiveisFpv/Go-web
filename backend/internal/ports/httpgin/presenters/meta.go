package presenters

type Pagination struct {
	Total     int `json:"total"`
	Page      int `json:"page"`
	Page_size int `json:"page_size"`
}
