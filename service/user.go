package service

import (
	"context"
	pb "user_exam/genproto/user_exam"
	l "user_exam/pkg/logger"
	"user_exam/storage"

	grpcClient "user_exam/service/grpc_client"
)

// UserService ...
type UserService struct {
	storage storage.StorageI
	logger  l.Logger
	client  grpcClient.IServiceManager
}

func (u UserService) CreateUser(ctx context.Context, user *pb.CreateUserRequest) (*pb.UserApi, error) {
	userResponse, err := u.storage.UserService().CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return userResponse, nil
}

func (u UserService) GetUserById(ctx context.Context, request *pb.GetUserByIdRequest) (*pb.UserApi, error) {
	user, err := u.storage.UserService().GetUserById(ctx, request)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) GetAllUser(ctx context.Context, request *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error) {
	users, err := u.storage.UserService().GetAllUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u UserService) UpdateUser(ctx context.Context, user *pb.User) (*pb.UserApi, error) {
	userR, err := u.storage.UserService().UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return userR, nil
}

func (u UserService) DeleteUser(ctx context.Context, request *pb.GetUserByIdRequest) (*pb.DeleteUserResponse, error) {
	err := u.storage.UserService().DeleteUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{Message: "successfully deleted"}, nil
}

func (u *UserService) CheckField(ctx context.Context, req *pb.CheckUser) (*pb.CheckRes, error) {
	check, err := u.storage.UserService().CheckField(ctx, req)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	return check, nil
}

func (u *UserService) GetUserByEmail(ctx context.Context, req *pb.EmailRequest) (*pb.UserApi, error) {
	res, err := u.storage.UserService().GetUserByEmail(ctx, req)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	return res, nil
}

func (u *UserService) GetUserByRefreshToken(ctx context.Context, req *pb.UserToken) (*pb.UserApi, error) {
	res, err := u.storage.UserService().GetUserByRefreshToken(ctx, req)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	return res, nil
}

func NewUserService(storage storage.StorageI, log l.Logger, client grpcClient.IServiceManager) *UserService {
	return &UserService{
		storage: storage,
		logger:  log,
		client:  client,
	}
}
