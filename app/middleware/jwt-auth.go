package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"chat/app"
	"chat/app/cassandra"
	"github.com/dgrijalva/jwt-go"
)

type AuthToken struct {
	jwt.StandardClaims
	Username string
}

func (tk *AuthToken) Sign() string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	signed, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return signed
}

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuthUrls := []string{"/v1/auth/login", "/v1/auth/sign-up", "/v1/ping"}
		path := r.URL.Path

		// Ignore some urls.
		for _, url := range notAuthUrls {
			if url == path {
				next.ServeHTTP(w, r)
				return
			}
		}

		headerToken := strings.Split(r.Header.Get("Authorization"), " ")
		if len(headerToken) != 2 {
			app.Respond(w, r, app.ErrorMessage{Code: 403, Message: "MissedAuthToken"})
			return
		}

		tokenPart := headerToken[1]
		tk := &AuthToken{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			app.Respond(w, r, app.ErrorMessage{Code: 403, Message: "InvalidAuthToken"})
			return
		}

		if !token.Valid {
			app.Respond(w, r, app.ErrorMessage{Code: 401, Message: "Unauthorized"})
			return
		}

		user := cassandra.UserFindOne(tk.Username)
		if user == nil {
			app.Respond(w, r, app.ErrorMessage{Code: 401, Message: "Unauthorized"})
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
