package postgres

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
	"user_exam/config"
	pb "user_exam/genproto/user_exam"
	"user_exam/pkg/db"
	"user_exam/storage/repo"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.UserStorageI
}

func (s *UserRepositoryTestSuite) SetupSuite() {
	pgPool, cleanUp := db.ConnectDBForSuite(config.Load())
	s.Repository = NewUserRepo(pgPool)
	s.CleanUpFunc = cleanUp
}

func (s *UserRepositoryTestSuite) TestUserCRUD() {
	user := pb.CreateUserRequest{
		FirstName: "Test User",
		LastName:  "Suite Test User",
		Email:     "vghjjhvjhv@gmail.com",
		Password:  "bvdhwvcjazc4askjb",
	}

	getUserByIdReq := pb.GetUserByIdRequest{
		UserId: uuid.NewString(),
	}

	createdUser, err := s.Repository.CreateUser(&user)
	s.Suite.NotNil(createdUser)
	s.Suite.NoError(err)
	s.Suite.Equal(user.FirstName, createdUser.FirstName)
	s.Suite.Equal(user.LastName, createdUser.LastName)

	getUser, err := s.Repository.GetUserById(&getUserByIdReq)
	s.Suite.NotNil(getUser)
	s.Suite.NoError(err)
	s.Suite.Equal(user.FirstName, getUser.FirstName)
	s.Suite.Equal(user.LastName, getUser.LastName)

	createdUser.FirstName = "Updated User Name"
	createdUser.LastName = "Updated Last Name"

	updatedUser, err := s.Repository.GetUserById(&getUserByIdReq)
	s.Suite.NotNil(updatedUser)
	s.Suite.NoError(err)

	getUpdatedUser, err := s.Repository.GetUserById(&getUserByIdReq)
	s.Suite.NotNil(getUpdatedUser)
	s.Suite.NoError(err)
	s.Suite.NotEqual(createdUser.FirstName, getUpdatedUser.FirstName)
	s.Suite.NotEqual(createdUser.LastName, getUpdatedUser.LastName)

	allUsers, err := s.Repository.GetAllUser(&pb.GetAllUserRequest{Page: 1, Limit: 10})
	s.Suite.NotNil(allUsers)
	s.Suite.NoError(err)

	deleteUser := s.Repository.DeleteUser(&pb.GetUserByIdRequest{UserId: getUserByIdReq.UserId})
	s.Suite.NotNil(deleteUser)
	s.Suite.NoError(err)
}

func (s *UserRepositoryTestSuite) TearDownSuite() {
	s.CleanUpFunc()
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
