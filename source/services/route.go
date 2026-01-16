package services

import (
	"ayam_bangkok/source/common/health"
	"ayam_bangkok/source/features/auth/register"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HealthCheck(r *gin.Engine, db *gorm.DB) {
	healthHandler := health.CheckHealth(db)
	r.GET("api/health", healthHandler.Check)
}

func AuthRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")
	
	{
		api.POST("/register", register.NewHandler(db))
	}
}