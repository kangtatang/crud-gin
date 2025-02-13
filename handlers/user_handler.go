package handlers

import (
	"net/http"

	"gin-user-management/models"

	"github.com/gin-gonic/gin"

	// "go.elastic.co/apm/v2/model"
	"gorm.io/gorm"
)

// UserHandler struct
type UserHandler struct {
	DB *gorm.DB
}

// NewUserHandler buat instance handler
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

// GetUsers mendapatkan semua user
func (h *UserHandler) GetUsers(c *gin.Context) {
	var users []models.User
	h.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GetUserByID mendapatkan user berdasarkan ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user []models.User

	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser menambahkan user baru
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user []models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// UpdateUser mengupdate data user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user []models.User

	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUser menghapus user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user []models.User

	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	h.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
