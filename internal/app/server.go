package app

import (
	"encoding/json"
	"net/http"
)

func (s *Server) favorites(w http.ResponseWriter, r *http.Request) {
	favorites := []string{"Chart1", "Insight1", "Audience1"}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(favorites)
	if err != nil {
		// TODO: Implement after defining handling strategy
	}
}

func (s *Server) addFavorite(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handling for "/favorites/add" endpoint
	panic("not implemented yet")
}

func (s *Server) removeFavorite(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handling for "/favorites/remove" endpoint
	panic("not implemented yet")
}

func (s *Server) updateFavorite(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handling for "/favorites/edit" endpoint
	panic("not implemented yet")
}
