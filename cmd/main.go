// cmd/main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rakeshrathoddev/gobank/db"
)

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type APIError struct {
	Error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleSingleAccount)).Methods("GET")

	log.Printf("JSON API Server running on %s", s.listenAddr)

	log.Fatal(http.ListenAndServe(s.listenAddr, router))
}

func main() {
	err := godotenv.Load("./internal/.env")
	if err != nil {
		log.Fatalf(".env file not loaded %s", err)
	}

	serviceURI := os.Getenv("MONGODB_URL")

	database, err := db.NewDatabase(serviceURI)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	defer database.Close()

	if err := database.CreateAccountTable(); err != nil {
		log.Fatalf("Failed to create accounts table: %s", err)
	}

	log.Println("Database connection successful")

	server := NewAPIServer(":9091", database)
	server.Run()
}
