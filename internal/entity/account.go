package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(client *Client) (*Account, error) {
	timeNow := time.Now()

	account := &Account{
		ID:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	err := account.validate()

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (a *Account) Credit(amount float64) {
	a.Balance += amount
	a.UpdatedAt = time.Now()
}

func (a *Account) Debit(amount float64) {
	a.Balance -= amount
	a.UpdatedAt = time.Now()
}

func (a *Account) validate() error {
	if a.Client == nil {
		return errors.New("Client must be filled")
	}

	return nil
}
