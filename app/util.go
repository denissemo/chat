package app

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Respond(w http.ResponseWriter, r *http.Request, data interface{}, params ...int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Add("Content-Type", "application/json")

	statusCode := 200

	if tmpData, ok := data.(ErrorMessage); ok {
		statusCode = tmpData.Code
		w.WriteHeader(statusCode)
	}

	if len(params) != 0 && params[0] != 0 {
		statusCode = params[0]
		w.WriteHeader(statusCode)
	}

	_ = json.NewEncoder(w).Encode(data)

	// Response log
	method := r.Method
	uri := r.URL.String()
	log.Printf("--> [%s] %s %d", method, uri, statusCode)
}

func LoadEnv() {
	mode := os.Getenv("MODE")

	if mode != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Print("WARNING: No .env file found")
		}
	}
}
