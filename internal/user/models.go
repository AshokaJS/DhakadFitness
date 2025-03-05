package user

import "time"

type GetGym struct {
	Id               int
	Owner            string
	Name             string
	Branch_id        string
	Location_Pincode int
}
type BuyPlan struct {
	Id                   int
	Gym_id               int
	Membership_Type      string
	Duration             string
	Price                int
	Scheduled_Start_Date time.Time
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Role     string
}

type Wallet struct {
	UserId  int
	Balance int
}

type Membership struct {
	Id                   string
	User_Id              int
	Gym_Id               int
	Plan_Id              int
	Scheduled_Start_Date int64
	Start_Date           int64
	End_Date             int64
	Membership_Type      string
	Gym_name             string
}

type Branches struct {
	Id               string
	location_Pincode int
}
