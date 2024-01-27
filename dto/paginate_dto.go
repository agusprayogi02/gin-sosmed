package dto

type Paginate struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

type PaginateRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
