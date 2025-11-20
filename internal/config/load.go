package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	defaultEnvPrefix  = "APP"
	defaultConfigPath = "./config.yaml"
)

func Load(path string, options ...func() error) (*App, error) {
	viper.SetConfigFile(path)
	if path == "" {
		viper.SetConfigFile(defaultConfigPath)
	}

	for _, option := range options {
		if err := option(); err != nil {
			return nil, err
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("config file not found: %w", err)
		}

		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg App
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &cfg, nil
}

func WithEnvs(prefix string) func() error {
	return func() error {
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		viper.SetEnvPrefix(defaultEnvPrefix)
		if prefix != "" {
			viper.SetEnvPrefix(prefix)
		}

		return nil
	}
}
