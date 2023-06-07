package app

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	server Server
	//assetService AssetService
}

type Server struct {
	port     int
	router   *http.ServeMux
	handlers Handler
}

func NewServer(port int) *Server {
	return &Server{
		port:   port,
		router: http.NewServeMux(),
	}
}

func (s *Server) Setup() {
	s.router.HandleFunc("/assets/", s.handleAssets)
	s.router.HandleFunc("/favs/", s.handleFavs)
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.port)

	log.Printf("Server listening on port %d\n", s.port)
	return http.ListenAndServe(addr, nil)
}
