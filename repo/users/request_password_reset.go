package users

import "time"

func (r *usersRepo) RequestPasswordReset(req ResetRequest, token string, expires time.Time) (*ResetRequest, error) {
	// Check if user exists
	var userID int
	query := `SELECT id FROM users WHERE email=$1`

	err := r.dbCon.QueryRow(query, req.Email).Scan(&userID)
	if err != nil {
		return nil, err
	}

	// Store token in DB
	query = `
		INSERT INTO password_resets (user_id, token, expires_at)
		VALUES ($1, $2, $3)
	`
	_, err = r.dbCon.Exec(query, userID, token, expires)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
