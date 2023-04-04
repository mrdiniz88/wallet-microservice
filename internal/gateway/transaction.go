package gateway

import "github.com/mrdiniz88/wallet-microservice/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
