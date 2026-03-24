package domain

type Product struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `gorm:"size:255;not null"`
	Price float64 `gorm:"not null"`
}
