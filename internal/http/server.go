package http

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"time"
)

const (
	readHeaderTimeout = 3 * time.Second
	shutdownTimeout   = 5 * time.Second
)

func RunServer(ctx context.Context, handler http.Handler, host, port string) {
	server := &http.Server{
		Addr:              net.JoinHostPort(host, port),
		Handler:           handler,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	httpCtx, stop := context.WithCancelCause(ctx)

	go func() {
		slog.Default().Info("starting http server", "addr", server.Addr)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			stop(err)
		}
	}()

	<-httpCtx.Done()

	if err := context.Cause(httpCtx); err != nil && !errors.Is(err, context.Canceled) {
		slog.Default().ErrorContext(httpCtx, "failed to start http server", "error", err)
		return
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Default().ErrorContext(shutdownCtx, "failed to shutdown http server", "error", err)
		return
	}

	slog.Default().InfoContext(httpCtx, "http server shutdown gracefully")
}
