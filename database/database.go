package database

import (
	"fmt"
	"log"

	"gin-user-management/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Sesuaikan dengan konfigurasi PostgreSQL Anda
	dsn := "host=localhost user=postgres password=yourpassword dbname=local_coba port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// AutoMigrate untuk membuat tabel jika belum ada
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Database connected successfully!")
	DB = db
}
