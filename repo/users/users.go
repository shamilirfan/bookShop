package users

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Users struct {
	ID       int    `json:"id" db:"id"`
	UserName string `json:"user_name" db:"user_name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type ResetRequest struct {
	Email string `json:"email" db:"email"`
}

type UsersRepo interface {
	CreateUser(newUser Users) (*Users, error)
	FindUser(email, password string) (*Users, error)
	RequestPasswordReset(req ResetRequest, token string, expires time.Time) (*ResetRequest, error)
	ResetPassword(token string, password string) error
}

type usersRepo struct{ dbCon *sqlx.DB }

func NewUsersRepo(dbCon *sqlx.DB) UsersRepo {
	return &usersRepo{dbCon: dbCon}
}
