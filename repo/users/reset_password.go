package users

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (r *usersRepo) ResetPassword(token string, password string) error {
	var userID int
	var expires time.Time

	query := `SELECT user_id, expires_at FROM password_resets WHERE token=$1`
	err := r.dbCon.QueryRow(query, token).Scan(&userID, &expires)
	if err != nil {
		return err
	}

	if time.Now().After(expires) {
		return err
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	query = `UPDATE users SET password=$1 WHERE id=$2`
	_, err = r.dbCon.Exec(query, string(hash), userID)
	if err != nil {
		return err
	}

	query = "DELETE FROM password_resets WHERE token=$1"
	_, err = r.dbCon.Exec(query, token)
	if err != nil {
		return err
	}

	return nil
}
