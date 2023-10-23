package response

type PaginationResponse[T any] struct {
	Items       []T  `json:"items"`
	PerPage     int  `json:"per_page"`
	Total       int  `json:"total"`
	TotalPages  int  `json:"total_pages"`
	HasNext     bool `json:"has_next"`
	HasPrev     bool `json:"has_prev"`
	Visible     int  `json:"visible"`
	CurrentPage int  `json:"current_page"`
}
