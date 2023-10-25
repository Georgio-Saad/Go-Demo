package models

import (
	"time"
)

type User struct {
	ID          uint       `gorm:"primarykey" json:"id" `
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	Username    string     `json:"username"`
	Password    string     `json:"-"`
	Locale      string     `json:"locale"`
	Email       string     `json:"email"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	CountryCode *string    `json:"country_code"`
	PhoneNumber *int       `json:"phone_number"`
	Verified    bool       `json:"verified"`
}
