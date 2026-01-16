package services

import (
	"ayam_bangkok/source/common/health"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HealthCheck(r *gin.Engine, db *gorm.DB) {
	healthHandler := health.CheckHealth(db)
	r.GET("api/health", healthHandler.Check)
}