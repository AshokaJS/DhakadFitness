package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/AshokaJS/DhakadFitness/pkg/middleware"
)

func AuthentionUtil(w http.ResponseWriter, r *http.Request) (int, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		err := errors.New("missing authorization header")
		return 0, err
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := middleware.ValidateToken(tokenStr)

	if err != nil {
		err := errors.New("invalid token")
		return 0, err
	}

	// Extract UserId from claims
	userId := claims.UserId

	return userId, nil
}
