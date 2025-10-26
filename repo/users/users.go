package users

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type Users struct {
	ID       int    `json:"id" db:"id"`
	UserName string `json:"user_name" db:"user_name" validate:"required,min=4,max=50"`
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required"`
}

type UsersRepo interface {
	SignUp()
	SignIn()
}

type usersRepo struct{ dbCon *sqlx.DB }

func NewBookRepo(dbCon *sqlx.DB) UsersRepo {
	repo := &usersRepo{dbCon: dbCon}

	// Validate struct
	validate := validator.New()
	err := validate.Struct(repo)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed on '%s' rule\n", err.Field(), err.Tag())
		}
		return nil
	}

	fmt.Println("Validation passed!")
	return repo
}
