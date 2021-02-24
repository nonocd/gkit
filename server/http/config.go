package http

import (
	"time"
)

// CfgOption is config option
type CfgOption func(*Config)

// Config is the main struct for http server
type Config struct {
	AppName       string // Application name
	RunMode       string // Running Mode: dev | prod
	Version       string
	HTTPAddr      string
	HTTPPort      int
	ServerTimeOut time.Duration
}

var (
	// BConfig is the default config for Application
	BConfig *Config
)

func init() {
	BConfig = newConfig()
}

func newConfig() *Config {
	return &Config{
		RunMode:       ProdMode,
		HTTPPort:      9000,
		ServerTimeOut: 60,
	}
}

// WithAppName set app name
func WithAppName(appname string) CfgOption {
	return func(cfg *Config) {
		cfg.AppName = appname
	}
}

// WithVersion set app version
func WithVersion(version string) CfgOption {
	return func(cfg *Config) {
		cfg.Version = version
	}
}
