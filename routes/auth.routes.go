package routes

import (
	"spotsync-api/handler"

	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(e *echo.Echo, authHandler *handler.AuthHandler) {
	e.POST("/api/v1/auth/register", authHandler.Register)
}
