package main

import (
	"log"
	"os"
	"strings"

	"github.com/brano-hozza/reserver-webapi/api"
	"github.com/brano-hozza/reserver-webapi/internal/db_service"
	"github.com/brano-hozza/reserver-webapi/internal/reserver"
	"github.com/gin-gonic/gin"

	"context"
	"time"

	"github.com/gin-contrib/cors"
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
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{""},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
	engine.Use(corsMiddleware)

	// setup context update  middleware
	departmentService := db_service.NewMongoService[reserver.Department](db_service.MongoServiceConfig{}, "department")
	defer departmentService.Disconnect(context.Background())
	doctorService := db_service.NewMongoService[reserver.Doctor](db_service.MongoServiceConfig{}, "doctor")
	defer doctorService.Disconnect(context.Background())
	roomService := db_service.NewMongoService[reserver.Room](db_service.MongoServiceConfig{}, "room")
	defer roomService.Disconnect(context.Background())
	reservationService := db_service.NewMongoService[reserver.RoomReservation](db_service.MongoServiceConfig{}, "reservation")
	defer reservationService.Disconnect(context.Background())
	engine.Use(func(ctx *gin.Context) {
		ctx.Set("department_service", departmentService)
		ctx.Set("doctor_service", doctorService)
		ctx.Set("room_service", roomService)
		ctx.Set("reservation_service", reservationService)
		ctx.Next()
	})

	// request routings
	reserver.AddRoutes(engine)
	engine.GET("/openapi", api.HandleOpenApi)
	engine.Run(":" + port)
}
