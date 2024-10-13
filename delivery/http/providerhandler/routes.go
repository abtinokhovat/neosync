package providerhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(e *echo.Echo) {
	v1 := e.Group("/v1/providers")

	v1.GET("/", h.getProvidersSorted)
}
