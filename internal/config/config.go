package config

import (
	"errors"
	"fmt"
)

type Config struct {
	APIServer `mapstructure:"api_server"`

	LogLevel string `mapstructure:"log_level"`
	Debug    bool   `mapstructure:"debug"`
}

func (c Config) Validate() error {
	if err := c.APIServer.Validate(); err != nil {
		return fmt.Errorf("invalid api_server config: %w", err)
	}
	if c.LogLevel == "" {
		return errors.New("log_level is required")
	}
	return nil
}
