package domain

import "time"

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
