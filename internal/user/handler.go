package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/AshokaJS/DhakadFitness/utils"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request, userService UserService) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
	userId, err := utils.AuthentionUtil(w, r)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	//fetch user profile
	user, err := userService.GetUserProfile(userId)
	if err != nil {
		http.Error(w, "failed to fetch user profile", http.StatusInternalServerError)
		return
	}

	// Send user profile response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func ProfileUpdateHandler(w http.ResponseWriter, r *http.Request, userService UserService) {
	if r.Method != http.MethodPatch {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
	var rUser User
	err := json.NewDecoder(r.Body).Decode(&rUser)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	userId, err := utils.AuthentionUtil(w, r)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	user, err := userService.UpdateProfile(userId, rUser)
	if err != nil {
		http.Error(w, "failed to update the profile", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "updated profile is \n")
	json.NewEncoder(w).Encode(user)

}

func WalletBalanceHandler(w http.ResponseWriter, r *http.Request, userService UserService) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid mehtod", http.StatusMethodNotAllowed)
		return
	}

	userId, err := utils.AuthentionUtil(w, r)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	wallet, err := userService.GetWalletBalance(userId)
	if err != nil {
		http.Error(w, "failed to fetch wallet balance", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "wallet balance of user is :- \n")
	json.NewEncoder(w).Encode(wallet)
}

func GymSearchHandler(w http.ResponseWriter, r *http.Request, userService UserService) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid mehtod", http.StatusMethodNotAllowed)
		return
	}
	_, err := utils.AuthentionUtil(w, r)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	queryparams := r.URL.Query()
	pincode := strings.TrimSpace(queryparams.Get("pincode"))

	gyms, err := userService.SearchGyms(pincode)
	if err != nil {
		http.Error(w, "unable to fetch gyms", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(gyms)
	w.WriteHeader(http.StatusOK)

}

func ActiveMembershipHandler(w http.ResponseWriter, r *http.Request, userService UserService) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid mehtod", http.StatusMethodNotAllowed)
		return
	}

	userId, err := utils.AuthentionUtil(w, r)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	membership, branches, err := userService.GetActiveMembership(userId)

	if err != nil {
		http.Error(w, "failed to fetch user memebership", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "membership of the user : \n")
	json.NewEncoder(w).Encode(membership)

	// var m = make(map[int]interface{})

	fmt.Fprint(w, "branches if the membership is global : \n")
	json.NewEncoder(w).Encode(branches)
}

func PurchaseMembershipHandler(w http.ResponseWriter, r *http.Request, userService UserService) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid mehtod", http.StatusMethodNotAllowed)
		return
	}

	userId, err := utils.AuthentionUtil(w, r)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	var plan BuyPlan
	err = json.NewDecoder(r.Body).Decode(&plan)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = userService.PurchaseGymPlan(userId, &plan)
	if err != nil {
		http.Error(w, "unable to purchase gym membership", http.StatusOK)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(plan)
}
