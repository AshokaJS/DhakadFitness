package gym

import "github.com/AshokaJS/DhakadFitness/utils"

type GymService interface {
	GetGymProfile(gymId int) (*[]utils.GetGym, error)
	CreateGym(gym *utils.GymStruct) (string, error)
	CreatePlan(plan utils.Plan) (string, error)
	DeletePlan(planId int) error
}

type ListResp struct {
	GymProfiles *[]utils.GetGym `json:"gym_profiles"`
}

type GymServiceImpl struct {
	Repo GymRepository
}

func NewGymService(repo GymRepository) GymService {
	return &GymServiceImpl{Repo: repo}
}

func (s *GymServiceImpl) GetGymProfile(gymId int) (*[]utils.GetGym, error) {
	resp := ListResp{}
	resp.GymProfiles, _ = s.Repo.GetGymProfile(gymId)
	return resp.GymProfiles, nil
}

func (s *GymServiceImpl) CreateGym(gym *utils.GymStruct) (string, error) {
	return s.Repo.CreateGym(gym)
}

func (s *GymServiceImpl) CreatePlan(plan utils.Plan) (string, error) {
	return s.Repo.AddPlan(plan)
}

func (s *GymServiceImpl) DeletePlan(planId int) error {
	return s.Repo.DeletePlan(planId)
}
