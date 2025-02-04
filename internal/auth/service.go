package auth

import (
	"errors"

	"github.com/AshokaJS/DhakadFitness/pkg/middleware"
	"golang.org/x/crypto/bcrypt"
)

// AuthService defines authentication-related business logic
type AuthService interface {
	Signup(username, email, password, role string) error
	Login(email, password string) (string, error)
}

// AuthServiceImpl is the implementation of AuthService
type AuthServiceImpl struct {
	Repo AuthRepository
}

// NewAuthService initializes a new service instance
func NewAuthService(repo AuthRepository) AuthService {
	return &AuthServiceImpl{Repo: repo}
}

// Signup registers a new user
func (s *AuthServiceImpl) Signup(username, email, password, role string) error {
	if email == "" || password == "" || (role != "GymUser" && role != "GymOwner") {
		return errors.New("invalid input")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error hashing password")
	}

	return s.Repo.CreateUser(username, email, string(hashedPassword), role)
}

// Login authenticates a user and returns a JWT token
func (s *AuthServiceImpl) Login(email, password string) (string, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid email or password")
	}

	return middleware.GenerateToken(user.Email, user.Role)
}
