package http

import (
	"fmt"
	"github.com/Alp4ka/classifier-mango/internal/manager"

	globaltelemetry "github.com/Alp4ka/classifier-mango/internal/telemetry"
	"github.com/gofiber/fiber/v2"

	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	"github.com/ansrivas/fiberprometheus/v2"
)

const (
	_appName = "http-api"
)

type Server struct {
	app         *fiber.App
	cfg         Config
	coreManager manager.Manager
}

func NewHTTPServer(cfg Config) *Server {
	server := &Server{
		cfg:         cfg,
		coreManager: manager.NewManager(cfg.CoreClient),
	}
	server.app = fiber.New(
		fiber.Config{
			AppName:               _appName,
			ErrorHandler:          server.mwErrorHandler,
			DisableStartupMessage: true,
			DisableKeepalive:      true,
		},
	)

	return server
}

func (s *Server) configureRouting() {
	// Middlewares.
	s.app.Use(s.mwRecoverer())
	s.app.Use(s.mwCors())
	s.app.Use(s.mwRequestID())
	s.app.Use(s.mwLogging())

	// API Group
	apiGroup := s.app.Group("/api")
	apiGroup.Use(s.mwRateLimiter())
	apiGroup.Use(s.mwContentChecker())

	sessionGroup := apiGroup.Group("/session")
	sessionGroup.Post("", s.hStart)
	sessionGroup.Delete("/:session_id", s.hFinish)

	actionGroup := apiGroup.Group("/action")
	actionGroup.Post("", s.hActionPost)
}

func (s *Server) Run() error {
	s.configureRouting()
	mlogger.L().Info("Listening HTTP API server", field.Int("port", s.cfg.Port))

	return s.app.Listen(fmt.Sprintf(":%d", s.cfg.Port))
}

func (s *Server) WithMetrics() *Server {
	prometheus := fiberprometheus.New(globaltelemetry.Namespace)
	prometheus.RegisterAt(s.app, "/metrics")

	return s
}

func (s *Server) Close() error {
	return s.app.Shutdown()
}
