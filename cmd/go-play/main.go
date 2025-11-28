package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/vendor116/go-play/internal"
	"github.com/vendor116/go-play/internal/config"
	"github.com/vendor116/go-play/internal/http"
	"github.com/vendor116/go-play/internal/http/serverv1"
	"golang.org/x/sync/errgroup"
)

var version = "dev"

func main() {
	os.Exit(run())
}

func run() int {
	var cfgPath string
	flag.StringVar(&cfgPath, "config", "", "path to config file")
	flag.Parse()

	cfg, err := config.Load[config.Config](cfgPath)
	if err != nil {
		slog.Default().Error("failed to load config", "error", err)
		return 1
	}

	err = cfg.Validate()
	if err != nil {
		slog.Default().Error("failed to validate config", "error", err)
		return 1
	}

	err = internal.SetupLogger(cfg.LogLevel, version)
	if err != nil {
		slog.Default().Error("failed to setup logger", "error", err)
		return 1
	}

	httpService := serverv1.RegisterRoutes(
		serverv1.NewServer(),
	)

	eg, egCtx := errgroup.WithContext(context.Background())
	ctx, cancel := signal.NotifyContext(egCtx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	http.RunHTTPServer(ctx, eg, httpService, cfg.HTTPServer)

	if egErr := eg.Wait(); egErr != nil && !errors.Is(egErr, context.Canceled) {
		slog.Default().Error("application completed", "error", egErr)
		return 1
	}

	return 0
}
