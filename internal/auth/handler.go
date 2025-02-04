package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// var AuthService1 AuthService

func SignupHandler(w http.ResponseWriter, r *http.Request, authService AuthService) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var req AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	err1 := authService.Signup(req.Username, req.Email, req.Password, req.Role)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user registered successfully"})
}

// func LoginHandler(w http.ResponseWriter, r *http.Request) {

// }

// func LogoutHandler(w http.ResponseWriter, r *http.Request) {

// }
