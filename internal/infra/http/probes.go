package http

import (
	"encoding/json"
	"net/http"
	"time"
)

type (
	Probes struct {
		livenessTS  time.Time
		readinessTS time.Time
	}
)

func NewProbes() *Probes {
	return &Probes{
		livenessTS:  time.Now(),
		readinessTS: time.Now(),
	}
}

func (h *Handler) handleLiveness(w http.ResponseWriter, r *http.Request) {
	// WIP: Better strategies can be implemented.
	if time.Since(h.livenessTS) > time.Minute {
		// Liveness probe failed
		response := struct {
			Status string `json:"status"`
		}{
			Status: "failed",
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write(responseJSON)
		return
	}

	// Liveness probe succeeded
	response := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(responseJSON)
}

func (h *Handler) handleReadiness(w http.ResponseWriter, r *http.Request) {
	// WIP: Better strategies can be implemented.
	if time.Since(h.readinessTS) > time.Minute {
		// Readiness probe failed
		response := struct {
			Status string `json:"status"`
		}{
			Status: "failed",
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write(responseJSON)
		return
	}

	// Readiness probe succeeded
	response := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(responseJSON)
}
