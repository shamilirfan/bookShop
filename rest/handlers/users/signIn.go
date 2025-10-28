package users

import (
	"bookShop/config"
	"bookShop/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var user Users

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		util.SendError(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	usr, _ := h.users.FindUser(user.Email, user.Password)

	if usr == nil {
		util.SendError(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	cnf := config.GetConfig()

	accessToken, err := util.CreateJwt(cnf.JwtSecretKey, util.Payload{
		Sub:      usr.ID,
		UserName: usr.UserName,
		Email:    usr.Email,
	})

	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, accessToken, http.StatusOK)
}
