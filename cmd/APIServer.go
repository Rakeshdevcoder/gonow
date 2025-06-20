// cmd/APIServer.go
package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/rakeshrathoddev/gobank/db"
)

type APIServer struct {
	listenAddr string
	db         *db.Database
}

func NewAPIServer(listenAddress string, database *db.Database) *APIServer {
	return &APIServer{
		listenAddr: listenAddress,
		db:         database,
	}
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	switch r.Method {
	case "GET":
		if err := s.handleGetAccount(w); err != nil {
			log.Printf("GET /account error")
			return s.handleError(w, err, "Error to get account")
		}
		return nil
	case "POST":
		if err := s.handleCreateAccount(w, r); err != nil {
			log.Printf("POST /account error")
			return s.handleError(w, err, "Error to create account")
		}
		return nil
	case "PUT":
		if err := s.handleUpdateAccount(w, r); err != nil {
			log.Printf("PUT /account error")
			return s.handleError(w, err, "Error to update account")
		}
		return nil
	case "DELETE":
		if err := s.handleDeleteAccount(w, r); err != nil {
			log.Printf("DELETE /delete error")
			return s.handleError(w, err, "Error to delete account")
		}
		return nil
	default:
		return writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"message":    "Method not Allowed",
			"statusCode": strconv.Itoa(http.StatusMethodNotAllowed),
		})
	}
}

func (s *APIServer) handleError(w http.ResponseWriter, err error, message string) error {
	log.Printf("API Error %s", err)
	return writeJSON(w, http.StatusInternalServerError, map[string]any{
		"message":    message,
		"statusCode": http.StatusInternalServerError,
		"error":      err.Error(),
	})
}
