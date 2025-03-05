package gym

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestGymServiceImpl_GetGymProfile(t *testing.T) {
	type fields struct {
		Repo GymRepository
	}
	type args struct {
		gymId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		given   func(mockRep *MockGymRepository)
		want    *[]GetGym
		wantErr bool
	}{
		{
			name:   "test GetGymProfile",
			fields: fields{Repo: &MockGymRepository{}},
			given: func(mockRep *MockGymRepository) {
				mockRep.EXPECT().GetGymProfile(mock.Anything).Return(&[]GetGym{
					{Id: 1, Owner: "kumar", Name: "Gold", Branch_id: 11, Location_Pincode: 411015},
					{Id: 1, Owner: "kumar", Name: "Gold", Branch_id: 12, Location_Pincode: 411016},
				}, nil)
			},
			args: args{gymId: 1},
			want: &[]GetGym{
				{Id: 1, Owner: "kumar", Name: "Gold", Branch_id: 11, Location_Pincode: 411015},
				{Id: 1, Owner: "kumar", Name: "Gold", Branch_id: 12, Location_Pincode: 411016},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockGymRepository{}
			tt.given(mockRepo)
			s := &GymServiceImpl{
				Repo: mockRepo,
			}
			got, err := s.GetGymProfile(tt.args.gymId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GymServiceImpl.GetGymProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GymServiceImpl.GetGymProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
