package storage

import (
	"user_exam/storage/postgres"
	"user_exam/storage/repo"

	"github.com/jmoiron/sqlx"
)

// IStorage ...
type IStorage interface {
	User() repo.UserStorageI
}

type Pg struct {
	db       *sqlx.DB
	userRepo repo.UserStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *Pg {
	return &Pg{
		db:       db,
		userRepo: postgres.NewUserRepo(db),
	}
}

func (s Pg) User() repo.UserStorageI {
	return s.userRepo
}
