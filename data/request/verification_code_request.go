package request

type VerificationCodeRequest struct {
	AlreadyUsed      bool   `json:"already_used" validate:"boolean"`
	VerificationCode string `json:"verification_code" validate:"required,numeric,len=6"`
	UserID           int    `json:"user_id" validate:"required"`
}
