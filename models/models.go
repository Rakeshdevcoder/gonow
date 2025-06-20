// cmd/models.go
package models

import "math/rand/v2"

type Account struct {
	ID            int    `json:"id"`
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	AccountNumber int    `json:"accountNumber"`
	Balance       int    `json:"balance"`
}

func NewAccount(firstname, lastname string) *Account {
	return &Account{
		Firstname:     firstname,
		Lastname:      lastname,
		AccountNumber: rand.IntN(1000000),
	}
}
