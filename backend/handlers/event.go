package handlers

import (
	"net/http"

	db "github.com/belmadge/Admin_Face/infra/repository"
	"github.com/belmadge/Admin_Face/models"
	"github.com/gin-gonic/gin"
)

// GetEvents returns all events
func GetEvents(c *gin.Context) {
	var events []models.Event
	db.DB.Find(&events)
	c.JSON(http.StatusOK, events)
}

// CreateEvent creates a new event
func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&event)
	c.JSON(http.StatusOK, gin.H{"message": "Event created successfully"})
}

// UpdateEvent updates an event
func UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	var event models.Event
	if err := db.DB.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&event)
	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

// DeleteEvent deletes an event
func DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	var event models.Event
	if err := db.DB.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	db.DB.Delete(&event)
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
