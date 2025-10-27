package users

import (
	"github.com/jmoiron/sqlx"
)

type Users struct {
	ID       int    `json:"id" db:"id"`
	UserName string `json:"user_name" db:"user_name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type UsersRepo interface {
	CreateUser(newUser Users) (*Users, error)
	SignIn()
}

type usersRepo struct{ dbCon *sqlx.DB }

func NewUsersRepo(dbCon *sqlx.DB) UsersRepo {
	repo := &usersRepo{dbCon: dbCon}
	return repo
}
