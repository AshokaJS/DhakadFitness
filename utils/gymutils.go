package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/AshokaJS/DhakadFitness/pkg/middleware"
)

func RoleAuthentication(w http.ResponseWriter, r *http.Request) (bool, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		err := errors.New("missing authorization header")
		return false, err
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := middleware.ValidateToken(tokenStr)
	if err != nil {
		err := errors.New("invalid token")
		return false, err
	}

	// Extract UserId from claims
	userRole := claims.Role

	if userRole != "GymOwner" {
		return false, nil
	}
	return true, nil
}
