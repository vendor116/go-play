package http

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"

	"github.com/vendor116/go-play/internal/config"
	"golang.org/x/sync/errgroup"
)

func RunHTTPServer(ctx context.Context, g *errgroup.Group, handler http.Handler, cfg config.HTTPServer) {
	server := &http.Server{
		Addr:              net.JoinHostPort(cfg.Host, cfg.Port),
		Handler:           handler,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
	}

	logger := slog.Default().With("address", server.Addr)

	g.Go(func() error {
		logger.Info("starting http server")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		logger.Info("http server shutdown gracefully")
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		if err := context.Cause(ctx); err != nil && !errors.Is(err, context.Canceled) {
			return err
		}

		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()

		logger.Info("shutting down http server")
		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.ErrorContext(shutdownCtx, "failed to shutdown http server", "error", err)
		}
		return nil
	})
}
