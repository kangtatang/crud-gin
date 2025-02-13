package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDB inisialisasi koneksi database
func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Auto Migrate table
	db.AutoMigrate(&User{})

	return db
}
