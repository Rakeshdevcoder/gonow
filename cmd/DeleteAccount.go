// cmd/DeleteAccount.go
package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/rakeshrathoddev/gobank/models"
)

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	var req models.Account

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return writeJSON(w, http.StatusBadRequest, map[string]any{
			"message":    "Invalid JSON format", // Fixed typo
			"statusCode": http.StatusBadRequest,
			"error":      err.Error(),
		})
	}

	if req.ID == 0 {
		return writeJSON(w, http.StatusBadRequest, map[string]any{
			"message":    "Account id is required",
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

	if err := s.db.DeleteAccount(req.ID); err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, map[string]any{
		"message":        "Account deleted successfully",
		"statusCode":     http.StatusOK,
		"deletedAccount": account,
	})
}
