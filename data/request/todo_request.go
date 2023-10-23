package request

type TodoRequest struct {
	Item      string `json:"item" validate:"required,max=20"`
	Completed bool   `json:"completed" validate:"boolean"`
}
