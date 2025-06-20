// db/Query.go
package db

import (
	"github.com/rakeshrathoddev/gobank/models"
)

func (d *Database) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS accounts (
	id SERIAL PRIMARY KEY,
	firstname VARCHAR(255) NOT NULL,
	lastname VARCHAR(255) NOT NULL,
	account_number INTEGER NOT NULL,
	balance INTEGER DEFAULT 0,
	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

	_, err := d.DB.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (d *Database) InsertAccount(account *models.Account) error {
	query := `INSERT INTO accounts (firstname, lastname, account_number, balance) VALUES ($1, $2, $3, $4) RETURNING id`

	err := d.DB.QueryRow(query, account.Firstname, account.Lastname, account.AccountNumber, account.Balance).Scan(&account.ID)

	if err != nil {
		return err
	}

	return nil
}

func (d *Database) GetAllAccounts() (map[int]*models.Account, error) {
	query := `SELECT id, firstname, lastname, account_number, balance FROM accounts`

	rows, err := d.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	accounts := make(map[int]*models.Account)

	for rows.Next() {
		account := &models.Account{}
		err := rows.Scan(&account.ID, &account.Firstname, &account.Lastname, &account.AccountNumber, &account.Balance)

		if err != nil {
			return nil, err
		}

		accounts[account.ID] = account
	}

	return accounts, nil
}

func (d *Database) GetAccountByID(id int) (*models.Account, error) {
	query := `SELECT id, firstname, lastname, account_number, balance FROM accounts WHERE id=$1`

	account := &models.Account{}

	err := d.DB.QueryRow(query, id).Scan(&account.ID, &account.Firstname, &account.Lastname, &account.AccountNumber, &account.Balance)

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (d *Database) UpdateAccount(account *models.Account) error {
	query := `UPDATE accounts SET firstname = $1, lastname = $2, balance = $3 WHERE id = $4`

	_, err := d.DB.Exec(query, account.Firstname, account.Lastname, account.Balance, account.ID)

	if err != nil {
		return err
	}

	return nil
}

func (d *Database) DeleteAccount(id int) error {
	query := `DELETE FROM accounts WHERE id=$1`
	_, err := d.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
