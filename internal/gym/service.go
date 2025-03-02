package gym

type GymService interface {
	GetGymProfile(gymId int) (*[]GetGym, error)
	CreateGym(gym *GymStruct) (string, error)
	CreatePlan(plan Plan) (string, error)
	DeletePlan(planId int) error
}

type GymServiceImpl struct {
	Repo GymRepository
}

func NewGymService(repo GymRepository) GymService {
	return &GymServiceImpl{Repo: repo}
}

func (s *GymServiceImpl) GetGymProfile(gymId int) (*[]GetGym, error) {
	return s.Repo.GetGymProfile(gymId)
}

func (s *GymServiceImpl) CreateGym(gym *GymStruct) (string, error) {
	return s.Repo.CreateGym(gym)
}

func (s *GymServiceImpl) CreatePlan(plan Plan) (string, error) {
	return s.Repo.AddPlan(plan)
}

func (s *GymServiceImpl) DeletePlan(planId int) error {
	return s.Repo.DeletePlan(planId)
}
