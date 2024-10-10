package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func handleHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "everything is good!"})
}

func handleTimeout(c echo.Context) error {
	return c.JSON(http.StatusServiceUnavailable, echo.Map{"message": "request timeout"})
}

func handleNotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, echo.Map{"message": "object not found"})
}

func handleNotAllowed(c echo.Context) error {
	return c.JSON(http.StatusMethodNotAllowed, echo.Map{"message": "method not allowed"})
}
