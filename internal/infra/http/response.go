package http

import (
	"encoding/json"
	"net/http"

	"github.com/jobquestvault/platform-go-challenge/internal/sys/errors"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var (
	MethodNotAllowedErr   = errors.NewError("method not allowed")
	InvalidResourceErr    = errors.NewError("invalid resource")
	NoUserErr             = errors.NewError("no user ID provided")
	NoResourceErr         = errors.NewError("no resource ID provided")
	NoAssetReqErr         = errors.NewError("no asset request provided")
	InvalidRequestErr     = errors.NewError("invalid request")
	InvalidRequestDataErr = errors.NewError("invalid request data")
)

func (h *Handler) handleSuccess(w http.ResponseWriter, payload interface{}, msg ...string) {
	var m string
	if len(msg) > 0 {
		m = msg[0]
	}

	response := APIResponse{
		Success: true,
		Message: m,
		Data:    payload,
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		h.Log().Error(errors.Wrap("error encoding handler success", err))
	}

	return
}

func (h *Handler) handleError(w http.ResponseWriter, handlerError error) {
	response := APIResponse{
		Success: false,
		Message: handlerError.Error(),
	}

	h.Log().Error("handler error:", handlerError)

	w.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		h.Log().Error(errors.Wrap("error encoding handler error", err))
	}

	return
}
