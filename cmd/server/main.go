package main

import (
	"net/http"

	"spotsync-api/config"

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

	e.Logger.Fatal(e.Start(":8080"))
}