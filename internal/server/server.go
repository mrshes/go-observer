package server

import (
	"context"
	"first-project/internal/config"
	"first-project/pkg/e"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.Config, handlers http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.Server.PORT,
			Handler: handlers,
		},
	}
}

func (s *Server) Run() error {
	e.Info("Server start!", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
