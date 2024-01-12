package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	XRequestIDKey = "X-Request-ID"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(XRequestIDKey)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set(XRequestIDKey, requestID)
		c.Writer.Header().Set(XRequestIDKey, requestID)
		c.Next()
	}
}
