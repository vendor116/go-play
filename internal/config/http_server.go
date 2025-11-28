package config

import (
	"errors"
	"time"
)

type HTTPServer struct {
	Host              string        `mapstructure:"host"`
	Port              string        `mapstructure:"port"`
	ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`
	ShutdownTimeout   time.Duration `mapstructure:"shutdown_timeout"`
}

func (as HTTPServer) Validate() error {
	if as.Port == "" {
		return errors.New("port is required")
	}
	if as.Host == "" {
		return errors.New("host is required")
	}
	if as.ReadHeaderTimeout == 0 {
		return errors.New("read_header_timeout is required")
	}
	if as.ShutdownTimeout == 0 {
		return errors.New("shutdown_timeout is required")
	}
	return nil
}
