package app

import (
	"context"
	"errors"

	"github.com/Alp4ka/classifier-mango/internal/api/http"
)

type App struct {
	cfg Config

	httpServer *http.Server
}

func New(cfg Config) *App {
	return &App{
		cfg: cfg,
		httpServer: http.NewHTTPServer(
			http.Config{
				RateLimit:  cfg.HTTPRateLimit,
				Port:       cfg.HTTPPort,
				APIKey:     cfg.APIKey,
				CoreClient: cfg.CoreClient,
			},
		).WithMetrics(),
	}
}

func (a *App) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func(errCh chan<- error) {
		errCh <- a.httpServer.Run()
	}(errCh)

	select {
	case <-ctx.Done():
		return nil
	case err := <-errCh:
		return err
	}
}

func (a *App) Close() (err error) {
	return errors.Join(a.httpServer.Close())
}
