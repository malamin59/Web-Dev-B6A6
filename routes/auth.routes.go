package routes

import (
	"spotsync-api/handler"
	"spotsync-api/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(e *echo.Echo, authHandler *handler.AuthHandler) {
	e.POST("/api/v1/auth/register", authHandler.Register)
	e.POST("/api/v1/auth/login", authHandler.Login)
	e.GET("/api/v1/profile", authHandler.Profile, middleware.JWTMiddleware)
	e.GET(
		"/api/v1/admin/users",
		authHandler.GetAllUsers,
		middleware.JWTMiddleware,
		middleware.RoleMiddleware("admin"),
	)
}
