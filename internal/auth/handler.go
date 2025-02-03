package auth

import (
	"encoding/json"
	"net/http"
)

type AuthRequest struct {
	Username string `json:"username"`
	Email    string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var req AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user registered successfully"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

}
