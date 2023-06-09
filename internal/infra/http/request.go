package http

import (
	"net/http"

	"github.com/google/uuid"
)

type ContextKey string

type AssetRequest struct {
	Type   string `json:"type"`
	Action string `json:"action"`
	Name   string `json:"name"`
}

const (
	UserIDCtxKey   = "user"
	ResIDCtxKey    = "resource"
	AssetReqCtxKey = "assetreq"
)

func (h *Handler) userID(r *http.Request) (userID string, ok bool) {
	value := r.Context().Value(UserIDCtxKey)
	if value == nil {
		return userID, false
	}

	userID, ok = value.(string)
	if !ok {
		return userID, false
	}

	return userID, true
}

func (h *Handler) resourceID(r *http.Request) (resID string, ok bool) {
	value := r.Context().Value(ResIDCtxKey)
	if value == nil {
		return resID, false
	}

	id, ok := value.(uuid.UUID)
	if !ok {
		return resID, false
	}

	return id.String(), true
}

func (h *Handler) assetReq(r *http.Request) (req AssetRequest, ok bool) {
	value := r.Context().Value(AssetReqCtxKey)
	if value == nil {
		return req, false
	}

	req, ok = value.(AssetRequest)
	if !ok {
		return req, false
	}

	return req, true
}
