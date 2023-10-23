package response

import "todogorest/models"

type PaginationResponse struct {
	Items       []models.Todo `json:"items"`
	PerPage     int           `json:"per_page"`
	Total       int           `json:"total"`
	HasNext     bool          `json:"has_next"`
	HasPrev     bool          `json:"has_prev"`
	Next        int           `json:"next"`
	Prev        int           `json:"prev"`
	CurrentPage int           `json:"current_page"`
}
