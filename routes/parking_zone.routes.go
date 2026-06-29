package routes

import (
	"spotsync-api/handler"
	"spotsync-api/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterParkingZoneRoutes(e *echo.Echo, parkingHandler *handler.ParkingZoneHandler) {

	e.POST(
		"/api/v1/parking-zones",
		parkingHandler.Create,
		middleware.JWTMiddleware,
		middleware.RoleMiddleware("admin"),
	)
	e.GET(
		"/api/v1/parking-zones",
		parkingHandler.GetAll,
		// middleware.JWTMiddleware,
	)
	e.GET("/api/v1/parking-zones/:id", parkingHandler.GetByID)
}
