package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")
		fmt.Printf("Header = %q\n", authHeader)
		fmt.Printf("Header = %q\n", authHeader)
		fmt.Println(strings.HasPrefix(authHeader, "Bearer "))
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing authorization header",
			})
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid authorization header",
			})
		}
		fmt.Println("Authorization Header:", authHeader)
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		// fmt.Println("JWT_SECRET:", os.Getenv("JWT_SECRET"))

		if err != nil {
			fmt.Println("JWT Parse Error:", err)
		}

		if token != nil {
			fmt.Println("Token Valid:", token.Valid)
		}

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid or expired token",
			})
		}

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid or expired token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid token claims",
			})
		}

		c.Set("userID", claims["user_id"])
		c.Set("email", claims["email"])
		c.Set("role", claims["role"])
		return next(c)
	}
}
