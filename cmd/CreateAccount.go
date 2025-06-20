// cmd/CreateAccount.go
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rakeshrathoddev/gobank/models"
)

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	var req models.Account

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Failed to decode Json")
		return writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message":    "Invalid JSON Format",
			"statusCode": http.StatusBadRequest,
			"error":      err.Error(),
		})
	}

	if req.Firstname == "" {
		return writeJSON(w, http.StatusBadRequest, map[string]any{
			"message":    "firstname is required",
			"statusCode": http.StatusBadRequest,
		})
	}

	if req.Lastname == "" {
		return writeJSON(w, http.StatusBadRequest, map[string]any{
			"message":    "lastname is required",
			"statusCode": http.StatusBadRequest,
		})
	}

	account := models.NewAccount(req.Firstname, req.Lastname)

	if err := s.db.InsertAccount(account); err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, map[string]any{
		"message":    "Account Created Successfully",
		"statusCode": http.StatusCreated,
		"account":    account,
	})
}
