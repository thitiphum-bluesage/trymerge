package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thitiphum-bluesage/trymerge/internal/interfaces/handlers"
	"github.com/thitiphum-bluesage/trymerge/internal/interfaces/middleware"
)

// RegisterRoutes configures the available routes in the application
func RegisterRoutes(e *echo.Echo) {
    e.GET("/hello", handlers.HelloHandler)
    e.GET("/secure", handlers.SecurePage, middleware.BasicAuthMiddleware("user", "pass"))
}
