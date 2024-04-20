package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/thitiphum-bluesage/trymerge/internal/interfaces/routes"
)

func main() {
    e := echo.New()

	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }


    // Register routes
    routes.RegisterRoutes(e)

    // Start server
	port := os.Getenv("SERVER_PORT")
    e.Logger.Fatal(e.Start(":" + port))
}
