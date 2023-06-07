package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jobquestvault/platform-go-challenge/internal/sys"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/cfg"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/log"
)

var (
	MethodNotAllowedErr = sys.NewError("method not allowed")
)

type Server struct {
	sys.Core
	port     int
	router   *http.ServeMux
	handlers Handler
}

func NewServer(log log.Logger, cfg *cfg.Config) *Server {
	return &Server{
		Core:   sys.NewCore(log, cfg),
		port:   cfg.Server.Port,
		router: http.NewServeMux(),
	}
}

func (s *Server) Setup(ctx context.Context) {
	s.router.HandleFunc("/assets/", s.handlers.handleAssets)
	s.router.HandleFunc("/favs/", s.handlers.handleFavs)
}

func (s *Server) Start(ctx context.Context) error {
	addr := fmt.Sprintf(":%d", s.port)
	s.Log().Info("Server listening on port:", s.port)
	return http.ListenAndServe(addr, nil)
}