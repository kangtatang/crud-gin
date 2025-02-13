package models

import "gorm.io/gorm"

// User struct untuk tabel users
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}
