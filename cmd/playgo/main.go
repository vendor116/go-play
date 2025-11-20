package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"

	"github.com/vendor116/playgo/internal"
	"github.com/vendor116/playgo/internal/api"
	"github.com/vendor116/playgo/internal/config"
	"github.com/vendor116/playgo/internal/generated"
	"github.com/vendor116/playgo/internal/http"
)

var (
	name    = "playgo"
	version = "dev"
)

func main() {
	var cfgPath, envPrefix string
	flag.StringVar(&cfgPath, "config", "config.yaml", "path to config file")
	flag.StringVar(&envPrefix, "env-prefix", "APP", "environment prefix")
	flag.Parse()

	cfg, err := config.Load(cfgPath, config.WithEnvs(envPrefix))
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	err = internal.SetupLogger(name, version, cfg.LogLevel)
	if err != nil {
		log.Fatalf("failed to setup logger: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	server := api.NewServer()

	http.RunServer(
		ctx,
		generated.HandlerFromMux(server, api.GetChiRouter(server)),
		cfg.HTTPServer.Host,
		cfg.HTTPServer.Port,
	)
}
