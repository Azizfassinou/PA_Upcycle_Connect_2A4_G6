package handlers

import (
	"backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var users []models.User
        if err := db.Find(&users).Error; err != nil {
            c.JSON(500, gin.H{"error": "Failed to fetch users"})
            return
        }
        c.JSON(http.StatusOK, users)
	}
}