package user

import (
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetUserbyId(id int) (*User, error)
	UpdateUserProfile(id int, rUser User) (*User, error)
	UserWalletBalance(id int) (*Wallet, error)
	UserActiveMemebrship(id int) (*Membership, *[]Branches, error)
	SearchGymsByPincode(code int) (*[]GetGym, error)
	BuyMembership(userId int, plan *BuyPlan) error
}

type UserRepoImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepoImpl{DB: db}
}

func (r *UserRepoImpl) GetUserbyId(id int) (*User, error) {
	var user User
	err := r.DB.QueryRow("SELECT * FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepoImpl) UpdateUserProfile(id int, rUser User) (*User, error) {
	var user User
	err := r.DB.QueryRow("SELECT * FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if rUser.ID == 0 {
		rUser.ID = user.ID
	}
	if rUser.Name == "" {
		rUser.Name = user.Name
	}
	if rUser.Email == "" {
		rUser.Email = user.Email
	}
	if rUser.Password == "" {
		rUser.Password = user.Password
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New("error hashing new password")
		}
		rUser.Password = string(hashedPassword)
	}
	if rUser.Role == "" {
		rUser.Role = user.Role
	}

	_, err = r.DB.Exec("UPDATE users SET name=$1, email=$2, password=$3, role=$4 WHERE id=$5", rUser.Name, rUser.Email, rUser.Password, rUser.Role, rUser.ID)

	if err != nil {
		return nil, errors.New("failed to update the user")
	}

	err = r.DB.QueryRow("SELECT users.name, users.email, users.password, users.role FROM users WHERE id=$1", rUser.ID).Scan(&rUser.Name, &rUser.Email, &rUser.Password, &rUser.Role)
	if err != nil {
		return nil, errors.New("failed to get details of updated user")
	}
	return &rUser, nil

}

func (r *UserRepoImpl) UserWalletBalance(id int) (*Wallet, error) {
	var w Wallet
	err := r.DB.QueryRow("SELECT * FROM wallet WHERE user_id=$1", id).Scan(&w.UserId, &w.Balance)
	if err != nil {
		return nil, errors.New("unable to fetch user wallet balance")
	}
	return &w, nil
}

func (r *UserRepoImpl) UserActiveMemebrship(id int) (*Membership, *[]Branches, error) {
	var m Membership
	curr_date := time.Now().Unix()
	err := r.DB.QueryRow("SELECT id, user_id, gym_id, plan_id, scheduled_start_date, start_date, end_date FROM memberships WHERE user_id=$1 And start_date<=$2 And end_date>=$3", id, curr_date, curr_date).Scan(&m.Id, &m.User_Id, &m.Gym_Id, &m.Plan_Id, &m.Scheduled_Start_Date, &m.Start_Date, &m.End_Date)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, errors.New("no active membership found")
		}
		return nil, nil, errors.New("unable to fetch user membership")
	}

	err = r.DB.QueryRow("SELECT membership_type FROM gym_plans WHERE id=$1", m.Plan_Id).Scan(&m.Membership_Type)
	if err != nil {
		return nil, nil, errors.New("unable to fetch type of membership")
	}
	err = r.DB.QueryRow("SELECT name FROM gyms WHERE id=$1", m.Gym_Id).Scan(&m.Gym_name)

	if err != nil {
		return nil, nil, errors.New("unable to fetch gym name")
	}

	var b []Branches
	if m.Membership_Type == "Global" {
		branchs, err := r.DB.Query("SELECT branch_id, location_pincode FROM branches WHERE gym_id=$1", m.Gym_Id)
		if err != nil {
			return nil, nil, errors.New("unable to fetch branches")
		}
		for branchs.Next() {
			var branch Branches
			if err = branchs.Scan(&branch.Id, &branch.location_Pincode); err != nil {
				return nil, nil, err
			}
			b = append(b, branch)
		}
		return &m, &b, nil
	} else {
		return &m, nil, nil
	}

}

func (r *UserRepoImpl) SearchGymsByPincode(code int) (*[]GetGym, error) {
	rows, err := r.DB.Query("SELECT gyms.id, gyms.owner, gyms.name, branches.branch_id, branches.location_pincode FROM gyms JOIN branches ON gyms.id=branches.gym_id WHERE branches.location_pincode BETWEEN $1 - 10 AND $1 + 10", code)
	if err != nil {
		return nil, errors.New("unable to fetch gyms")
	}

	var gyms []GetGym

	for rows.Next() {
		var gym GetGym
		if err := rows.Scan(&gym.Id, &gym.Owner, &gym.Name, &gym.Branch_id, &gym.Location_Pincode); err != nil {
			return nil, err
		}
		gyms = append(gyms, gym)
	}
	return &gyms, nil
}

func (r *UserRepoImpl) BuyMembership(userId int, plan *BuyPlan) error {
	var balance int
	err := r.DB.QueryRow("SELECT available_balance FROM wallet WHERE user_id=$1", userId).Scan(&balance)
	if err != nil {
		return errors.New("error in updating users wallet balance")
	}
	if balance < (plan.Price) {
		return errors.New("wallet balance is low")
	}
	balance = balance - (plan.Price)
	_, err = r.DB.Exec("UPDATE wallet SET available_balance = $1 WHERE user_id = $2", balance, userId)

	if err != nil {
		return errors.New("unable to buy membership")
	}
	date := (plan.Scheduled_Start_Date)
	new_date := date.Unix() //it is start date in Unix()
	end_d := date.AddDate(0, 0, 30)
	end_date := end_d.Unix() // it is end date in unix()
	_, err = r.DB.Exec("INSERT INTO memberships (user_id, gym_id,plan_id,scheduled_start_date,start_date,end_date) VALUES ($1, $2, $3, $4, $5, $6)", userId, plan.Gym_id, plan.Id, new_date, new_date, end_date)
	return err
}
