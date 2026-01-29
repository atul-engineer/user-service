package api

import (
	"encoding/json"
	"net/http"
)

type UserResponse struct {
	Id       int `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool `json:"is_active"`
}

func WriteJsonResponse(w http.ResponseWriter, statusCode int, ur UserResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(ur)
}