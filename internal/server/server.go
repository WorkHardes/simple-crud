package server

import (
	"context"
	"net/http"

	"example.com/simple-crud/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	serverAddr := cfg.SrvCfg.Host + ":" + cfg.SrvCfg.Port
	s := &http.Server{
		Addr:           serverAddr,
		Handler:        handler,
		ReadTimeout:    cfg.SrvCfg.ReadTimeout,
		WriteTimeout:   cfg.SrvCfg.WriteTimeout,
		MaxHeaderBytes: cfg.SrvCfg.MaxHeaderBytes,
	}

	return &Server{
		httpServer: s,
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
