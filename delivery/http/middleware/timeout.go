package middleware

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"time"
)

// ContextTimeout propagate http request context with timeout duration
func ContextTimeout(timeout time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			timeoutCtx, cancel := context.WithTimeout(c.Request().Context(), timeout)
			c.SetRequest(c.Request().WithContext(timeoutCtx))
			defer cancel()
			return next(c)
		}
	}
}

// Timeout runs request in an extra goroutine and listens for the context cancellation
func Timeout(timeoutHandler echo.HandlerFunc) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			doneCh := make(chan error, 1)

			go func() {
				// capture the result of the handler and send it to doneCh
				doneCh <- next(c)
			}()

			select {
			case res := <-doneCh:
				return res
			case <-c.Request().Context().Done():
				if errors.Is(c.Request().Context().Err(), context.DeadlineExceeded) {
					doneCh <- timeoutHandler(c)
				} else {
					doneCh <- c.Request().Context().Err()
				}

				return <-doneCh
			}
		}
	}
}
