package middleware

import (
	"bytes"
	"io"

	"github.com/The-EpaG/Go-BasicServer/pkg/logging"

	"github.com/gin-gonic/gin"
)

func HTTPLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		reqBody, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			logging.Error(ctx, "Error while reading request body: "+err.Error())
			return
		}
		// Set again the request body
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))

		logging.Info(ctx.Request.Context(), "Incoming Request: "+ctx.Request.Method+" "+ctx.Request.URL.String())
		logging.Info(ctx.Request.Context(), "Incoming Request Payload: "+string(reqBody))

		// Req processing
		ctx.Next()
	}
}
