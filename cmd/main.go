// cmd/main.go
package main

import (
	"encoding/json"
	"fmt"
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

// Wrapper that converts apiFunc to http.HandlerFunc
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

	dbhost := os.Getenv("DB_HOST")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, dbhost, dbport, dbname)

	database, err := db.NewDatabase(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database:%s", err)
	}
	defer database.Close()

	if err := database.CreateAccountTable(); err != nil {
		log.Fatalf("Failed to create accounts table:%s", err)
	}

	log.Println("Database connection succesful")

	server := NewAPIServer(":9091", database)
	server.Run()
}
