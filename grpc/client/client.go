package client

import "go_catalog_service/config"

type ServiceManagerI interface{}

type grpcClients struct{}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {

	return &grpcClients{}, nil
}
