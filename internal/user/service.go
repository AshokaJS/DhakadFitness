package user

import (
	"errors"
)

type UserService interface {
	GetUserProfile(id int) (*User, error)
	UpdateProfile(id int, rUSer User) (*User, error)
}

type UserServiceImpl struct {
	Repo UserRepository
}

// constructor hai which initializes a new service instance
func NewUserService(repo UserRepository) UserService {
	return &UserServiceImpl{Repo: repo}
}

func (s *UserServiceImpl) GetUserProfile(id int) (*User, error) {
	user, err := s.Repo.GetUserbyId(id)
	if err != nil {
		return nil, errors.New("failed to fetch user profile")
	}
	return user, nil
}

func (s *UserServiceImpl) UpdateProfile(id int, rUser User) (*User, error) {
	user, err := s.Repo.UpdateUserProfile(id, rUser)
	if err != nil {
		return nil, errors.New("failed to update the user profile")
	}
	return user, nil
}
