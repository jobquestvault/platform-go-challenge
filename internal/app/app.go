package app

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Port int
}

func NewServer(port int) *Server {
	return &Server{
		Port: port,
	}
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.Port)
	http.HandleFunc("/favorites", s.favorites)
	http.HandleFunc("/favorites/add", s.addFavorite)
	http.HandleFunc("/favorites/remove", s.removeFavorite)
	http.HandleFunc("/favorites/update", s.updateFavorite)

	log.Printf("Server listening on port %d\n", s.Port)
	return http.ListenAndServe(addr, nil)
}
