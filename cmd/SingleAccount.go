// cmd/SingleAccount.go
package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *APIServer) handleSingleAccount(w http.ResponseWriter, r *http.Request) error {

	vars := mux.Vars(r)
	idstr := vars["id"]

	if idstr == "" {
		return writeJSON(w, http.StatusBadRequest, map[string]any{
			"message":    "Account id query parameter is required",
			"statusCode": http.StatusBadRequest,
		})
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		return writeJSON(w, http.StatusBadRequest, map[string]any{
			"message":    "Invalid account ID format",
			"statusCode": http.StatusBadRequest,
		})
	}

	account, err := s.db.GetAccountByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return writeJSON(w, http.StatusNotFound, map[string]any{
				"message":    "Account not found",
				"statusCode": http.StatusNotFound,
			})
		}
	}

	return writeJSON(w, http.StatusOK, map[string]any{
		"message":    "Single Account retrieved successfully",
		"statusCode": http.StatusOK,
		"account":    account,
	})
}
