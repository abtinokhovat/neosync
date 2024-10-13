package http

import (
	"errors"
	"github.com/labstack/echo/v4"
	"neosync/pkg/richerror"
	"net/http"
)

func Ok(c echo.Context, res any) error {
	return c.JSON(http.StatusOK, res)
}

func Error(err error) *echo.HTTPError {
	msg, code := mapError(err)
	return echo.NewHTTPError(code, msg)
}

func mapError(err error) (string, int) {
	var richError richerror.RichError
	switch {
	case errors.As(err, &richError):
		var re richerror.RichError
		errors.As(err, &re)
		msg := re.Message()

		code := mapKindToHTTPStatusCode(re.Kind())

		// we should not expose unexpected error messages
		if code >= 500 && msg == "" {
			msg = "something went wrong"
		}

		return msg, code
	default:
		return err.Error(), http.StatusBadRequest
	}
}

func MapHTTPStatusCodeToKind(code int) richerror.Kind {
	switch code {
	case http.StatusUnprocessableEntity:
		return richerror.KindInvalid
	case http.StatusNotFound:
		return richerror.KindNotFound
	case http.StatusForbidden:
		return richerror.KindForbidden
	case http.StatusInternalServerError:
		return richerror.KindUnexpected
	default:
		return richerror.KindInvalid
	}
}

func mapKindToHTTPStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindInvalid:
		return http.StatusUnprocessableEntity
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindForbidden:
		return http.StatusForbidden
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
