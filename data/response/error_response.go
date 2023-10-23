package response

type ErrorResponse struct {
	StatusCode int         `json:"status_code"`
	Code       string      `json:"code"`
	Data       interface{} `json:"data,omitempty"`
}

type ErrorMessage struct {
	Message string `json:"message,omitempty"`
}
