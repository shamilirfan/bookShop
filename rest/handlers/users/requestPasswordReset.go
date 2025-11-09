package users

import (
	"bookShop/config"
	"bookShop/repo/users"
	"bookShop/util"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func generateToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
	}

	return hex.EncodeToString(b)
}

func (h *Handler) Request_Password_Reset(w http.ResponseWriter, r *http.Request) {
	var req users.ResetRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.SendError(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	// Generate token & expiry
	token := generateToken()
	expires := time.Now().Add(15 * time.Minute)

	_, err = h.users.RequestPasswordReset(req, token, expires)
	if err != nil {
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}

	// Create the reset link
	port := strconv.Itoa(int(config.GetConfig().HttpPort))
	resetLink := fmt.Sprintf("http://localhost:"+port+"/reset-password-form?token=%s", token)

	// Send email using updated SendEmail (HTML formatted)
	subject := "Reset your BookShop password"
	err = util.SendEmail(req.Email, subject, resetLink)

	if err != nil {
		fmt.Println("Email error:", err)
		util.SendError(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "Password reset link sent to your email", http.StatusOK)
}
