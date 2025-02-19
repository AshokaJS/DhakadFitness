package user

import (
	"encoding/json"
	"fmt"
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

func ProfileUpdateHandler(w http.ResponseWriter, r *http.Request, userService UserService) {
	if r.Method != http.MethodPatch {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
	var rUser User
	err := json.NewDecoder(r.Body).Decode(&rUser)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := middleware.ValidateToken(tokenStr)
	if err != nil {
		http.Error(w, "invalid token", http.StatusInternalServerError)
		return
	}

	userId := claims.UserId

	user, err := userService.UpdateProfile(userId, rUser)
	if err != nil {
		http.Error(w, "failed to update the profile", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "updated profile is \n")
	json.NewEncoder(w).Encode(user)

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
