package gym

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/AshokaJS/DhakadFitness/utils"
)

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
	// Scheduled_Start_Date time.Time `json:"start_date"`  // required only when purchase
}

func GymProfileHandler(w http.ResponseWriter, r *http.Request, gymService GymService) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
	_, err := utils.AuthentionUtil(w, r)
	if err != nil {
		http.Error(w, "invalid token", http.StatusBadRequest)
		return
	}

	parts := strings.Split(r.URL.Path, "/")

	gymId, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		http.Error(w, "invalid gym id ", http.StatusBadRequest)
		return
	}

	gym, err := gymService.GetGymProfile(gymId)
	if err != nil {
		http.Error(w, "gym not found", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(gym)
}

func GymProfileCreateHandler(w http.ResponseWriter, r *http.Request, gymService GymService) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
	_, err := utils.AuthentionUtil(w, r)
	if err != nil {
		http.Error(w, "invalid token", http.StatusBadRequest)
		return
	}

	role, err := utils.RoleAuthentication(w, r)
	if err != nil {
		http.Error(w, "unable to fetch user role", http.StatusBadRequest)
		return
	}

	if !role {
		http.Error(w, "user is not GymOwner", http.StatusBadRequest)
		return
	}

	var gym GymStruct

	err = json.NewDecoder(r.Body).Decode(&gym)
	if err != nil {
		http.Error(w, "unable to fetch gym details from request body", http.StatusBadRequest)
		return
	}

	ok, err := gymService.CreateGym(&gym)
	if err != nil {
		fmt.Fprint(w, err)
		http.Error(w, "unbale to create gym", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, ok)
	// w.WriteHeader(http.StatusOK)
}

func GymPlansHandler(w http.ResponseWriter, r *http.Request, gymService GymService) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
	_, err := utils.AuthentionUtil(w, r)
	if err != nil {
		http.Error(w, "invalid token", http.StatusBadRequest)
		return
	}

	role, err := utils.RoleAuthentication(w, r)
	if err != nil {
		http.Error(w, "unable to fetch user role", http.StatusBadRequest)
		return
	}

	if !role {
		http.Error(w, "user is not GymOwner", http.StatusBadRequest)
		return
	}

	var plan Plan
	err = json.NewDecoder(r.Body).Decode(&plan)
	if err != nil {
		http.Error(w, "unable to decode request body", http.StatusBadRequest)
	}

	ok, err := gymService.CreatePlan(plan)
	if err != nil {
		fmt.Fprint(w, err)
		http.Error(w, "unbale to add gym plan", http.StatusBadRequest)
	}

	fmt.Fprint(w, ok)
	// w.WriteHeader(http.StatusOK)
}

func DeletePlanHandler(w http.ResponseWriter, r *http.Request, gymService GymService) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	_, err := utils.AuthentionUtil(w, r)
	if err != nil {
		http.Error(w, "invalid token", http.StatusBadRequest)
		return
	}

	role, err := utils.RoleAuthentication(w, r)
	if err != nil {
		http.Error(w, "unable to fetch user role", http.StatusBadRequest)
		return
	}

	if !role {
		http.Error(w, "user is not GymOwner", http.StatusBadRequest)
		return
	}

	parts := strings.Split(r.URL.Path, "/")

	planId, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		http.Error(w, "invalid plan id", http.StatusBadRequest)
	}

	err = gymService.DeletePlan(planId)
	if err != nil {
		http.Error(w, "failed to delete plan", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "gym plan deleted sucessfully")
}
