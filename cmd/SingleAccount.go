// cmd/SingleAccount.go
package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *APIServer) handleSingleAccount(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		return writeJSON(w, http.StatusBadRequest, map[string]any{
			"message":    "Account id parameter is required",
			"statusCode": http.StatusBadRequest,
		})
	}

	account, err := s.db.GetAccountByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return writeJSON(w, http.StatusNotFound, map[string]any{
				"message":    "Account not found",
				"statusCode": http.StatusNotFound,
			})
		}
		return err
	}

	return writeJSON(w, http.StatusOK, map[string]any{
		"message":    "Single Account retrieved successfully",
		"statusCode": http.StatusOK,
		"account":    account,
	})
}
