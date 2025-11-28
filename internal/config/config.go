package config

import (
	"errors"
	"fmt"
)

type Config struct {
	HTTPServer `mapstructure:"http_server"`

	LogLevel string `mapstructure:"log_level"`
}

func (c Config) Validate() error {
	if err := c.HTTPServer.Validate(); err != nil {
		return fmt.Errorf("invalid http_server config: %w", err)
	}

	if c.LogLevel == "" {
		return errors.New("log_level is required")
	}

	return nil
}
