// cmd/UpdateAccount.go
package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/rakeshrathoddev/gobank/models"
)

func (s *APIServer) handleUpdateAccount(w http.ResponseWriter, r *http.Request) error {
	var req models.Account

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return writeJSON(w, http.StatusBadRequest, map[string]any{
			"message":    "Invalid JSON Format",
			"statusCode": http.StatusBadRequest,
		})
	}

	if req.ID == 0 {
		return writeJSON(w, http.StatusBadRequest, map[string]any{
			"message":    "Account ID is required",
			"statusCode": http.StatusBadRequest,
		})
	}

	account, err := s.db.GetAccountByID(req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
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

	return writeJSON(w, http.StatusOK, map[string]any{
		"message":    "Account Updated Successfully",
		"statusCode": http.StatusOK,
		"account":    account,
	})
}
