package createAccount

import (
	"github.com/mrdiniz88/wallet-microservice/internal/entity"
	"github.com/mrdiniz88/wallet-microservice/internal/gateway"
)

type InputCreateAccountDTO struct {
	ClientID string
}

type OutputCreateAccountDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(a gateway.AccountGateway, c gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: a,
		ClientGateway:  c,
	}
}

func (uc *CreateAccountUseCase) Execute(input InputCreateAccountDTO) (*OutputCreateAccountDTO, error) {
	client, err := uc.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	account, err := entity.NewAccount(client)
	if err != nil {
		return nil, err
	}

	err = uc.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	return &OutputCreateAccountDTO{
		ID: account.ID,
	}, nil
}
