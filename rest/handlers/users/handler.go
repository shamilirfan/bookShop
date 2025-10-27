package users

import user "bookShop/repo/users"

type Handler struct{ users user.UsersRepo }

func NewHandler(users user.UsersRepo) *Handler {
	return &Handler{users: users}
}
