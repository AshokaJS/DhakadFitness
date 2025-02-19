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
	UpdateUserProfile(id int, rUser User) (*User, error)
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

func (r *UserRepoImpl) UpdateUserProfile(id int, rUser User) (*User, error) {
	var user User
	err := r.DB.QueryRow("SELECT * FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if rUser.ID == 0 {
		rUser.ID = user.ID
	}
	if rUser.Name == "" {
		rUser.Name = user.Name
	}
	if rUser.Email == "" {
		rUser.Role = user.Role
	}
	if rUser.Role == "" {
		rUser.Role = user.Role
	}

	_, err1 := r.DB.Exec("UPDATE users SET name=$1, email=$2, password=$3, role=$4 WHERE id=$5", rUser.Name, rUser.Email, rUser.Password, rUser.Role, rUser.ID)

	if err1 != nil {
		return nil, errors.New("failed to update the user")
	}

	err2 := r.DB.QueryRow("SELECT users.name, users.email, users.password, users.role FROM users WHERE id=$1", rUser.ID).Scan(&rUser.Name, &rUser.Email, &rUser.Password, &rUser.Role)
	if err2 != nil {
		return nil, errors.New("failed to get details of updated user")
	}
	return &rUser, nil

}
