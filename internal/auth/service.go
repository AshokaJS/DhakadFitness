package auth

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Signup(ctx context.Context, username, email, password, role string) error
	Authenticate(ctx context.Context, email, role, password string) (*User, error)
}

type AuthServiceImpl struct {
	Repo AuthRepository
}

// constructor hai which initializes a new service instance
func NewAuthService(repo AuthRepository) AuthService {
	return &AuthServiceImpl{Repo: repo}
}

var ErrInvalidEmail = errors.New("invalid email")
var ErrUsrEmailPresent = errors.New("user with this email is already present in the database")

// Signup registers a new user
func (s *AuthServiceImpl) Signup(ctx context.Context, username, email, password, role string) error {
	if email == "" || password == "" || (role != "GymUser" && role != "GymOwner") {
		return errors.New("invalid input")
	}
	if !strings.Contains(email, "@") {
		return ErrInvalidEmail
	}
	_, err := s.Repo.GetUserByEmail(ctx, email)
	if err == nil {
		return ErrUsrEmailPresent
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error hashing password")
	}

	return s.Repo.CreateUser(ctx, username, email, string(hashedPassword), role)

}

func (s *AuthServiceImpl) Authenticate(ctx context.Context, email, role, password string) (*User, error) {
	user, err := s.Repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err1 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err1 != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
