package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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
		*Probes
	}
)

type (
	EmptyResponse string
)

const emptyRes = ""

func NewHandler(svc service.AssetService, log log.Logger, cfg *cfg.Config) *Handler {
	return &Handler{
		Core:    sys.NewCore(log, cfg),
		service: svc,
		Probes:  NewProbes(),
	}
}

func (h *Handler) handleAPIV1(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")

	userIDIndex := -1
	for i, segment := range pathSegments {
		_, err := uuid.Parse(segment)
		if err == nil {
			userIDIndex = i
			break
		}
	}

	if userIDIndex == -1 {
		http.NotFound(w, r)
		return
	}

	userIDSegment := pathSegments[userIDIndex]
	ctx := context.WithValue(r.Context(), UserIDCtxKey, userIDSegment)
	r = r.WithContext(ctx)

	resIDSegment := pathSegments[len(pathSegments)-1]
	resID, err := uuid.Parse(resIDSegment)
	if err != nil {
		h.Log().Debug("Not a resource URL:", r.URL.Path)
	}

	ctx = context.WithValue(r.Context(), ResIDCtxKey, resID)
	r = r.WithContext(ctx)

	resource := strings.ToLower(pathSegments[userIDIndex+1])

	switch resource {
	case "assets":
		h.handleAssets(w, r)
	case "faved":
		h.handleFavs(w, r)
	default:
		h.handleError(w, InvalidResourceErr)
	}
}

func (h *Handler) handleAssets(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAssets(w, r)
	case http.MethodPut:
		h.updateAsset(w, r)
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
		h.updateAsset(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		h.handleError(w, MethodNotAllowedErr)
	}
}

func (h *Handler) getAssets(w http.ResponseWriter, r *http.Request) {
	userID, ok := h.userID(r)
	if !ok {
		h.handleError(w, NoUserErr)
		return
	}

	repo := h.service.Repo()
	assets, err := repo.GetAssets(r.Context(), userID)
	if err != nil {
		h.handleError(w, errors.Wrap("get assets error", err))
		return
	}

	msg := fmt.Sprintf("count: %d", len(assets))

	h.handleSuccess(w, assets, msg)
}

func (h *Handler) getFaved(w http.ResponseWriter, r *http.Request) {
	userID, ok := h.userID(r)
	if !ok {
		h.handleError(w, NoUserErr)
		return
	}

	repo := h.service.Repo()
	assets, err := repo.GetAssets(r.Context(), userID, port.Faved)
	if err != nil {
		h.handleError(w, errors.Wrap("get faved error", err))
		return
	}

	msg := fmt.Sprintf("count: %d", len(assets))

	h.handleSuccess(w, assets, msg)
}

func (h *Handler) updateAsset(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, InvalidRequestErr)
	}

	var req AssetRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		h.handleError(w, InvalidRequestDataErr)
	}

	ctx := context.WithValue(r.Context(), AssetReqCtxKey, req)
	r = r.WithContext(ctx)

	switch req.Action {
	case "fav":
		h.favAsset(w, r)
	case "unfav":
		h.unfavAsset(w, r)
	case "update":
		h.updateName(w, r)
	default:
		http.Error(w, "Invalid action", http.StatusBadRequest)
	}
}

func (h *Handler) favAsset(w http.ResponseWriter, r *http.Request) {
	userID, ok := h.userID(r)
	if !ok {
		h.handleError(w, NoUserErr)
	}

	resID, ok := h.resourceID(r)
	if !ok {
		h.handleError(w, NoResourceErr)
		return
	}

	req, ok := h.assetReq(r)
	if !ok {
		h.handleError(w, NoAssetReqErr)
		return
	}

	repo := h.service.Repo()
	err := repo.AddFav(r.Context(), userID, req.Type, resID)

	if err != nil {
		h.handleError(w, errors.Wrap("add fav error", err))
		return
	}

	//msg := fmt.Sprintf("Reg. count: %d", len(assets))

	h.handleSuccess(w, emptyRes)
}

func (h *Handler) unfavAsset(w http.ResponseWriter, r *http.Request) {
	userID, ok := h.userID(r)
	if !ok {
		h.handleError(w, NoUserErr)
	}

	resID, ok := h.resourceID(r)
	if !ok {
		h.handleError(w, NoResourceErr)
		return
	}

	req, ok := h.assetReq(r)
	if !ok {
		h.handleError(w, NoAssetReqErr)
		return
	}

	repo := h.service.Repo()
	err := repo.RemoveFav(r.Context(), userID, req.Type, resID)

	if err != nil {
		h.handleError(w, errors.Wrap("add fav error", err))
		return
	}

	//msg := fmt.Sprintf("Reg. count: %d", len(assets))

	h.handleSuccess(w, emptyRes)
}

func (h *Handler) updateName(w http.ResponseWriter, r *http.Request) {
	userID, ok := h.userID(r)
	if !ok {
		h.handleError(w, NoUserErr)
	}

	resID, ok := h.resourceID(r)
	if !ok {
		h.handleError(w, NoResourceErr)
		return
	}

	req, ok := h.assetReq(r)
	if !ok {
		h.handleError(w, NoAssetReqErr)
		return
	}

	repo := h.service.Repo()
	err := repo.UpdateFav(r.Context(), userID, req.Type, resID, req.Name)

	if err != nil {
		h.handleError(w, errors.Wrap("update fav error", err))
		return
	}

	//msg := fmt.Sprintf("Reg. count: %d", len(assets))

	h.handleSuccess(w, emptyRes)
}
