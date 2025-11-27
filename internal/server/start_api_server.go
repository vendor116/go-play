package server

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"

	"github.com/vendor116/playgo/internal/config"
	"golang.org/x/sync/errgroup"
)

func StartAPIServer(ctx context.Context, handler http.Handler, cfg config.APIServer) error {
	server := &http.Server{
		Addr:              net.JoinHostPort(cfg.Host, cfg.Port),
		Handler:           handler,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
	}

	logger := slog.Default().With("addr", server.Addr)

	g, gCtx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		logger.Info("starting api server")

		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		logger.Info("api server shutdown gracefully")
		return nil
	})

	g.Go(func() error {
		select {
		case <-gCtx.Done():
			return gCtx.Err()
		case <-ctx.Done():
			logger.Warn("shutting down api server")

			shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
			defer cancel()

			if err := server.Shutdown(shutdownCtx); err != nil {
				logger.ErrorContext(shutdownCtx, "failed to shutdown api server", "error", err)
			}

			return nil
		}
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}
