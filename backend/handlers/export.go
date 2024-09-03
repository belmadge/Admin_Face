package handlers

import (
	"encoding/csv"
	"net/http"
	"os"
	"strconv"
	"time"

	db "github.com/belmadge/Admin_Face/infra/repository"
	"github.com/belmadge/Admin_Face/models"
	"github.com/gin-gonic/gin"
)

// ExportEvents exports event data to CSV
func ExportEvents(c *gin.Context) {
	var events []models.Event
	db.DB.Find(&events)

	fileName := "events_" + time.Now().Format("20060102") + ".csv"
	file, err := os.Create(fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create CSV file"})
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Name", "Location", "Date", "Description"})

	for _, event := range events {
		writer.Write([]string{
			strconv.Itoa(int(event.ID)),
			event.Name,
			event.Location,
			event.Date.Format("2006-01-02"),
			event.Description,
		})
	}

	c.File(fileName)
}
