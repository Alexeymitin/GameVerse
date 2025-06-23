package utils

import (
	"encoding/json"
	"net/http"
)

func EncodeJSON(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func DecodeJSON[T any](req *http.Request) (T, error) {
	var payload T
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
