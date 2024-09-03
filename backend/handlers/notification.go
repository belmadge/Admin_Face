package handlers

import (
	"net/http"

	db "github.com/belmadge/Admin_Face/infra/repository"
	"github.com/belmadge/Admin_Face/models"
	"github.com/gin-gonic/gin"
)

// GetNotifications returns all notifications
func GetNotifications(c *gin.Context) {
	var notifications []models.Notification
	db.DB.Find(&notifications)
	c.JSON(http.StatusOK, notifications)
}

// CreateNotification creates a new notification
func CreateNotification(c *gin.Context) {
	var notification models.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&notification)
	c.JSON(http.StatusOK, gin.H{"message": "Notification created successfully"})
}

// DeleteNotification deletes a notification
func DeleteNotification(c *gin.Context) {
	id := c.Param("id")
	var notification models.Notification
	if err := db.DB.First(&notification, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	db.DB.Delete(&notification)
	c.JSON(http.StatusOK, gin.H{"message": "Notification deleted successfully"})
}
