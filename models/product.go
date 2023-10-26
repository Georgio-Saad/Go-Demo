package models

import "time"

type Product struct {
	ID        uint      `gorm:"primarykey" json:"id" `
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Slug      string    `json:"slug" gorm:"notnull;unique"`
	Product   string    `json:"product" gorm:"notnull;unique"`
}
