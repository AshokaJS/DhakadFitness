package utils

type GymStruct struct {
	Id               int    `json:"id"`
	Owner            string `json:"owner"`
	Name             string `json:"name"`
	Branch_Id        int    `json:"branch_id"`
	Location_Pincode int    `json:"pincode"`
}

type Plan struct {
	Id              int    `json:"id"`
	Gym_id          int    `json:"gym_id"`
	Membership_Type string `json:"membership_type"`
	Duration        string `json:"duration"`
	Price           int    `json:"price"`
}

type GetGym struct {
	Id               int
	Owner            string
	Name             string
	Branch_id        int
	Location_Pincode int
}
