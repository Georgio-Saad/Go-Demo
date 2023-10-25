package request

import "time"

type SigninUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserRequest struct {
	Username    string     `json:"username" validate:"required,max=20"`
	Password    string     `json:"password" validate:"required,min=8,max=30"`
	Email       string     `json:"email" validate:"email,required"`
	DateOfBirth *time.Time `json:"date_of_birth" validate:"omitempty"`
	CountryCode *string    `json:"country_code" validate:"required_with=PhoneNumber"`
	PhoneNumber *int       `json:"phone_number" validate:"required_with=CountryCode"`
	Role        string     `json:"role" validate:"required"`
}

type UpdateUserRequest struct {
	Username    string     `json:"username" validate:"required,max=20"`
	Email       string     `json:"email" validate:"email,required"`
	DateOfBirth *time.Time `json:"date_of_birth" validate:"omitempty"`
	CountryCode *string    `json:"country_code" validate:"required_with=PhoneNumber"`
	PhoneNumber *int       `json:"phone_number" validate:"required_with=CountryCode"`
}

type RefreshUserRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type VerifyUserRequest struct {
	UserID           string `json:"user_id"`
	VerificationCode string `json:"verification_code"`
	Login            bool   `json:"login"`
}
