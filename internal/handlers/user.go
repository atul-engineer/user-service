package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
			IsActive: false,
		}
		userService := data.NewUserService(db)
		usr, err := userService.CreateUser(r.Context(), &user)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}
		userTask := data.NewUserTask()
		go userTask.ProcessEvent(r.Context())
		userTask.WriteEvent(data.UserCreateEvent{
			Id: usr.Id,
			Email: usr.Email,
			Name: usr.Name,
			IsActive: true,
		})
		close(userTask.Event)

		api.WriteJsonResponse(w, http.StatusCreated, api.UserResponse{
			Id: usr.Id,
			Name: usr.Name,
			Email: usr.Email,
			IsActive: usr.IsActive,
		})
	}
}

func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return 
		}
		userId := r.URL.Query().Get("user_id")
		fmt.Println(userId)
		//if userId{}
		userService := data.NewUserService(db)
		users, _ := userService.GetUsers(r.Context())
		json.NewEncoder(w).Encode(&users)
	}
}