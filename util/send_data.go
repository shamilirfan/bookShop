package util

import (
	"encoding/json"
	"net/http"
)

// SendData sends a JSON response with data
func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode) 
	json.NewEncoder(w).Encode(&data)
}
