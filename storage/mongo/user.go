package mongo

import (
	"log"

	"github.com/abdukhashimov/sample/storage/repo"
)

type userRepo struct {
}

func NewUserRepo() repo.UserI {
	return &userRepo{}
}

func (u *userRepo) Create() {
	log.Println("Mongodb Create")
}
