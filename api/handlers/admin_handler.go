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

func GetStats(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var stats models.DashboardStat
		var userCount int64
		db.Model(&models.User{}).Count(&userCount)
		var adCount int64
		db.Model(&model.Announcements{}).Where("status = ?", "available").Count(&adCount)
		
		stats.TotalUsers = int(userCount)
		stats.ActiveAds = int(adCount)
		stats.TotalUpcycled = 15000
		stats.WasteAvoided = 2.5
		Stats.TotalImpact = 5000
		
		c.JSON(http.StatusOK, stats)
	}
}