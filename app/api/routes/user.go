package routes

import (
	"chat/app/api/controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/users/contacts", controllers.Contacts).Methods("GET")
	//{username:[0-9A-Za-z_]+}
}
