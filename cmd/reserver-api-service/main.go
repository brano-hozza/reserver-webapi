package main

import (
	"log"
	"os"
	"strings"

	"github.com/brano-hozza/reserver-webapi/api"
	"github.com/brano-hozza/reserver-webapi/internal/reserver"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Server started")
	port := os.Getenv("RESERVER_API_PORT")
	if port == "" {
		port = "8080"
	}
	environment := os.Getenv("RESERVER_API_ENVIRONMENT")
	if !strings.EqualFold(environment, "production") { // case insensitive comparison
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())
	// request routings
	reserver.AddRoutes(engine)
	engine.GET("/openapi", api.HandleOpenApi)
	engine.Run(":" + port)
}
