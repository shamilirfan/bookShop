package util

import (
	"encoding/json"
	"net/http"
)

// SendError sends a JSON error response
func SendError(w http.ResponseWriter, message interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode) // একবারই
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": message,
	})
}
