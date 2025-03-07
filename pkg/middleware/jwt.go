package middleware

import (
	"errors"
	"time"

	"github.com/AshokaJS/DhakadFitness/config"
	"github.com/golang-jwt/jwt/v5"
)

// Claims represents JWT claims
type Claims struct {
	UserId    int
	UserEmail string `json:"userEmail"`
	Role      string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT token for a user
func GenerateToken(userId int, userEmail string, role string) (string, error) {
	secret := config.GetEnv("JWT_SECRET")

	claims := Claims{
		UserId:    userId,
		UserEmail: userEmail,
		Role:      role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Second)), // 24 hours expiration
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateToken parses and validates a JWT token
func ValidateToken(tokenStr string) (*Claims, error) {
	secret := config.GetEnv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
