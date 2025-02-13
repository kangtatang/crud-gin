package helpers

import "github.com/gin-gonic/gin"

// Standard API response format
type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // Data bisa null jika tidak ada
}

// Function to create a response
func NewResponse(status int, message string, data interface{}) APIResponse {
	return APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

// Function to send JSON response
func JSONResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, NewResponse(status, message, data))
}
