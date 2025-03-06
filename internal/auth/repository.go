package auth

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, username, email, password, role string) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type AuthRepoImpl struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &AuthRepoImpl{DB: db}
}

// CreateUser inserts a new user into the database
func (repo *AuthRepoImpl) CreateUser(ctx context.Context, username, email, password, role string) error {
	_, err := repo.DB.Exec("INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4)", username, email, password, role)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		log.Printf("Failed to create user: %v", err)
		return errors.New("failed to create user")
	}
	var usr_id int
	err = repo.DB.QueryRow("SELECT id FROM users WHERE email=$1", email).Scan(&usr_id)
	if err != nil {
		log.Printf("Failed to fetch user for adding wallet balance: %v", err)
		return errors.New("failed to fetch user for adding wallet balance")
	}
	_, err = repo.DB.Exec("INSERT INTO wallet (user_id, available_balance) VALUES ($1,$2)", usr_id, 5000)
	if err != nil {
		log.Printf("Failed to add wallet balance: %v", err)
		return errors.New("failed to add wallet balance")
	}
	return nil
}

func (repo *AuthRepoImpl) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := repo.DB.QueryRow("SELECT id, name, email, password, role FROM users WHERE email=$1", email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		log.Printf("Failed to get user: %v", err)
		return nil, errors.New("error retrieving user")
	}
	return &user, nil
}
