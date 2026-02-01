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
	menucreate "ayam_bangkok/source/features/menu/menu_create"
	menudelete "ayam_bangkok/source/features/menu/menu_delete"
	menuget "ayam_bangkok/source/features/menu/menu_get"
	menugetbyid "ayam_bangkok/source/features/menu/menu_get_by_id"
	menuupdate "ayam_bangkok/source/features/menu/menu_update"
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

	// menu routes
	menuRoute := routeGroup.Group("/menu").Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		menuRoute.GET("", menuget.NewHandler(r.db))
		menuRoute.GET("/detail/:menu_id", menugetbyid.NewHandler(r.db))
		menuRoute.POST("/create", menucreate.NewHandler(r.db))
		menuRoute.PUT("/update/:menu_id", menuupdate.NewHandler(r.db))
		menuRoute.DELETE("/delete/:menu_id", menudelete.NewHandler(r.db))
	}
}