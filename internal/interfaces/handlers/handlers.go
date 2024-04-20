package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HelloHandler responds to the "/hello" route
func HelloHandler(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World from TryMerge!")
}

func SecurePage(c echo.Context) error {
    return c.String(http.StatusOK, "You have accessed a secure area!")
}