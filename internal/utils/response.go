package utils

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func ResponseWithSuccess(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(SuccessResponse{Message: message, Data: data})
}

func ResposnseWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}
