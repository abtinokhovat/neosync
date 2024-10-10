package middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"neosync/internal/logger"
)

func Logger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:           true,
		LogStatus:        true,
		LogHost:          true,
		LogRemoteIP:      true,
		LogRequestID:     true,
		LogMethod:        true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogLatency:       true,
		LogProtocol:      true,
		LogError:         true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			slogAttributes := []slog.Attr{
				slog.String("uri", v.URI),
				slog.Int("status", v.Status),
				slog.String("host", v.Host),
				slog.String("remote_ip", v.RemoteIP),
				slog.String("request_id", v.RequestID),
				slog.String("method", v.Method),
				slog.String("content_length", v.ContentLength),
				slog.Int64("response_size", v.ResponseSize),
				slog.Int64("latency_ms", v.Latency.Milliseconds()),
				slog.String("protocol", v.Protocol),
			}

			msg := "request"
			level := slog.LevelInfo
			if v.Error != nil {
				slogAttributes = append(slogAttributes, slog.String("error", v.Error.Error()))
				level = slog.LevelError
			}

			logger.L().WithGroup("echo").LogAttrs(context.Background(), level, msg,
				slogAttributes...,
			)

			return nil
		},
	})
}
