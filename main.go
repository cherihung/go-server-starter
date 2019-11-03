package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cherihung/go-server-starter/handler"
	"github.com/cherihung/go-server-starter/middleware"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	var router = gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.GET("/", handler.Default)
	router.GET("/.well-known/health", handler.HealthCheck)

	return router
}

func main() {
	var serverErr error

	r := setupRouter()
	addr := fmt.Sprintf(":%d", 8080)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	serverErr = server.ListenAndServe()

	if serverErr != nil {
		log.Fatal("server start error: ", serverErr)
	}
}
