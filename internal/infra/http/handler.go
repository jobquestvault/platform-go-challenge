package http

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/jobquestvault/platform-go-challenge/internal/domain/port"
	"github.com/jobquestvault/platform-go-challenge/internal/domain/service"
	"github.com/jobquestvault/platform-go-challenge/internal/sys"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/cfg"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/errors"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/log"
)

type (
	Handler struct {
		sys.Core
		service service.AssetService
	}
)

func NewHandler(svc service.AssetService, log log.Logger, cfg *cfg.Config) *Handler {
	return &Handler{
		Core:    sys.NewCore(log, cfg),
		service: svc,
	}
}

func (h *Handler) handleAPIV1(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")

	uuidIndex := -1
	for i, segment := range pathSegments {
		_, err := uuid.Parse(segment)
		if err == nil {
			uuidIndex = i
			break
		}
	}

	if uuidIndex == -1 {
		http.NotFound(w, r)
		return
	}

	uuidSegment := pathSegments[uuidIndex]
	ctx := context.WithValue(r.Context(), UserCtxKey, uuidSegment)
	r = r.WithContext(ctx)

	resource := strings.ToLower(pathSegments[uuidIndex+1])

	switch resource {
	case "assets":
		h.handleAssets(w, r)
	case "favs":
		h.handleFavs(w, r)
	default:
		h.handleError(w, InvalidResourceErr)
	}
}

func (h *Handler) handleAssets(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAssets(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		h.handleError(w, MethodNotAllowedErr)
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
		h.handleError(w, MethodNotAllowedErr)
	}
}

func (h *Handler) getAssets(w http.ResponseWriter, r *http.Request) {
	userID, ok := h.userID(r)
	if !ok {
		h.handleError(w, NoUserErr)
	}

	repo := h.service.Repo()
	assets, err := repo.GetAssets(r.Context(), userID)
	if err != nil {
		h.handleError(w, errors.Wrap("get assets error", err))
	}

	msg := fmt.Sprintf("Reg. count: %d", len(assets))

	h.handleSuccess(w, assets, msg)
}

func (h *Handler) getFaved(w http.ResponseWriter, r *http.Request) {
	userID, ok := h.userID(r)
	if !ok {
		h.handleError(w, NoUserErr)
	}

	repo := h.service.Repo()
	assets, err := repo.GetAssets(r.Context(), userID, port.Faved)
	if err != nil {
		h.handleError(w, errors.Wrap("get faved error", err))
	}

	msg := fmt.Sprintf("Reg. count: %d", len(assets))

	h.handleSuccess(w, assets, msg)
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
