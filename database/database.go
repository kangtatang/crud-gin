package database

import (
	"fmt"
	"log"
	"os"

	"gin-user-management/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load konfigurasi dari .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Buat DSN (Database Source Name)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Koneksi ke PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto Migrate: Membuat tabel jika belum ada
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Simpan koneksi ke variabel global
	DB = db
	fmt.Println("Database connected & migrated successfully!")
}

// package database

// import (
// 	"fmt"
// 	"log"

// 	"gin-user-management/models"

// 	"gorm.io/driver/postgres"
// 	"github.com/joho/godotenv"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func InitDB() {
// 	// Sesuaikan dengan konfigurasi PostgreSQL Anda
// 	dsn := "host=localhost user=postgres password=yourpassword dbname=local_coba port=5432 sslmode=disable TimeZone=Asia/Jakarta"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}

// 	// AutoMigrate untuk membuat tabel jika belum ada
// 	err = db.AutoMigrate(&models.User{})
// 	if err != nil {
// 		log.Fatalf("Failed to migrate database: %v", err)
// 	}

// 	fmt.Println("Database connected successfully!")
// 	DB = db
// }
