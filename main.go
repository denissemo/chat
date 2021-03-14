package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"chat/app"
	"chat/app/api/routes"
	"chat/app/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func init() {
	app.LoadEnv()
}

func main() {
	var router *mux.Router
	router = mux.NewRouter()
	router = router.PathPrefix("/v1").Subrouter()
	router.Use(middleware.RequestLogger)

	routes.AuthRoutes(router)

	router.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "PONG")
	})

	port := os.Getenv("PORT")
	if port == "" {
		// Set default port
		port = "3000"
	}

	log.Printf("INFO: Server started on http://127.0.0.1:%s", port)

	if err := http.ListenAndServe(":"+port, handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)); err != nil {
		log.Fatal(err)
	}
}
