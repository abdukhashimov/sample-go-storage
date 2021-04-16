package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/abdukhashimov/sample/storage/repo"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo() repo.UserI {
	return &userRepo{}
}

func (u *userRepo) Create() {
	log.Println("I am a create function")
}
