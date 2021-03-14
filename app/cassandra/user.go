package cassandra

import (
	"fmt"
	"log"
	"strings"
	"time"

	"chat/app"
	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	LastSeen  time.Time `json:"last_seen"`
	Password  string    `json:"password"`
}

func (user *User) Validate() (err app.ErrorMessage, ok bool) {
	if user.Username == "" || user.Password == "" || user.FirstName == "" {
		return app.ErrorMessage{Code: 400, Message: "MissedRequiredParams"}, false
	}

	// Save email and username to lowercase.
	user.Username = strings.ToLower(user.Username)

	// Check duplicate username.
	var existedUsername string
	q := "SELECT username FROM users WHERE username = ? LIMIT 1"
	if err := Session.Query(q, user.Username).Consistency(gocql.One).Scan(&existedUsername); err != nil {
		if err != gocql.ErrNotFound {
			log.Print("User validate: ", err)
			return app.ErrorMessage{Code: 409, Message: "UnknownError"}, false
		}
	}

	if existedUsername == user.Username {
		return app.ErrorMessage{Code: 409, Message: "UsernameAlreadyExist"}, false
	}

	return app.ErrorMessage{}, true
}

func FindOneUser(username string) *User {
	user := User{}
	q := "SELECT * FROM users WHERE username = ? LIMIT 1"
	if err := Session.Query(q, username).Consistency(gocql.One).Scan(
		&user.Username, &user.CreatedAt, &user.FirstName, &user.LastName, &user.LastSeen, &user.Password,
	); err != nil {
		log.Print("Find user: ", err)
		return nil
	}

	return &user
}

func (user *User) Create() (ok bool) {
	user.SetPassword(user.Password)

	q := "INSERT INTO users (username, created_at, first_name, last_name, last_seen, password) VALUES (?, ?, ?, ?, ?, ?)"
	if err := Session.Query(
		q, user.Username, time.Now(), user.FirstName, user.LastName, time.Now(), user.Password,
	).Exec(); err != nil {
		log.Print("User saving: ", err)
		return false
	}

	return true
}

func (user *User) SetPassword(password string) {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(passwordHash)
}

func (user *User) ToString() string {
	return fmt.Sprintf("\nUser(\n\tusername: %s,\n\tname: %s,\n\tlast_seen: %s\n);",
		user.Username, user.FirstName+" "+user.LastName, user.LastSeen,
	)
}
