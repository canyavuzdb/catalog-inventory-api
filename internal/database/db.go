package database

import (
	"fmt"
	"log"

	"github.com/canyavuzdb/catalog-inventory-api/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=catalog_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	fmt.Println("Database connected")

	err = db.AutoMigrate(&domain.Product{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	seedProducts(db)

	DB = db
}

func seedProducts(db *gorm.DB) {
	var count int64
	db.Model(&domain.Product{}).Count(&count)

	if count > 0 {
		return
	}

	products := []domain.Product{
		{Name: "Laptop", Price: 1500},
		{Name: "Keyboard", Price: 100},
		{Name: "Mouse", Price: 50},
		{Name: "Monitor", Price: 300},
	}

	if err := db.Create(&products).Error; err != nil {
		log.Fatal("failed to seed products:", err)
	}

	fmt.Println("Seed data inserted")
}
