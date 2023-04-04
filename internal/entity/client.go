package entity

import (
	"errors"
	"net/mail"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	Accounts  []*Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(name string, email string) (*Client, error) {

	timeNow := time.Now()

	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	err := client.Validate()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("Name is required")
	}

	if c.Email == "" {
		return errors.New("Email is required")
	}

	_, err := mail.ParseAddress(c.Email)

	if err != nil {
		return errors.New("Email invalid")
	}

	return nil
}

func (c *Client) Update(name string, email string) error {
	c.Name = name
	c.Email = email

	err := c.Validate()

	if err != nil {
		return err
	}

	c.UpdatedAt = time.Now()

	return nil
}

func (c *Client) AddAccount(account *Account) error {

	if account.Client.ID != c.ID {
		return errors.New("Account does not belong to client")
	}

	c.Accounts = append(c.Accounts, account)

	return nil
}
