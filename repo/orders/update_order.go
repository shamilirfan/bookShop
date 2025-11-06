package orders

func (r *orderRepo) Update(newOrder Order) (string, error) {
	query := `
	UPDATE orders
	SET status = $1
	WHERE id = $2
	`
	_, err := r.dbCon.Exec(query, newOrder.Status, newOrder.ID)
	if err != nil {
		return "", err
	}

	return newOrder.Status, nil
}
