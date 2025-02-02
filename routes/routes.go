package routes

import (
	"net/http"

	"github.com/AshokaJS/DhakadFitness/internal/auth"
	"github.com/AshokaJS/DhakadFitness/internal/gym"
	"github.com/AshokaJS/DhakadFitness/internal/user"
)

func SetupRoutes() {
	http.HandleFunc("/auth/signup", auth.SignupHandler)
	http.HandleFunc("/auth/login", auth.LoginHandler)
	http.HandleFunc("/auth/logout", auth.LogoutHandler)

	// user ke endpoints
	http.HandleFunc("/user/profile", user.ProfileHandler)
	http.HandleFunc("/user/update", user.ProfileUpdateHandler)
	http.HandleFunc("/user/wallet", user.GetWalletBalanceHandler)
	//yaha gym search wala handler likhna hai yaad se likh dena
	// http.HandleFunc("user/membership",user.MembershipHandler)
	http.HandleFunc("/memberships/buy", user.BuyMembershipHandler)
	http.HandleFunc("/memberships/user", user.GetUserMembershipsHandler)
	http.HandleFunc("/memberships/scheduled", user.GetScheduledMembershipsHandler)
	http.HandleFunc("/user/membership/:membershipID", user.DeleteMembershipHandler)
	http.HandleFunc("/user/gyms", user.AccessibleGymHandler)

	//gym ke endpoints
	http.HandleFunc("gym/:gymID", gym.ProfileGymHandler)
	http.HandleFunc("gym/create", gym.ProfileCreateHandler)
	http.HandleFunc("gym/update/:gymID", gym.ProfileUpdadteHandler)
	http.HandleFunc("/gym/:gym_id/plans", gym.GymPlansHandler)
	http.HandleFunc("gym/:gymID/plans/:planID", gym.DeletePlanHandler)
	http.HandleFunc("gym/:gymID/plans/:planID", gym.UpdatePlanHandler)

	// Health Check Route
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Gym API is running!"))
	})
}
