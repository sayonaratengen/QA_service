package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/sayonaratengen/QA_service/internal/handler/dto"
)

func parseIDFromPath(prefix string, r *http.Request) (int, error) {
	path := strings.TrimPrefix(r.URL.Path, prefix)
	if path == "" {
		return 0, errors.New(MsgEmptyIDInURL)
	}

	parts := strings.SplitN(path, "/", 2)
	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, errors.New(MsgInvalidNumericID)
	}

	return id, nil
}

func writeJSON(w http.ResponseWriter, status int, obj any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		http.Error(w, MsgInternalError, http.StatusInternalServerError)
	}
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, dto.ErrorResponse{Error: msg})
}
