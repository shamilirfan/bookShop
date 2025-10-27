package users

import "log"

func (r usersRepo) CreateUser(newUser Users) (*Users, error) {
	query := `
	INSERT INTO users(user_name, email, password)
	VALUES($1, $2, $3)
	RETURNING id
	`
	row := r.dbCon.QueryRow(
		query,
		newUser.UserName,
		newUser.Email,
		newUser.Password,
	)

	err := row.Scan(&newUser.ID)
	if err != nil {
		log.Printf("Error scanning new book ID: %v", err)
		return nil, err
	}

	return &newUser, nil
}
