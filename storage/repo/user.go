package repo

import (
	pb "user_exam/genproto/user_exam"
)

// UserStorageI ...
type UserStorageI interface {
	CreateUser(mailReq *pb.CreateUserRequest) (*pb.UserApi, error)
	GetUserById(isUnReq *pb.GetUserByIdRequest) (*pb.UserApi, error)
	GetAllUser(crUsReq *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error)
	UpdateUser(logInReq *pb.User) (*pb.UserApi, error)
	DeleteUser(logInReq *pb.GetUserByIdRequest) error
	CheckField(req *pb.CheckUser) (*pb.CheckRes, error)
	GetUserByEmail(req *pb.EmailRequest) (*pb.UserApi, error)
	GetUserByRefreshToken(req *pb.UserToken) (*pb.UserApi, error)
}
