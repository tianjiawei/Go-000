package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestIDHandler sets unique request id.
// If header `X-Request-ID` is already present in the request, that is considered the
// request id. Otherwise, generates a new unique ID.
func RequestIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.Request.Header.Get("X-Request-ID")
		if rid == "" {
			rid = uuid.New().String()
			c.Request.Header.Set("X-Request-ID", rid)
		}
		c.Next()
	}
}
