package main

import (
	"ayam_bangkok/source/common/health"
	"ayam_bangkok/source/config"
	"ayam_bangkok/source/pkg/db"
	"ayam_bangkok/source/pkg/logger"
	"ayam_bangkok/source/services"
	"ayam_bangkok/source/services/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init(true)
	logger.Info().Msg("App Started")

	cfg := config.Load()
	dbConn, err := db.Database(cfg)
	if err != nil {
		logger.Error().Msg("Database connection failed")
		return
	}
	logger.Info().Msg("Database connected")

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	// HealthCheck /health
	healthHandler := health.CheckHealth(dbConn)
	r.GET("/health", healthHandler.Check)
	
	// Static File GLOBAL
	r.Static("/storage", "./storage")

	// Mounting routers
	route_api_v1 := r.Group("api/v1")
	mounthRote := services.NewRouters(dbConn)
	mounthRote.MountRouters(route_api_v1)

	// Run Server
	logger.Info().Msg("Server running on port 8080")
	if err := r.Run(":8080"); err != nil {
		logger.Log.Fatal().Msg("Server failed to start")
	}
}