package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type InfoHandler struct{}

func NewInfoHandler() *InfoHandler {
	return &InfoHandler{}
}

func (h *InfoHandler) GetInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)

	err := encoder.Encode(struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}{
		Name:    "playgo",
		Version: "dev",
	})
	if err != nil {
		slog.Default().ErrorContext(r.Context(), "calling http handler GetInfo", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
}
