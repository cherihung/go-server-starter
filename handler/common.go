package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Default handler
func Default(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Success")
}

// HealthCheck handler
func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
