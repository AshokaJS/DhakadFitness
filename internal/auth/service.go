package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Signup(username, email, password, role string) error
	Authenticate(email, role, password string) (*User, error)
}

type AuthServiceImpl struct {
	Repo AuthRepository
}

// constructor hai which initializes a new service instance
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

func (s *AuthServiceImpl) Authenticate(email, role, password string) (*User, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err1 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err1 != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
