package utils

import (
	"encoding/json"
	"net/http"
)

// WriteJSON sends a JSON response with the given status code
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}