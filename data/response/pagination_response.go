package response

type PaginationResponse[T interface{}] struct {
	Items       []T  `json:"items"`
	PerPage     int  `json:"per_page"`
	Total       int  `json:"total"`
	FirstPage   int  `json:"first_page"`
	LastPage    int  `json:"last_page"`
	HasNext     bool `json:"has_next"`
	HasPrev     bool `json:"has_prev"`
	Visible     int  `json:"visible"`
	CurrentPage int  `json:"current_page"`
}
