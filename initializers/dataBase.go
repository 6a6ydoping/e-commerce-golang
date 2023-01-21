package initializers

import (
	"e-commerce-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDataBase() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database...")
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Client{}, &models.Admin{}, &models.Seller{}, &models.Item{})
}
