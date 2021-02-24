package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"gkit/pkg/logs"
)

const (
	// DevMode is for develop
	DevMode = "dev"
	// ProdMode is for production
	ProdMode = "prod"
)

var sigs = []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, os.Interrupt}

// Server defines application
type Server struct {
	Handlers *gin.Engine
	Server   *http.Server
	Cfg      *Config
}

// NewServer will create a sever with the BConfig.
func NewServer(opts ...CfgOption) *Server {
	srv := NewServerWithCfg(BConfig)
	for _, o := range opts {
		o(srv.Cfg)
	}
	return srv
}

// NewServerWithCfg will create a sever with the cfg.
func NewServerWithCfg(cfg *Config) *Server {
	app := &Server{
		Server: &http.Server{},
		Cfg:    cfg,
	}
	app.setGinMode()
	return app
}

// setGinMode is set gin mode
func (app *Server) setGinMode() {
	var ginMode string
	switch app.Cfg.RunMode {
	case DevMode:
		ginMode = gin.DebugMode
	case ProdMode:
		ginMode = gin.ReleaseMode
	default:
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)
}

// Run application.
func (app *Server) Run() {
	app.Server.Handler = app.Handlers
	app.Server.ReadHeaderTimeout = time.Duration(app.Cfg.ServerTimeOut) * time.Second
	app.Server.WriteTimeout = time.Duration(app.Cfg.ServerTimeOut) * time.Second

	addr := fmt.Sprintf("%s:%d", app.Cfg.HTTPAddr, app.Cfg.HTTPPort)
	app.Server.Addr = addr

	logs.Infoln("http server starting")
	go func() {
		if err := app.Server.ListenAndServe(); err != nil {
			logs.Fatalf("Server err: %v", err)
		}
	}()
	logs.Infoln("http server Running on http://%s", app.Server.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, sigs...)
	<-quit
	app.Stop()
}

// Stop application.
func (app *Server) Stop() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := app.Server.Shutdown(ctx); err != nil {
		logs.Fatalf("Server Shutdown: %v", err)
		return
	}
	logs.Infoln("Server exiting")
}
