package user

import "errors"

type UserService interface {
	GetUserProfile(id int) (*User, error)
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
