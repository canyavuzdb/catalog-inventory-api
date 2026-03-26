package domain

import "time"

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"size:255;not null"`
	Description string  `gorm:"size:500"`
	SKU         string  `gorm:"size:100;uniqueIndex;not null"`
	Price       float64 `gorm:"not null"`
	CategoryID  uint    `gorm:"not null"`
	Category    Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
