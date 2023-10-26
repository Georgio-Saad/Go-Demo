package models

import (
	"time"
)

type User struct {
	ID             uint       `gorm:"primarykey" json:"id" `
	CreatedAt      time.Time  `json:"-"`
	UpdatedAt      time.Time  `json:"-"`
	Username       string     `json:"username" gorm:"notnull;unique"`
	ProfilePicture *string    `json:"profile_picture"`
	Password       string     `json:"-"`
	Email          string     `json:"email" gorm:"notnull;unique"`
	DateOfBirth    *time.Time `json:"date_of_birth"`
	CountryCode    *string    `json:"country_code"`
	PhoneNumber    *int       `json:"phone_number" gorm:"unique"`
	Verified       bool       `json:"verified"`
	Role           string     `json:"role"`
	ProductID      uint       `json:"-"`
	Product        Product    `json:"product"`
}
