package gym

import (
	"database/sql"
	"errors"
	"fmt"
)

type GymRepository interface {
	GetGymProfile(gymId int) (*[]GetGym, error)
	CreateGym(gym *GymStruct) (string, error)
	AddPlan(plan Plan) (string, error)
	DeletePlan(planId int) error
}

type GymRepositoryImpl struct {
	DB *sql.DB
}

func NewGymRepository(db *sql.DB) GymRepository {
	return &GymRepositoryImpl{DB: db}
}

func (r *GymRepositoryImpl) GetGymProfile(gymId int) (*[]GetGym, error) {
	rows, err := r.DB.Query("SELECT gyms.id, gyms.owner, gyms.name, branches.branch_id, branches.location_pincode FROM gyms JOIN branches ON gyms.id = branches.gym_id WHERE gyms.id = $1;", gymId)
	var gyms []GetGym
	if err != nil {
		return nil, errors.New("unable to fetch gym")
	}

	for rows.Next() {
		var gym GetGym
		err = rows.Scan(&gym.Id, &gym.Owner, &gym.Name, &gym.Branch_id, &gym.Location_Pincode)
		if err != nil {
			return nil, errors.New("unable to fetch gym details")
		}
		gyms = append(gyms, gym)
	}

	return &gyms, nil

}

func (r *GymRepositoryImpl) CreateGym(gym *GymStruct) (string, error) {

	// var flag bool
	flag := true
	ids, _ := r.DB.Query("SELECT id FROM gyms")
	for ids.Next() {
		var id int
		_ = ids.Scan(&id)
		if id == gym.Id {
			flag = false
		}
	}
	if flag {
		_, err := r.DB.Exec("INSERT INTO gyms (id, owner, name) VALUES ($1, $2, $3)", gym.Id, gym.Owner, gym.Name)
		if err != nil {
			return " ", fmt.Errorf("error occured is : %v", err)
		}
	}

	_, err := r.DB.Exec("INSERT INTO branches (branch_id, gym_id,location_pincode) VALUES ($1, $2, $3)", gym.Branch_Id, gym.Id, gym.Location_Pincode)
	if err != nil {
		return "", fmt.Errorf("error occured is : %v", err)
	}
	ok := "gym created successfully"
	return ok, nil
}

func (r *GymRepositoryImpl) AddPlan(plan Plan) (string, error) {
	_, err := r.DB.Exec("INSERT INTO gym_plans (id, gym_id, membership_type, duration, price) VALUES ($1, $2, $3,$4,$5)", plan.Id, plan.Gym_id, plan.Membership_Type, plan.Duration, plan.Price)
	if err != nil {
		return " ", fmt.Errorf("error occured is : %v", err)
	}
	ok := "plan added sucessfully"
	return ok, nil
}

func (r *GymRepositoryImpl) DeletePlan(planId int) error {
	_, err := r.DB.Exec("DELETE FROM gym_plans WHERE id=$1", planId)
	return err
}
