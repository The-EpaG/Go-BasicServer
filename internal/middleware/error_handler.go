package middleware

import (
	"github.com/The-EpaG/Go-BasicServer/pkg/exceptions"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Execute next middleware

		if len(c.Errors) > 0 {
			// Get the last error
			err := c.Errors.Last()
			apiErr, ok := err.Err.(exceptions.APIError)

			if !ok {
				apiErr = exceptions.ErrInternal
			}

			c.JSON(apiErr.HttpCode, apiErr.ErrorCode)
		}
	}
}
