package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"

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

const (
	emptyRes    = ""
	defPageSize = 12
)

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
		h.handleFaved(w, r)
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

func (h *Handler) handleFaved(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getFaved(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		h.handleError(w, MethodNotAllowedErr)
	}
}

func (h *Handler) getAssets(w http.ResponseWriter, r *http.Request) {
	page, size := h.pageAndSize(r)

	repo := h.service.Repo()
	assets, pages, err := repo.GetAssets(r.Context(), page, size)
	if err != nil {
		h.handleError(w, errors.Wrap("get assets error", err))
		return
	}

	h.handleSuccess(w, assets, len(assets), pages)
}

func (h *Handler) getFaved(w http.ResponseWriter, r *http.Request) {
	userID, ok := h.userID(r)
	if !ok {
		h.handleError(w, NoUserErr)
		return
	}

	repo := h.service.Repo()
	assets, pages, err := repo.GetFaved(r.Context(), userID, 1, 40)
	if err != nil {
		h.handleError(w, errors.Wrap("get faved error", err))
		return
	}

	h.handleSuccess(w, assets, len(assets), pages)
}

func (h *Handler) createAsset(w http.ResponseWriter, r *http.Request) {
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

	default:
		http.Error(w, "Invalid action", http.StatusBadRequest)
	}
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

func (h *Handler) deleteAsset(w http.ResponseWriter, r *http.Request) {
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
	case "unfav":
		h.unfavAsset(w, r)

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

	h.handleSuccess(w, emptyRes, 1, 0)
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
		h.handleError(w, errors.Wrap("remove fav error", err))
		return
	}

	h.handleSuccess(w, emptyRes, 1, 0)
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
	err := repo.UpdateFav(r.Context(), userID, req.Type, resID, req.Name, req.Description)

	if err != nil {
		h.handleError(w, errors.Wrap("update fav error", err))
		return
	}

	h.handleSuccess(w, emptyRes, 1, 0)
}

func (h *Handler) pageAndSize(r *http.Request) (page, size int) {
	params := r.URL.Query()

	p := params.Get("page")
	if p == "" {
		page = 1
	}

	s := params.Get("size")
	if s == "" {
		s = fmt.Sprintf("%d", h.Cfg().Prop.PageSize)
	}

	page, err := strconv.Atoi(p)
	if err != nil {
		page = 1
	}

	size, err = strconv.Atoi(s)
	if err != nil {
		size = defPageSize
	}

	return page, size
}
