package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cherihung/go-server-starter/config"
	"github.com/cherihung/go-server-starter/handler"
	"github.com/cherihung/go-server-starter/middleware"
	"github.com/gin-gonic/gin"
	color "github.com/gookit/color"
)

var (
	router     *gin.Engine
	appConfigs *config.AppConfiguration
)

func init() {
	var err error
	appConfigs, err = config.NewAppConfiguration()
	mode := "debug"

	if err != nil {
		fmt.Println(fmt.Errorf("config error: %s", err))
		os.Exit(1)
	}

	if appConfigs.ReleaseMode {
		mode = "release"
	}

	gin.SetMode(mode)
}

func setupRouter() *gin.Engine {

	router = gin.Default()

	router.Use(middleware.CORSMiddleware())

	initializeCommonRoutes()
	initializeHeroRoutes()

	return router
}

func initializeCommonRoutes() {
	router.GET("/", handler.Default)
	router.GET("/.well-known/health", handler.HealthCheck)
}

func initializeHeroRoutes() {
	route := router.Group("/heros")
	{
		route.GET("/", handler.GetHeros)
		route.GET("/:id", handler.GetHeroByID)
	}
}

func main() {
	var serverErr error

	r := setupRouter()
	addr := fmt.Sprintf(":%d", appConfigs.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	if appConfigs.SSL {
		color.Cyan.Println("starting SSL...", addr)
		serverErr = server.ListenAndServeTLS("_certs/server.crt", "_certs/server.key")
	} else {
		color.Cyan.Println("starting...", addr)
		serverErr = server.ListenAndServe()
	}

	if serverErr != nil {
		log.Fatal("server start error: ", serverErr)
	}
}
