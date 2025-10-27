package users

import (
	"bookShop/repo/users"
	"bookShop/util"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	ID       int    `json:"id" db:"id"`
	UserName string `json:"user_name" db:"user_name" validate:"required,min=4,max=50"`
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required"`
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser Users

	// Decode JSON request
	err1 := json.NewDecoder(r.Body).Decode(&newUser)
	if err1 != nil {
		util.SendError(w, "Please give a valid JSON", http.StatusBadRequest)
		return
	}

	// Validate input
	validate := validator.New()
	err2 := validate.Struct(newUser)
	if err2 != nil {
		util.SendError(w, err2.Error(), http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err3 := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err3 != nil {
		util.SendError(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	newUser.Password = string(hashedPassword)

	createdUser, err4 := h.users.CreateUser(users.Users{
		UserName: newUser.UserName,
		Email:    newUser.Email,
		Password: newUser.Password,
	})

	if err4 != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createdUser, http.StatusCreated)
}
