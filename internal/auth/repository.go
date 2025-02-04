package auth

import (
	"database/sql"
	"errors"
	"log"
)

// User struct represents a user
type User struct {
	ID       int
	Email    string
	Password string
	Role     string
}

// AuthRepository defines database operations for authentication
type AuthRepository interface {
	CreateUser(username, email, password, role string) error
	GetUserByEmail(email string) (*User, error)
}

// AuthRepoImpl is the implementation of AuthRepository
type AuthRepoImpl struct {
	DB *sql.DB
}

// NewAuthRepository initializes a new repository instance
func NewAuthRepository(db *sql.DB) AuthRepository {
	return &AuthRepoImpl{DB: db}
}

// CreateUser inserts a new user into the database
func (repo *AuthRepoImpl) CreateUser(username, email, password, role string) error {
	_, err := repo.DB.Exec("INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4)", username, email, password, role)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return errors.New("failed to create user")
	}
	return nil
}

// GetUserByEmail fetches a user by email
func (repo *AuthRepoImpl) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	err := repo.DB.QueryRow("SELECT id, email, password, role FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to fetch user")
	}
	return user, nil
}
