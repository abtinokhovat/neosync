package providerhandler

import (
	"github.com/labstack/echo/v4"
	"neosync/pkg/http"
)

func (h Handler) getProvidersSorted(c echo.Context) error {

	//TODO: replace this with the sorted function
	resp, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		return http.Error(err)
	}

	return http.Ok(c, resp)
}
