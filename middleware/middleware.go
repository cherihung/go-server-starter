package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware : setting, passing Headers
func CORSMiddleware(allowedOrigin string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length")
		// ctx.Writer.Header().Set("Access-Control-Expose-Headers", "")

		if ctx.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
