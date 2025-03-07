package routes

import (
	"net/http"

	"github.com/AshokaJS/DhakadFitness/internal/auth"
	"github.com/AshokaJS/DhakadFitness/internal/gym"
	"github.com/AshokaJS/DhakadFitness/internal/user"
	"github.com/AshokaJS/DhakadFitness/pkg/middleware"
)

func SetupRoutes(authService auth.AuthService, userService user.UserService, gymService gym.GymService) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/signup", func(w http.ResponseWriter, r *http.Request) {
		auth.SignupHandler(w, r, authService)
	})

	mux.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		auth.LoginHandler(w, r, authService)
	})
	// http.HandleFunc("/auth/logout", auth.LogoutHandler)

	// user ke endpoints

	mux.HandleFunc("/user/profile", func(w http.ResponseWriter, r *http.Request) {
		user.ProfileHandler(w, r, userService)
	})
	mux.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		user.ProfileUpdateHandler(w, r, userService)
	})

	mux.HandleFunc("/user/wallet", func(w http.ResponseWriter, r *http.Request) {
		user.WalletBalanceHandler(w, r, userService)
	})

	mux.HandleFunc("/user/gyms", func(w http.ResponseWriter, r *http.Request) {
		user.GymSearchHandler(w, r, userService)
	})
	mux.HandleFunc("/user/membership", func(w http.ResponseWriter, r *http.Request) {
		user.ActiveMembershipHandler(w, r, userService)
	})

	mux.HandleFunc("/user/plan", func(w http.ResponseWriter, r *http.Request) {
		user.PurchaseMembershipHandler(w, r, userService)
	})

	// //gym ke endpoints

	mux.HandleFunc("/gym/id/", func(w http.ResponseWriter, r *http.Request) {
		gym.GymProfileHandler(w, r, gymService)
	})
	mux.HandleFunc("/gym/create", func(w http.ResponseWriter, r *http.Request) {
		gym.GymProfileCreateHandler(w, r, gymService)
	})

	mux.HandleFunc("/gym/addplan", func(w http.ResponseWriter, r *http.Request) {
		gym.GymPlansHandler(w, r, gymService)
	})

	mux.HandleFunc("/gym/plan/", func(w http.ResponseWriter, r *http.Request) {
		gym.DeletePlanHandler(w, r, gymService)
	})

	// Health Check Route
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Gym API is running"))
	})

	// handler := cors.Default().Handler(mux)
	handler := middleware.EnableCORS(mux)
	return handler
}
