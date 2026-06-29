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

	e.Logger.Fatal(e.Start(":8080"))
}
