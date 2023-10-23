package response

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Code       string      `json:"code"`
	Data       interface{} `json:"data,omitempty"`
}
