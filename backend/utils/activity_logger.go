package utils

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ActivityLogger middleware logs user activities
func ActivityLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next() // Process the request

		latency := time.Since(startTime)
		logMessage := fmt.Sprintf("User: %v | Method: %s | Path: %s | Latency: %v",
			c.GetString("user_id"),
			c.Request.Method,
			c.Request.URL.Path,
			latency)

		// Here you can save logMessage to a file or database
		fmt.Println(logMessage)
	}
}
