package internal

import (
	"log/slog"
	"os"
)

func SetupLogger(level string, version string) error {
	var lvl slog.Level
	if err := lvl.UnmarshalText([]byte(level)); err != nil {
		return err
	}

	slog.SetDefault(
		slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{
					Level: lvl,
				},
			),
		).With(
			slog.String("version", version),
		),
	)
	return nil
}
