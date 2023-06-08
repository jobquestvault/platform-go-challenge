package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jobquestvault/platform-go-challenge/internal/domain/service"
	"github.com/jobquestvault/platform-go-challenge/internal/sys"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/cfg"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/errors"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/log"
)

type Server struct {
	sys.Core
	http.Server
	port     int
	router   *http.ServeMux
	handlers *Handler
}

var (
	MethodNotAllowedErr = errors.NewError("method not allowed")
)

const (
	apiV1 = "/api/v1"
)

func NewServer(svc service.AssetService, log log.Logger, cfg *cfg.Config) *Server {
	router := http.NewServeMux()

	srv := Server{
		Core:     sys.NewCore(log, cfg),
		port:     cfg.Server.Port,
		router:   router,
		handlers: NewHandler(svc, log, cfg),
	}

	srv.Server = http.Server{
		Addr:    srv.Address(),
		Handler: router,
	}

	return &srv

}

func (s *Server) Setup(ctx context.Context) {
	s.router.HandleFunc(apiV1+"/assets/", s.handlers.handleAssets)
	s.router.HandleFunc(apiV1+"/favs/", s.handlers.handleFavs)
}

func (s *Server) Start(ctx context.Context) error {
	s.Log().Info("Server listening on port:", s.port)
	return s.Server.ListenAndServe()
}

func (s *Server) Address() string {
	return fmt.Sprintf(":%d", s.port)
}
