package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"user-service/internal/api"
	"user-service/internal/data"
)

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
}


func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return 
		}
		var userRequest UserRequest
		if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
			log.Printf(err.Error())
			http.Error(w, "invalid request", http.StatusBadRequest)
			return 
		}
		user := data.User{
			Name: userRequest.Name,
			Email: userRequest.Email,
			IsActive: true,
		}
		userService := data.NewUserService(db)
		usr, err := userService.CreateUser(r.Context(), &user)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}
		api.WriteJsonResponse(w, http.StatusCreated, api.UserResponse{
			Id: usr.Id,
			Name: usr.Name,
			Email: usr.Email,
			IsActive: usr.IsActive,
		})
	}
}