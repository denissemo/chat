package controllers

import (
	"encoding/json"
	"net/http"

	"chat/app"
	"chat/app/cassandra"
	"chat/app/middleware"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	type loginBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	body := &loginBody{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		app.Respond(w, r, app.ErrorMessage{Code: 400, Message: "InvalidBody"})
		return
	}

	user := cassandra.FindOneUser(body.Username)
	if user == nil {
		app.Respond(w, r, app.ErrorMessage{Code: 403, Message: "InvalidUsername"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		app.Respond(w, r, app.ErrorMessage{Code: 403, Message: "InvalidCredentials"})
		return
	}

	tokenType := middleware.AuthToken{
		Username: user.Username,
	}
	token := tokenType.Sign()

	user.Password = ""
	response := make(map[string]interface{})
	response["user"] = user
	response["token"] = "JWT " + token

	app.Respond(w, r, response)
	return
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
