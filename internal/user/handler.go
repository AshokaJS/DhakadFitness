package user

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/AshokaJS/DhakadFitness/pkg/middleware"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request, userService UserService) {
	// Extract token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := middleware.ValidateToken(tokenStr)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Extract UserId from claims
	userId := claims.UserId

	//fetch user profile
	user, err := userService.GetUserProfile(userId)
	if err != nil {
		http.Error(w, "failed to fetch user profile", http.StatusInternalServerError)
		return
	}

	// Send user profile response
	// fmt.Fprint(w, http.StatusOK, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func ProfileUpdateHandler(w http.ResponseWriter, r *http.Request) {
	//
}

func WalletBalanceHandler(w http.ResponseWriter, r *http.Request) {
	//
}

func GymSearchHandler(w http.ResponseWriter, r *http.Request) {
	//
}

func ActiveMembershipHandler(w http.ResponseWriter, r *http.Request) {
	//
}

func PurchaseMembershipHandler(w http.ResponseWriter, r *http.Request) {
	//
}

func ViewAllMembershipHandler(w http.ResponseWriter, r *http.Request) {
	//
}

func CancelMembershipHandler(w http.ResponseWriter, r *http.Request) {
	//
}

func AccessibleGymHandler(w http.ResponseWriter, r *http.Request) {
	//
}

// func SignupHandler(w http.ResponseWriter, r *http.Request, authService AuthService) {

// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	var req AuthRequest
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 	}

// 	err1 := authService.Signup(req.Username, req.Email, req.Password, req.Role)
// 	if err1 != nil {
// 		fmt.Println(err1)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "user registered successfully"})
// }
