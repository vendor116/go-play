package internal

import (
	"log/slog"
	"os"
)

func SetupLogger(name, version, level string) error {
	var l slog.Level

	if err := l.UnmarshalText([]byte(level)); err != nil {
		return err
	}

	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: l,
	})

	slog.SetDefault(slog.New(h).With(
		slog.String("name", name),
		slog.String("version", version),
	))

	return nil
}
