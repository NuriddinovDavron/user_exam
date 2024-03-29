package mongorepo

import (
	"context"
	pb "user_exam/genproto/user_exam"
)

type UserStorageI interface {
	CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserApi, error)
	GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.UserApi, error)
	GetAllUser(ctx context.Context, req *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error)
	UpdateUser(ctx context.Context, req *pb.User) (*pb.UserApi, error)
	DeleteUser(ctx context.Context, req *pb.GetUserByIdRequest) error
	CheckField(ctx context.Context, req *pb.CheckUser) (*pb.CheckRes, error)
	GetUserByEmail(ctx context.Context, req *pb.EmailRequest) (*pb.UserApi, error)
	GetUserByRefreshToken(ctx context.Context, req *pb.UserToken) (*pb.UserApi,error)
}
