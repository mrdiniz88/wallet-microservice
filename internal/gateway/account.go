package gateway

import "github.com/mrdiniz88/wallet-microservice/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
