package services

import (
	"ayam_bangkok/source/common/health"
	"ayam_bangkok/source/features/auth/login"
	"ayam_bangkok/source/features/auth/profile"
	"ayam_bangkok/source/features/auth/register"
	updatepassword "ayam_bangkok/source/features/auth/update_password"
	updateprofile "ayam_bangkok/source/features/auth/update_profile"
	"ayam_bangkok/source/services/middleware"

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
		api.POST("/login", login.NewHandler(db))
		api.Use(middleware.AuthMiddleware()).GET("/profile", profile.NewHandler(db))
		api.Use(middleware.AuthMiddleware()).PUT("/profile/update", updateprofile.NewHandler(db))
		api.Use(middleware.AuthMiddleware()).PATCH("/profile/update-password", updatepassword.NewHandler(db))
	}
}