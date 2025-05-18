package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/AshokaJS/DhakadFitness/internal/auth/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignup(t *testing.T) {
	// create an instance of mock
	// var ctx context.Context
	ctx := context.Background()
	mock1 := &mocks.AuthRepository{}

	mock1.On("GetUserByEmail", mock.Anything, "test@gmail.com").Return(nil, errors.New("user not found")).Once()

	mock1.On("CreateUser", mock.Anything, "TestUser", "test@gmail.com", mock.Anything, "GymUser").Return(nil).Once()
	s := AuthServiceImpl{
		Repo: mock1,
	}
	err := s.Signup(ctx, "TestUser", "test@gmail.com", "test123", "GymUser")

	assert.NoError(t, err)
	mock1.AssertExpectations(t)
}

func TestSignupTableDriven(t *testing.T) {
	ctx := context.Background()
	mockRepo := &mocks.AuthRepository{}

	s := AuthServiceImpl{
		Repo: mockRepo,
	}

	mockRepo.On("GetUserByEmail", mock.Anything, "test@gmail.com").Return(nil, errors.New("user not found")).Once()
	mockRepo.On("GetUserByEmail", mock.Anything, "test1@gmail.com").Return(nil, errors.New("user not found")).Once()
	// mockRepo.On("GetUserByEmail", mock.Anything, "test2gmail.com").Return(nil, errors.New("user not found")).Once()

	mockRepo.On("CreateUser", mock.Anything, "TestUser", "test@gmail.com", mock.Anything, "GymUser").Return(nil).Once()
	mockRepo.On("CreateUser", mock.Anything, "TestUser1", "test1@gmail.com", mock.Anything, "GymUser").Return(nil).Once()
	// mockRepo.On("CreateUser", mock.Anything, "TestUser2", "test2gmail.com", mock.Anything, "GymUser").Return(nil).Once()

	tests := []struct {
		name          string
		username      string
		email         string
		password      string
		role          string
		expectedError error
	}{
		{
			name:          "valid email 1",
			username:      "TestUser",
			email:         "test@gmail.com",
			password:      "test123",
			role:          "GymUser",
			expectedError: nil,
		},
		{
			name:          "valid email 2",
			username:      "TestUser1",
			email:         "test1@gmail.com",
			password:      "test123",
			role:          "GymUser",
			expectedError: nil,
		},
		{
			name:          "invalid email",
			username:      "TestUser2",
			email:         "test2gmail.com",
			password:      "test123",
			role:          "GymUser",
			expectedError: ErrInvalidEmail,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := s.Signup(ctx, tc.username, tc.email, tc.password, tc.role)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
	mockRepo.AssertExpectations(t)
}

func TestAuthenticateTableDriv(t *testing.T) {

}
