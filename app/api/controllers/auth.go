package controllers

import (
	"encoding/json"
	"net/http"

	"chat/app"
	"chat/app/cassandra"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := &cassandra.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		app.Respond(w, r, app.ErrorMessage{Code: 400, Message: "InvalidBody"})
		return
	}

	if err, ok := user.Validate(); !ok {
		app.Respond(w, r, err)
		return
	}

	if ok := user.Create(); !ok {
		app.Respond(w, r, app.ErrorMessage{Code: 409, Message: "UserCreationError"})
		return
	}
	user.Password = "" // Don`t send password in response.

	app.Respond(w, r, user)
	return
}
