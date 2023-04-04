package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	account, err := NewAccount(client)

	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account, err := NewAccount(nil)

	assert.Nil(t, account)
	assert.NotNil(t, err)
	assert.Error(t, err, "Client must be filled")

}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	account, _ := NewAccount(client)

	account.Credit(100)

	assert.Equal(t, float64(100), account.Balance)
}

func TestDeditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	account, _ := NewAccount(client)

	account.Credit(100)
	account.Debit(50)

	assert.Equal(t, float64(50), account.Balance)
}
