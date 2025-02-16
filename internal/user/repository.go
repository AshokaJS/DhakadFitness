package user

import (
	"database/sql"
	"errors"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Role     string
}
type UserRepository interface {
	GetUserbyId(id int) (*User, error)
}

type UserRepoImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepoImpl{DB: db}
}

func (r *UserRepoImpl) GetUserbyId(id int) (*User, error) {
	var user User
	err := r.DB.QueryRow("SELECT * FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
