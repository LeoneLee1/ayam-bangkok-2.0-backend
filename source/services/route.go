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
	bookingroomcreate "ayam_bangkok/source/features/booking_room/booking_room_create"
	bookingroomlist "ayam_bangkok/source/features/booking_room/booking_room_list"
	bookingroomstatus "ayam_bangkok/source/features/booking_room/booking_room_status"
	bookingroomupdate "ayam_bangkok/source/features/booking_room/booking_room_update"
	menucreate "ayam_bangkok/source/features/menu/menu_create"
	menudelete "ayam_bangkok/source/features/menu/menu_delete"
	menuget "ayam_bangkok/source/features/menu/menu_get"
	menugetbyid "ayam_bangkok/source/features/menu/menu_get_by_id"
	menuupdate "ayam_bangkok/source/features/menu/menu_update"
	ordergetmenu "ayam_bangkok/source/features/order/order_get_menu"
	ordermenu "ayam_bangkok/source/features/order/order_menu"
	ordermenucancel "ayam_bangkok/source/features/order/order_menu_cancel"
	ordermenustatus "ayam_bangkok/source/features/order/order_menu_status"
	roomcreate "ayam_bangkok/source/features/room/room_create"
	roomdelete "ayam_bangkok/source/features/room/room_delete"
	roomget "ayam_bangkok/source/features/room/room_get"
	roomgetbyid "ayam_bangkok/source/features/room/room_get_by_id"
	roomupdate "ayam_bangkok/source/features/room/room_update"
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

	// order menu routes
	orderMenuRoute := routeGroup.Group("/order-menu")
	orderMenuRoute.Use(middleware.AuthMiddleware())
	{
		orderMenuRoute.GET("", ordergetmenu.NewHandler(r.db))
		orderMenuRoute.GET("/status/:menu_id", ordermenustatus.NewHandler(r.db))
		orderMenuRoute.POST("/create/:menu_id", ordermenu.NewHandler(r.db))
		orderMenuRoute.DELETE("/cancel/:menu_id", ordermenucancel.NewHandler(r.db))
	}

	// room routes
	roomRoute := routeGroup.Group("/room")
	roomRoute.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		roomRoute.POST("/create", roomcreate.NewHandler(r.db))
		roomRoute.GET("", roomget.NewHandler(r.db))
		roomRoute.GET("/detail/:room_id", roomgetbyid.NewHandler(r.db))
		roomRoute.PUT("/update/:room_id", roomupdate.NewHandler(r.db))
		roomRoute.DELETE("/delete/:room_id", roomdelete.NewHandler(r.db))
	}

	// booking room routes
	bookingRoomRoute := routeGroup.Group("/booking-rooms")
	bookingRoomRoute.Use(middleware.AuthMiddleware())
	{
		bookingRoomRoute.POST("/create/:room_id", bookingroomcreate.NewHandler(r.db))
		bookingRoomRoute.GET("/status/:room_id", bookingroomstatus.NewHandler(r.db))
		bookingRoomRoute.GET("/list/:room_id", bookingroomlist.NewHandler(r.db))
		bookingRoomRoute.PUT("update/:booking_room_id", bookingroomupdate.NewHandler(r.db))
	}
}