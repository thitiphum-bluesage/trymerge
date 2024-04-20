package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thitiphum-bluesage/trymerge/internal/interfaces/handlers"
)

// RegisterRoutes configures the available routes in the application
func RegisterRoutes(e *echo.Echo) {
    e.GET("/hello", handlers.HelloHandler)
}
