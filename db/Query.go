// db/Query.go
package db

import (
	"context"
	"time"

	"github.com/rakeshrathoddev/gobank/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *Database) CreateAccountTable() error {
	// MongoDB creates collections automatically, no need to create explicitly
	return nil
}

func (d *Database) InsertAccount(account *models.Account) error {
	collection := d.Database.Collection("accounts")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	account.CreatedAt = time.Now()
	result, err := collection.InsertOne(ctx, account)
	if err != nil {
		return err
	}

	account.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *Database) GetAllAccounts() (map[string]*models.Account, error) {
	collection := d.Database.Collection("accounts")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	accounts := make(map[string]*models.Account)

	for cursor.Next(ctx) {
		var account models.Account
		if err := cursor.Decode(&account); err != nil {
			return nil, err
		}
		accounts[account.ID.Hex()] = &account
	}

	return accounts, nil
}

func (d *Database) GetAccountByID(id string) (*models.Account, error) {
	collection := d.Database.Collection("accounts")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var account models.Account
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	return &account, nil
}

func (d *Database) UpdateAccount(account *models.Account) error {
	collection := d.Database.Collection("accounts")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"firstname": account.Firstname,
			"lastname":  account.Lastname,
			"balance":   account.Balance,
			"updatedAt": time.Now(),
		},
	}

	_, err := collection.UpdateOne(ctx, bson.M{"_id": account.ID}, update)
	return err
}

func (d *Database) DeleteAccount(id string) error {
	collection := d.Database.Collection("accounts")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
