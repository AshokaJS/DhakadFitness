package routes

import (
	"net/http"

	"github.com/AshokaJS/DhakadFitness/internal/auth"
	"github.com/AshokaJS/DhakadFitness/internal/user"
)

func SetupRoutes(authService auth.AuthService, userService user.UserService) {
	http.HandleFunc(" /auth/signup", func(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/user/update", user.ProfileUpdateHandler)
	http.HandleFunc("/user/wallet", user.WalletBalanceHandler)
	http.HandleFunc("/user/gyms?like=gold", user.GymSearchHandler)
	http.HandleFunc("user/membership", user.ActiveMembershipHandler)
	http.HandleFunc("/memberships/buy", user.PurchaseMembershipHandler)
	http.HandleFunc("/memberships/user", user.ViewAllMembershipHandler)
	http.HandleFunc("/memberships/scheduled", user.CancelMembershipHandler)
	http.HandleFunc("/user/gyms", user.AccessibleGymHandler)

	// //gym ke endpoints
	// http.HandleFunc("gym/:gymID", gym.ProfileGymHandler)
	// http.HandleFunc("gym/create", gym.ProfileCreateHandler)
	// http.HandleFunc("gym/update/:gymID", gym.ProfileUpdadteHandler)
	// http.HandleFunc("/gym/:gym_id/plans", gym.GymPlansHandler)
	// http.HandleFunc("gym/:gymID/plans/:planID", gym.DeletePlanHandler)
	// http.HandleFunc("gym/:gymID/plans/:planID", gym.UpdatePlanHandler)

	// Health Check Route
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Gym API is running"))
	})
}
