package app

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handleAssets(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getAssets(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		//err := fmt.Fprintln(w, "Method not allowed")
	}
}

func (s *Server) handleFavs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getFaved(w, r)
	case http.MethodPut:
		s.updateFav(w, r)
	case http.MethodDelete:
		s.removeFav(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		//err := fmt.Fprintln(w, "Method not allowed")
	}
}

func (s *Server) getAssets(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(assets)
	if err != nil {
		// TODO: Implement after defining handling strategy
	}
}

func (s *Server) getFaved(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(assets)
	if err != nil {
		// TODO: Implement after defining handling strategy
	}
}

func (s *Server) addFav(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handling for "/favorites/add" endpoint
	panic("not implemented yet")
}

func (s *Server) removeFav(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handling for "/favorites/remove" endpoint
	panic("not implemented yet")
}

func (s *Server) updateFav(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handling for "/favorites/edit" endpoint
	panic("not implemented yet")
}
