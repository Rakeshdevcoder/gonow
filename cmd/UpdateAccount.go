// cmd/UpdateAccount.go
package main

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s *APIServer) handleUpdateAccount(w http.ResponseWriter, r *http.Request) error {
	var req struct {
		ID        string `json:"id"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Balance   int    `json:"balance"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return writeJSON(w, http.StatusBadRequest, map[string]any{
			"message":    "Invalid JSON Format",
			"statusCode": http.StatusBadRequest,
		})
	}

	if req.ID == "" {
		return writeJSON(w, http.StatusBadRequest, map[string]any{
			"message":    "Account ID is required",
			"statusCode": http.StatusBadRequest,
		})
	}

	account, err := s.db.GetAccountByID(req.ID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return writeJSON(w, http.StatusNotFound, map[string]any{
				"message":    "Account not found",
				"statusCode": http.StatusNotFound,
			})
		}
		return err
	}

	if req.Firstname != "" {
		account.Firstname = req.Firstname
	}

	if req.Lastname != "" {
		account.Lastname = req.Lastname
	}

	if req.Balance != 0 {
		account.Balance = req.Balance
	}

	if err := s.db.UpdateAccount(account); err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, map[string]any{
		"message":    "Account Updated Successfully",
		"statusCode": http.StatusOK,
		"account":    account,
	})
}
