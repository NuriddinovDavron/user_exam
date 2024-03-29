package grpcClient

import (
	"fmt"
	"user_exam/config"
	pbp "user_exam/genproto/product_exam"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	ProductService() pbp.ProductServiceClient
}

type serviceManager struct {
	cfg            config.Config
	productService pbp.ProductServiceClient
}

func New(cfg config.Config) (IServiceManager, error) {
	// dail to post-service
	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.ProductServiceHost, cfg.ProductServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dail host: %s port : %d", cfg.PostgresHost, cfg.ProductServicePort)
	}

	return &serviceManager{
		cfg:            cfg,
		productService: pbp.NewProductServiceClient(connPost),
	}, nil
}

func (s *serviceManager) ProductService() pbp.ProductServiceClient {
	return s.productService
}
