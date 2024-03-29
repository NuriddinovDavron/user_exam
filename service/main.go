package service

import (
	"context"
	"user_exam/config"
	pb "user_exam/genproto/user_exam"
	"user_exam/pkg/logger"
	grpcClient2 "user_exam/service/grpc_client"
	storage2 "user_exam/storage"
	"fmt"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Service struct {
	UserService *UserService
}

func New(cfg *config.Config, log logger.Logger) (*Service, error) {
	// postgres, err := db.New(*cfg)
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot connect to database:", err.Error())
	// }

	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	collection := client.Database("userdb").Collection("users")
	storage := storage2.New(collection, log)
	grpcClient, err := grpcClient2.New(*cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to grpc client:%v", err.Error())
	}

	return &Service{UserService: NewUserService(storage, log, grpcClient)}, nil
} 

func (s *Service) Run(log logger.Logger, cfg *config.Config) {
	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, s.UserService)

	listen, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("error while creating a listener", logger.Error(err))
		return
	}

	defer logger.Cleanup(log)

	log.Info("main: sqlConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
		logger.String("rpc port", cfg.RPCPort))

	if err := server.Serve(listen); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
