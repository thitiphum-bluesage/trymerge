package main

import (
	"github.com/labstack/echo/v4"
	"github.com/thitiphum-bluesage/trymerge/internal/interfaces/routes"
)

func main() {
    e := echo.New()

    // Register routes
    routes.RegisterRoutes(e)

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}
