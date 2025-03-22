package middleware

import (
	"github.com/The-EpaG/Go-BasicServer/pkg/logging"

	"github.com/gin-gonic/gin"
)

// RequestIDMiddleware assigns a unique ID to the request
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.Request.Header.Get(string(logging.GetTraceKey()))

		// Generates and inserts the unique ID into the request context; if the traceID is empty, it creates one
		ctx := logging.InjectRequestID(c.Request.Context(), traceId)

		print(ctx.Value(logging.GetTraceKey()).(string))

		c.Request = c.Request.WithContext(ctx)
		print(c.Request.Context().Value(logging.GetTraceKey()).(string))

		// Proceed with the request
		c.Next()
	}
}
