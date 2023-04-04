package gateway

import "github.com/mrdiniz88/wallet-microservice/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(cliet *entity.Client) error
}
