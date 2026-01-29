package model

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func JSONResponse(w http.ResponseWriter, status int, success bool, msg string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := APIResponse{
		Success: success,
		Message: msg,
		Data:    data,
	}

	json.NewEncoder(w).Encode(resp)
}
