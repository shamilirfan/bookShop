package users

import (
	"database/sql"
)

func (r *usersRepo) FindUser(email, password string) (*Users, error) {
	var user Users
	query := `
		SELECT id, user_name, email, password
		FROM users
		WHERE email = $1 AND password = $2
	`

	err := r.dbCon.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
