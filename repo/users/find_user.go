package users

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

func (r *usersRepo) FindUser(email, password string) (*Users, error) {
	var user Users
	query := `
		SELECT id, user_name, email, password
		FROM users
		WHERE email = $1
	`

	// Step 1: শুধু email দিয়ে user fetch কর
	err := r.dbCon.Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // user নেই
		}
		return nil, err
	}

	// Step 2: bcrypt দিয়ে password যাচাই
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// password mismatch
		return nil, err
	}

	return &user, nil
}
