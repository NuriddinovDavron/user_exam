package storage

import (
	"user_exam/pkg/logger"
	"user_exam/storage/mongodb"
	"user_exam/storage/mongorepo"

	"go.mongodb.org/mongo-driver/mongo"
)

type StorageI interface {
	UserService() mongorepo.UserStorageI
}

type storagePg struct {
	userService mongorepo.UserStorageI
}

func New(collection *mongo.Collection, log logger.Logger) StorageI {
	// return &storagePg{userService: postgres.NewUserRepo(db, log)}
	return &storagePg{userService: mongodb.NewUserRepo(collection, log)}
}

func (s *storagePg) UserService() mongorepo.UserStorageI {
	return s.userService
}