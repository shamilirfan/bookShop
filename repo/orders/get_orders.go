package orders

func (r *orderRepo) Get() ([]*Order, error) {
	var orderList []*Order

	query :=
		`
	SELECT 
		id, 
		user_id, 
		road_number,
		holding_number,
		area, 
		thana,
		district, 
		phone_number,
		status, 
		created_at
	FROM orders 
	`
	err := r.dbCon.Select(&orderList, query)
	if err != nil {
		return nil, err
	}

	return orderList, nil
}
