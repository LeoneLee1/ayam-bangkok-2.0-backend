package services

import (
	absenceclockinandout "ayam_bangkok/source/features/absence/absence_clock_in_and_out"
	absenceexportexcel "ayam_bangkok/source/features/absence/absence_export_excel"
	absenceget "ayam_bangkok/source/features/absence/absence_get"
	absencegetbyfilter "ayam_bangkok/source/features/absence/absence_get_by_filter"
	"ayam_bangkok/source/features/auth/login"
	"ayam_bangkok/source/features/auth/profile"
	"ayam_bangkok/source/features/auth/register"
	updatepassword "ayam_bangkok/source/features/auth/update_password"
	updateprofile "ayam_bangkok/source/features/auth/update_profile"
	"ayam_bangkok/source/services/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routers struct {
	db *gorm.DB
}

func NewRouters(db *gorm.DB) *Routers {
	return &Routers{
		db: db,
	}
}

func (r *Routers) MountRouters(routeGroup *gin.RouterGroup) {

	// auth routes
	authRoute := routeGroup.Group("")
	{
		authRoute.POST("/register", register.NewHandler(r.db))
		authRoute.POST("/login", login.NewHandler(r.db))
		
		// use auth middleware
		authRoute.Use(middleware.AuthMiddleware()).GET("/profile", profile.NewHandler(r.db))
		authRoute.Use(middleware.AuthMiddleware()).PUT("/profile/update", updateprofile.NewHandler(r.db))
		authRoute.Use(middleware.AuthMiddleware()).PATCH("/profile/update-password", updatepassword.NewHandler(r.db))
	}

	// absence routes
	absenceRoute := routeGroup.Group("/absence").Use(middleware.AuthMiddleware())
	absenceRouteAdmin := absenceRoute.Use(middleware.AdminMiddleware())
	{
		absenceRoute.POST("", absenceclockinandout.NewHandler(r.db))
		absenceRoute.GET("/detail", absenceget.NewHandler(r.db))
		absenceRouteAdmin.GET("/filter", absencegetbyfilter.NewHandler(r.db))
		absenceRouteAdmin.GET("/export", absenceexportexcel.NewHandler(r.db))
	}
}