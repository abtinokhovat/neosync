package http

import (
	"errors"
	"fmt"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	middlewares "neosync/delivery/http/middleware"
	"neosync/internal/logger"
	"net/http"
	"os"
	"time"
)

// Router is function which register a bunch of handlers on echo instance
type Router interface {
	SetRoutes(e *echo.Echo)
}

type Config struct {
	EnableBanner bool          `koanf:"enable_banner"`
	Port         int           `koanf:"port"`
	Timeout      time.Duration `koanf:"timeout"`
	MetricPort   int           `koanf:"metric_port"`
}

type Server struct {
	cfg     Config
	Router  *echo.Echo
	routers *[]Router
}

func NewServer(cfg Config, routers []Router) *Server {
	return &Server{
		Router:  echo.New(),
		cfg:     cfg,
		routers: &routers,
	}
}

func (s *Server) Start() {
	s.Router.HideBanner = !s.cfg.EnableBanner

	// pre-routing middleware
	s.Router.Pre(middleware.RemoveTrailingSlash())

	// register middlewares
	s.middlewares()

	go s.serveMetrics()

	// register routes
	s.routes()

	// start server
	s.start()
}

// routes registers routes on the server
func (s *Server) routes() {
	echo.NotFoundHandler = handleNotFound
	echo.MethodNotAllowedHandler = handleNotAllowed

	// Routes
	s.Router.GET("/healthz", handleHealthCheck)

	// register routes from handlers
	if len(*s.routers) <= 0 {
		logger.L().Warn("no routers were available to register")
		return
	}

	for _, router := range *s.routers {
		router.SetRoutes(s.Router)
	}
}

// serveMetrics run a prometheus friendly metric http server
func (s *Server) serveMetrics() {
	logger.L().Info(fmt.Sprintf("started metric server on port %d", s.cfg.MetricPort))
	metrics := echo.New()
	metrics.HideBanner = true

	metrics.Use(middleware.Recover())

	metrics.GET("/metrics", echoprometheus.NewHandler())
	if err := metrics.Start(fmt.Sprintf(":%d", s.cfg.MetricPort)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.L().Error("metric server start error", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

// middlewares registers middlewares on the server
func (s *Server) middlewares() {
	s.Router.Use(middlewares.CORS())
	s.Router.Use(middleware.RequestID())
	s.Router.Use(middleware.Recover())
	s.Router.Use(middlewares.GZIP())
	s.Router.Use(middlewares.Logger())
	s.Router.Use(echoprometheus.NewMiddleware("neosync"))

	// timeout middlewares should be registered at the end off the middleware chain
	logger.L().Info("timeout", slog.Duration("timeout", s.cfg.Timeout))
	s.Router.Use(middlewares.ContextTimeout(s.cfg.Timeout))
	s.Router.Use(middlewares.Timeout(handleTimeout))
}

// start starts the http server
func (s *Server) start() {
	// Start server
	address := fmt.Sprintf(":%d", s.cfg.Port)
	logger.L().Info(fmt.Sprintf("started http server on %s", address))
	if err := s.Router.Start(address); err != nil {
		logger.L().Error("failed to start the http server", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
