package models

import (
	"time"
)

type VerificationCode struct {
	ID               uint      `gorm:"primarykey" json:"id" `
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
	VerificationCode string    `json:"verification_code"`
	AlreadyUsed      bool      `json:"already_used"`
	UserID           int       `json:"-" gorm:"notnull;unique"`
	User             User      `json:"user"`
}
