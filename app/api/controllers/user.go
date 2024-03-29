package controllers

import (
	"chat/app"
	"chat/app/cassandra"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(cassandra.User)
	user.Password = ""
	app.Respond(w, r, user)
	return
}

func Contacts(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(cassandra.User)

	contacts := cassandra.FindAllContacts(user.Username)
	app.Respond(w, r, contacts)
	return
}
