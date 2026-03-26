package database

import (
	"fmt"
	"log"

	"github.com/canyavuzdb/catalog-inventory-api/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(runSeed bool) {
	dsn := "host=localhost user=postgres password=postgres dbname=catalog_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	fmt.Println("Database connected")

	if runSeed {
		seedCategories(db)
		seedProducts(db)
	}

	DB = db
}

func seedCategories(db *gorm.DB) {
	var count int64
	db.Model(&domain.Category{}).Count(&count)
	if count > 0 {
		return
	}

	categories := []domain.Category{
		{Name: "Electronics"},
		{Name: "Accessories"},
		{Name: "Office"},
	}

	if err := db.Create(&categories).Error; err != nil {
		log.Fatal("failed to seed categories:", err)
	}

	fmt.Println("Seed categories inserted")
}

func seedProducts(db *gorm.DB) {
	var count int64
	db.Model(&domain.Product{}).Count(&count)
	if count > 0 {
		return
	}

	var electronics domain.Category
	var accessories domain.Category

	if err := db.Where("name = ?", "Electronics").First(&electronics).Error; err != nil {
		log.Fatal("failed to find Electronics category:", err)
	}

	if err := db.Where("name = ?", "Accessories").First(&accessories).Error; err != nil {
		log.Fatal("failed to find Accessories category:", err)
	}

	products := []domain.Product{
		{
			Name:        "Laptop",
			Description: "14 inch business laptop",
			SKU:         "LAP-001",
			Price:       1500,
			CategoryID:  electronics.ID,
		},
		{
			Name:        "Keyboard",
			Description: "Mechanical keyboard",
			SKU:         "KEY-001",
			Price:       100,
			CategoryID:  accessories.ID,
		},
		{
			Name:        "Mouse",
			Description: "Wireless mouse",
			SKU:         "MOU-001",
			Price:       50,
			CategoryID:  accessories.ID,
		},
	}

	if err := db.Create(&products).Error; err != nil {
		log.Fatal("failed to seed products:", err)
	}

	fmt.Println("Seed products inserted")
}
