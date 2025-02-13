package handlers

import (
	"errors"
	"net/http"
	"time"

	"gin-user-management/database"
	"gin-user-management/helpers"
	"gin-user-management/models"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

// UserHandler struct
type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		DB: database.DB, // Pastikan ini mengambil database yang sudah diinisialisasi
	}
}

// GetUsers mendapatkan semua user
func (h *UserHandler) GetUsers(c *gin.Context) {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		helpers.JSONResponse(c, http.StatusInternalServerError, "Gagal mengambil data pengguna", nil)
		return
	}

	if len(users) == 0 {
		helpers.JSONResponse(c, http.StatusOK, "Belum ada data pengguna ditemukan", nil)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Data pengguna berhasil diambil", users)
}

// GetUserByID mendapatkan user berdasarkan ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL

	var user models.User
	if err := h.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.JSONResponse(c, http.StatusNotFound, "User not found", nil)
			return
		}
		helpers.JSONResponse(c, http.StatusInternalServerError, "Failed to retrieve user", nil)
		return
	}

	// Berikan respons sukses
	helpers.JSONResponse(c, http.StatusOK, "User retrieved successfully", user)
}

// CreateUser menambahkan user baru
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	// Binding JSON ke struct user
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.JSONResponse(c, http.StatusBadRequest, "Invalid request data", nil)
		return
	}

	// Cek apakah email sudah digunakan
	var existingUser models.User
	if err := h.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		helpers.JSONResponse(c, http.StatusConflict, "Email already exists", nil)
		return
	}

	// Simpan user ke database
	if err := h.DB.Create(&user).Error; err != nil {
		helpers.JSONResponse(c, http.StatusInternalServerError, "Failed to create user", nil)
		return
	}

	helpers.JSONResponse(c, http.StatusCreated, "User created successfully", user)
}

// UpdateUser mengupdate data user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL

	var user models.User
	if err := h.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.JSONResponse(c, http.StatusNotFound, "User not found", nil)
			return
		}
		helpers.JSONResponse(c, http.StatusInternalServerError, "Failed to retrieve user", nil)
		return
	}

	var updateData models.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		helpers.JSONResponse(c, http.StatusBadRequest, "Invalid request data", nil)
		return
	}

	// Update data user yang ditemukan
	user.Name = updateData.Name
	user.Email = updateData.Email
	user.UpdatedAt = time.Now().Unix() // Update timestamp

	if err := h.DB.Save(&user).Error; err != nil {
		helpers.JSONResponse(c, http.StatusInternalServerError, "Failed to update user", nil)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "User updated successfully", user)
}

// DeleteUser menghapus user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL

	var user models.User
	if err := h.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.JSONResponse(c, http.StatusNotFound, "User not found", nil)
			return
		}
		helpers.JSONResponse(c, http.StatusInternalServerError, "Failed to retrieve user", nil)
		return
	}

	// Hapus user
	if err := h.DB.Delete(&user).Error; err != nil {
		helpers.JSONResponse(c, http.StatusInternalServerError, "Failed to delete user", nil)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "User deleted successfully", nil)
}
