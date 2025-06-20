// cmd/GetAccount.go
package main

import (
	"net/http"
)

func (s *APIServer) handleGetAccount(w http.ResponseWriter) error {

	accounts, err := s.db.GetAllAccounts()
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, map[string]any{
		"message":    "All accounts retrieved successfully",
		"statusCode": http.StatusOK,
		"count":      len(accounts),
		"accounts":   accounts,
	})
}
