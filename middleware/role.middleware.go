package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RoleMiddleware(requiredRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			role, ok := c.Get("role").(string)
			// role := c.Get("role").(string)
			fmt.Println("Current Role:", role)
			fmt.Println("Required Role:", requiredRole)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "role not found",
				})
			}


			if role != requiredRole {
				return c.JSON(http.StatusForbidden, map[string]string{
					"error": "access denied",
				})
			}

			return next(c)
		}
	}
}
