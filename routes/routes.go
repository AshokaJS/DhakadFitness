package routes

import (
	"net/http"

	"github.com/AshokaJS/DhakadFitness/internal/auth"
	"github.com/AshokaJS/DhakadFitness/internal/gym"
	"github.com/AshokaJS/DhakadFitness/internal/user"
)

func SetupRoutes(authService auth.AuthService, userService user.UserService, gymService gym.GymService) {
	http.HandleFunc("/auth/signup", func(w http.ResponseWriter, r *http.Request) {
		auth.SignupHandler(w, r, authService)
	})
	http.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		auth.LoginHandler(w, r, authService)
	})
	// http.HandleFunc("/auth/logout", auth.LogoutHandler)

	// user ke endpoints

	http.HandleFunc("/user/profile", func(w http.ResponseWriter, r *http.Request) {
		user.ProfileHandler(w, r, userService)
	})
	http.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		user.ProfileUpdateHandler(w, r, userService)
	})

	http.HandleFunc("/user/wallet", func(w http.ResponseWriter, r *http.Request) {
		user.WalletBalanceHandler(w, r, userService)
	})

	http.HandleFunc("/user/gyms", func(w http.ResponseWriter, r *http.Request) {
		user.GymSearchHandler(w, r, userService)
	})
	http.HandleFunc("/user/membership", func(w http.ResponseWriter, r *http.Request) {
		user.ActiveMembershipHandler(w, r, userService)
	})

	http.HandleFunc("/user/plan", func(w http.ResponseWriter, r *http.Request) {
		user.PurchaseMembershipHandler(w, r, userService)
	})

	// //gym ke endpoints

	http.HandleFunc("/gym/id/", func(w http.ResponseWriter, r *http.Request) {
		gym.GymProfileHandler(w, r, gymService)
	})
	http.HandleFunc("/gym/create", func(w http.ResponseWriter, r *http.Request) {
		gym.GymProfileCreateHandler(w, r, gymService)
	})

	http.HandleFunc("/gym/addplan", func(w http.ResponseWriter, r *http.Request) {
		gym.GymPlansHandler(w, r, gymService)
	})

	http.HandleFunc("/gym/plan/", func(w http.ResponseWriter, r *http.Request) {
		gym.DeletePlanHandler(w, r, gymService)
	})

	// Health Check Route
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Gym API is running"))
	})
}
