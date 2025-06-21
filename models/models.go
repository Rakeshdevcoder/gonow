// models/model.go
package models

import (
	"math/rand/v2"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Firstname     string             `json:"firstname" bson:"firstname"`
	Lastname      string             `json:"lastname" bson:"lastname"`
	AccountNumber int                `json:"accountNumber" bson:"account_number"`
	Balance       int                `json:"balance" bson:"balance"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
}

func NewAccount(firstname, lastname string) *Account {
	return &Account{
		Firstname:     firstname,
		Lastname:      lastname,
		AccountNumber: rand.IntN(1000000),
		Balance:       0,
	}
}
