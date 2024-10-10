package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
)

const (
	defaultFilePath        = "logs/logs.json"
	defaultUseLocalTime    = false
	defaultFileMaxSizeInMB = 10
	defaultFileAgeInDays   = 30
)

type Config struct {
	FilePath         string
	UseLocalTime     bool
	FileMaxSizeInMB  int
	FileMaxAgeInDays int
}

// l singleton logger instance
var l *slog.Logger

var defaultOptions = slog.HandlerOptions{
	Level: slog.LevelDebug,
}

// init wake logger with default config
func init() {
	setDefaultLogger()
}

func setDefaultLogger() {
	fileWriter := &lumberjack.Logger{
		Filename:  defaultFilePath,
		LocalTime: defaultUseLocalTime,
		MaxSize:   defaultFileMaxSizeInMB,
		MaxAge:    defaultFileAgeInDays,
	}

	l = slog.New(
		slog.NewJSONHandler(io.MultiWriter(fileWriter, os.Stdout), &defaultOptions),
	)
}

// L logger instance
func L() *slog.Logger {
	return l
}

// New makes a logger with the customized config
func New(cfg Config, opt *slog.HandlerOptions) *slog.Logger {
	fileWriter := &lumberjack.Logger{
		Filename:  cfg.FilePath,
		LocalTime: cfg.UseLocalTime,
		MaxSize:   cfg.FileMaxSizeInMB,
		MaxAge:    cfg.FileMaxAgeInDays,
	}

	logger := slog.New(
		slog.NewJSONHandler(io.MultiWriter(fileWriter, os.Stdout), opt),
	)

	return logger
}
