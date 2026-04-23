package user

import (
	"encoding/json"
	"net/http"
)

// Common helpers for user handlers.
func SendJSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func SendError(w http.ResponseWriter, code int, msg string) {
	SendJSON(w, code, map[string]string{"error": msg})
}