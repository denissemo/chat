package routes

import (
	"chat/app/api/controllers"
	"github.com/gorilla/mux"
)

func AuthRoutes(router *mux.Router) {
	router.HandleFunc("/auth/login", controllers.Login).Methods("POST")
	router.HandleFunc("/auth/sign-up", controllers.SignUp).Methods("POST")
}
