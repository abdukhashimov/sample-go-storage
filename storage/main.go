package storage

import (
	"github.com/abdukhashimov/sample/storage/mongo"
	"github.com/abdukhashimov/sample/storage/repo"
)

type StorageI interface {
	User() repo.UserI
}

type storageObj struct {
	userRepo repo.UserI
}

func New() StorageI {

	return &storageObj{
		userRepo: mongo.NewUserRepo(),
	}
}

func (s *storageObj) User() repo.UserI {
	return s.userRepo
}
