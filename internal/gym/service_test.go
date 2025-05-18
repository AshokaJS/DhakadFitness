package gym

import (
	"testing"

	"github.com/AshokaJS/DhakadFitness/internal/gym/mocks"
	"github.com/AshokaJS/DhakadFitness/utils"
	"github.com/stretchr/testify/suite"
)

type GymServiceTestSuite struct {
	suite.Suite
	Repo *mocks.GymRepository
}

func TestGymServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GymServiceTestSuite))

}

func (suite *GymServiceTestSuite) SetupTest() {
	suite.Repo = new(mocks.GymRepository)
}

func (suite *GymServiceTestSuite) TearDownTest() {
	suite.Repo.AssertExpectations(suite.T())
}

func (suite *GymServiceTestSuite) TestCreateGym() {
	suite.Repo.On("CreateGym", &utils.GymStruct{
		Name: "Gym1",
	}).Return("Gym1", nil)

	service := NewGymService(suite.Repo)
	result, err := service.CreateGym(&utils.GymStruct{
		Name: "Gym1",
	})

	suite.Nil(err)
	suite.Equal("Gym1", result)
}