package users

import (
	"bookShop/util"
	"net/http"
)

func (h *Handler) Reset_Password(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	password := r.FormValue("password")

	err := h.users.ResetPassword(token, password)
	if err != nil {
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("✅ Password reset successful"))
}
