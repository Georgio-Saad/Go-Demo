package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primarykey" json:"id" `
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Username    string         `json:"username"`
	Password    string         `json:"-"`
	Locale      string         `json:"locale"`
	Email       string         `json:"email"`
	DateOfBirth *time.Time     `json:"date_of_birth"`
	CountryCode *string        `json:"country_code"`
	PhoneNumber *int           `json:"phone_number"`
	Verified    bool           `json:"verified"`
}
