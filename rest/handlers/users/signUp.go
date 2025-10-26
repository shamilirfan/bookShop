package users

import (
	"bookShop/util"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Users struct {
	ID       int    `json:"id" db:"id"`
	UserName string `json:"user_name" db:"user_name" validate:"required,min=4,max=50"`
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required"`
}

var users []Users

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user Users

	err1 := json.NewDecoder(r.Body).Decode(&user)
	if err1 != nil {
		util.SendError(w, "Please give a valid JSON", http.StatusBadRequest)
		return
	}

	// Validate input
	validate := validator.New()
	err2 := validate.Struct(user)
	if err2 != nil {
		util.SendError(w, err2.Error(), http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		util.SendError(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	user.ID = len(users) + 1
	users = append(users, user)

	util.SendData(w, user, http.StatusCreated)
}
