package http

import (
	"encoding/json"
	"net/http"

	"github.com/jobquestvault/platform-go-challenge/internal/domain/model"
	"github.com/jobquestvault/platform-go-challenge/internal/sys"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/cfg"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/log"
)

type (
	Handler struct {
		sys.Core
	}
)

func NewHandler(log log.Logger, cfg *cfg.Config) *Handler {
	return &Handler{
		Core: sys.NewCore(log, cfg),
	}
}

func (h *Handler) handleAssets(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAssets(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		h.Log().Error(MethodNotAllowedErr)
	}
}

func (h *Handler) handleFavs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getFaved(w, r)
	case http.MethodPut:
		h.updateFav(w, r)
	case http.MethodDelete:
		h.removeFav(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		h.Log().Error(MethodNotAllowedErr)
	}
}

func (h *Handler) getAssets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(model.Assets)
	if err != nil {
		// TODO: Implement after defining handling strategy
	}
}

func (h *Handler) getFaved(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(model.Assets)
	if err != nil {
		// TODO: Implement after defining handling strategy
	}
}

func (h *Handler) addFav(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handling for "/favorites/add" endpoint
	panic("not implemented yet")
}

func (h *Handler) removeFav(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handling for "/favorites/remove" endpoint
	panic("not implemented yet")
}

func (h *Handler) updateFav(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handling for "/favorites/edit" endpoint
	panic("not implemented yet")
}
