package handler

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
}

func UserViewHandler(users map[string]User) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("user_id")
		if userId == "" {
			http.Error(rw, "user_id is empty", http.StatusBadRequest)
			return
		}

		user, ok := users[userId]
		if !ok {
			http.Error(rw, "user not found", http.StatusNotFound)
			return
		}

		jsonUser, err := json.Marshal(user)
		if err != nil {
			http.Error(rw, "can't provide a json. internal error",
				http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(jsonUser)
	}
}

func UserCreateHandler(users map[string]User) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(rw, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(rw, "invalid json body", http.StatusBadRequest)
			return
		}

		if user.ID == "" {
			http.Error(rw, "id is required", http.StatusBadRequest)
			return
		}

		if _, ok := users[user.ID]; ok {
			http.Error(rw, "user already exists", http.StatusConflict)
			return
		}

		users[user.ID] = user

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(rw).Encode(user)
	}
}
