package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")

		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)

		log.Printf("[RequestID: %s] Started %s %s", requestID, c.Request.Method, c.Request.URL.Path)

		c.Next()

		status := c.Writer.Status()

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				log.Printf("[RequestID: %s] Error: %v", requestID, err.Err)
			}
		}

		log.Printf("[RequestID: %s] Completed with status %d", requestID, status)
	}
}
