package server

import (
	"context"
	"net/http"

	"example.com/simple-crud/internal/config"
	"example.com/simple-crud/internal/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	httpServer *http.Server
	router     *mux.Router
}

func NewServer(cfg *config.Config) *Server {
	serverAddr := cfg.SrvCfg.Host + ":" + cfg.SrvCfg.Port
	srv := &http.Server{
		Addr:           serverAddr,
		ReadTimeout:    cfg.SrvCfg.ReadTimeout,
		WriteTimeout:   cfg.SrvCfg.WriteTimeout,
		MaxHeaderBytes: cfg.SrvCfg.MaxHeaderBytes,
	}

	s := &Server{
		httpServer: srv,
	}

	s.router = routers.New()
	s.router.Use(cors.AllowAll().Handler)

	srv.Handler = s.router

	return s
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
