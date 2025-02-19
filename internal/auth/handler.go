package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AshokaJS/DhakadFitness/pkg/middleware"
	"github.com/google/uuid"
)

type AuthRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request, authService AuthService) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	_, ok := ctx.Value("request-id").(string)

	if !ok {
		//if value is not found in the request context then we have to create new value.
		rid := uuid.New()
		ctx = context.WithValue(ctx, "request-id", rid)
	}

	var req AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	err1 := authService.Signup(ctx, req.Username, req.Email, req.Password, req.Role)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user registered successfully"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request, authService AuthService) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	ctx := r.Context()
	_, ok := ctx.Value("request-id").(string)

	if !ok {
		rid := uuid.New()
		ctx = context.WithValue(ctx, "request-id", rid)
	}

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := authService.Authenticate(ctx, req.Email, req.Role, req.Password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "user logged in successfully"})
	json.NewEncoder(w).Encode(LoginResponse{Token: token})

}
