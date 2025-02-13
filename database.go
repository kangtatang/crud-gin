package main // Atau sesuaikan jika menggunakan package lain

import (
	"gin-user-management/models" // ✅ Import package models
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	var err error
	db, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// ✅ Panggil models.User dengan prefix package
	db.AutoMigrate(&models.User{})

	return db
}
