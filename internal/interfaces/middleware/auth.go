package middleware

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// BasicAuthMiddleware authenticates requests using basic auth
func BasicAuthMiddleware(username, password string) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            u, p, ok := c.Request().BasicAuth()
            if !ok || u != os.Getenv("ADMIN_USERNAME") || p != os.Getenv("ADMIN_PASSWORD") {
                return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
            }
            return next(c)
        }
    }
}
