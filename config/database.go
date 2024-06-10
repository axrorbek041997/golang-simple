package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"simple/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.User{})

	DB = database
}
