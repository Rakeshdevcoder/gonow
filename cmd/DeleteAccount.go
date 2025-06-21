// cmd/DeleteAccount.go
package main

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	var req struct {
		ID string `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return writeJSON(w, http.StatusBadRequest, map[string]any{
			"message":    "Invalid JSON format",
			"statusCode": http.StatusBadRequest,
			"error":      err.Error(),
		})
	}

	if req.ID == "" {
		return writeJSON(w, http.StatusBadRequest, map[string]any{
			"message":    "Account id is required",
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

	if err := s.db.DeleteAccount(req.ID); err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, map[string]any{
		"message":        "Account deleted successfully",
		"statusCode":     http.StatusOK,
		"deletedAccount": account,
	})
}
