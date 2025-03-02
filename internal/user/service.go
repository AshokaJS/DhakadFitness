package user

import (
	"errors"
	"strconv"
)

type UserService interface {
	GetUserProfile(id int) (*User, error)
	UpdateProfile(id int, rUSer User) (*User, error)
	GetWalletBalance(id int) (*Wallet, error)
	GetActiveMembership(id int) (*Membership, *[]Branches, error)
	SearchGyms(pincode string) (*[]GetGym, error)
	PurchaseGymPlan(userId int, plan *BuyPlan) error
}

type UserServiceImpl struct {
	Repo UserRepository
}

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

func (s *UserServiceImpl) GetWalletBalance(id int) (*Wallet, error) {
	wallet, err := s.Repo.UserWalletBalance(id)
	if err != nil {
		return nil, errors.New("unable to fetch wallet balance")
	}
	return wallet, nil
}

func (s *UserServiceImpl) GetActiveMembership(id int) (*Membership, *[]Branches, error) {
	membership, branches, err := s.Repo.UserActiveMemebrship(id)

	if err != nil {
		if errors.Is(err, errors.New("no active membership found")) {
			return nil, nil, errors.New("no active membership found")
		}
		return nil, nil, errors.New("unable to fetch user membership")
	}

	return membership, branches, nil

}

func (s *UserServiceImpl) SearchGyms(pincode string) (*[]GetGym, error) {
	var code int
	if pincode != "" {
		pincodeInt, err := strconv.Atoi(pincode)
		if err != nil {
			return nil, errors.New("unable to convert pincode from string to int")
		}
		code = pincodeInt
	}
	return s.Repo.SearchGymsByPincode(code)
}

func (s *UserServiceImpl) PurchaseGymPlan(userId int, plan *BuyPlan) error {
	return s.Repo.BuyMembership(userId, plan)
}
