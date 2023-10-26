package request

type PaginationRequest struct {
	Page string `json:"page"`
	Size string `json:"size"`
}
