package request

import "time"

type SigninUserRequest struct {
	Username string `json:"username" validate:"string,required"`
	Password string `json:"password" validate:"string,required"`
}

type CreateUserRequest struct {
	Username    string     `json:"username" validate:"string,required,max=20"`
	Password    string     `json:"password" validate:"string,required,min=8,max=30"`
	Locale      string     `json:"locale" validate:"string,required,len=2"`
	Email       string     `json:"email" validate:"email,required"`
	DateOfBirth *time.Time `json:"date_of_birth" validate:"date,omitempty"`
	CountryCode *string    `json:"country_code" validate:"string,required_with=PhoneNumber"`
	PhoneNumber *int       `json:"phone_number" validate:"string,required_with=CountryCode"`
}

type UpdateUserRequest struct {
	Username    string     `json:"username" validate:"string,required,max=20"`
	Locale      string     `json:"locale" validate:"string,required,len=2"`
	Email       string     `json:"email" validate:"email,required"`
	DateOfBirth *time.Time `json:"date_of_birth" validate:"date,omitempty"`
	CountryCode *string    `json:"country_code" validate:"string,required_with=PhoneNumber"`
	PhoneNumber *int       `json:"phone_number" validate:"string,required_with=CountryCode"`
}
