package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/AshokaJS/DhakadFitness/pkg/middleware"
	"github.com/AshokaJS/DhakadFitness/utils"
)


func SignupHandler(w http.ResponseWriter, r *http.Request, authService AuthService) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	ctx := utils.GetContext(r)

	var req AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = authService.Signup(ctx, req.Username, req.Email, req.Password, req.Role)
	if err != nil {
		if errors.Is(err, ErrInvalidEmail) {
			http.Error(w, "enter correct email", http.StatusBadRequest)
			return
		}
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user registered successfully"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request, authService AuthService) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	ctx := utils.GetContext(r)

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
