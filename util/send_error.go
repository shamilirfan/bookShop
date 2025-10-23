package util

import (
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, data interface{}, statusCode int) {
	// Write header
	w.WriteHeader(statusCode)

	// Encoding
	encoder := json.NewEncoder(w) // w হচ্ছে কোথায় JSON ডেটা যাবে — সেটা বোঝায়।
	encoder.Encode(&data)         // productList কে JSON format এ convert করে
}
