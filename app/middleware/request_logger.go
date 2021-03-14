package middleware

import (
	"log"
	"net/http"
)

var RequestLogger = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		method := request.Method
		uri := request.URL.String()
		log.Printf("<-- [%s] %s\n", method, uri)

		next.ServeHTTP(writer, request)
	})
}
