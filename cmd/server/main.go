package main

import (
	"net/http"

	"spotsync-api/config"
	"spotsync-api/handler"
	"spotsync-api/repository"
	"spotsync-api/routes"
	"spotsync-api/service"

	"github.com/labstack/echo/v4"
)

func main() {

	// Database Connection
	config.ConnectDatabase()

	// Echo Server
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "SpotSync API is running")
	})

	userRepo := repository.NewUserRepository(config.DB)

	authService := service.NewAuthService(userRepo)

	authHandler := handler.NewAuthHandler(authService)

	routes.RegisterAuthRoutes(e, authHandler)

	parkingRepo := repository.NewParkingZoneRepository(config.DB)
	parkingService := service.NewParkingZoneService(parkingRepo)
	parkingHandler := handler.NewParkingZoneHandler(parkingService)
	routes.RegisterParkingZoneRoutes(e, parkingHandler)

	// Reservation Module
	// =========================
	reservationRepo := repository.NewReservationRepository(config.DB)
	reservationService := service.NewReservationService(
		reservationRepo,
		parkingRepo,
	)
	reservationHandler := handler.NewReservationHandler(reservationService)
	routes.RegisterReservationRoutes(e, reservationHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
