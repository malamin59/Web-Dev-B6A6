package routes

import (
	"spotsync-api/handler"
	"spotsync-api/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterReservationRoutes(e *echo.Echo, reservationHandler *handler.ReservationHandler) {

	e.POST(
		"/api/v1/reservations",
		reservationHandler.Create,
		middleware.JWTMiddleware,
	)
	e.GET(
		"/api/v1/reservations/me",
		reservationHandler.GetMyReservations,
		middleware.JWTMiddleware,
	)
	e.PATCH(
		"/api/v1/reservations/:id/cancel",
		reservationHandler.Cancel,
		middleware.JWTMiddleware,
	)
	e.GET(
		"/api/v1/admin/reservations",
		reservationHandler.GetAll,
		middleware.JWTMiddleware,
		middleware.RoleMiddleware("admin"),
	)
}
