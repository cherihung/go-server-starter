package middleware

import (
	"fmt"
	"time"

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

// CustomLogger - custom logger
func CustomLogger(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
