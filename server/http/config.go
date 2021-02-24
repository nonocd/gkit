package http

import "time"

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
		RunMode:  "",
		HTTPPort: 9000,
	}
}
