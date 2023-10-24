package models

import (
	"time"

	"gorm.io/gorm"
)

type VerificationCode struct {
	ID               uint           `gorm:"primarykey" json:"id" `
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	VerificationCode string         `json:"verification_code"`
	AlreadyUsed      bool           `json:"already_used"`
	UserID           int
	User             User
}
