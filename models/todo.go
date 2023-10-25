package models

import (
	"time"
)

type Todo struct {
	ID        uint      `gorm:"primarykey" json:"id" `
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Item      string    `json:"item"`
	Completed bool      `json:"completed"`
}
