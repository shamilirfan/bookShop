package orders

func (r *orderRepo) Create(order Order) (*Order, error) {
	query := `
		INSERT INTO orders (
			user_id, 
			road_number,
			holding_number,
			area, 
			thana,
			district, 
			phone_number,
			status
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, 'Pending')
		RETURNING id, created_at
	`
	err := r.dbCon.QueryRow(query,
		order.UserID,
		order.RoadNumber,
		order.HoldingNumber,
		order.Area,
		order.Thana,
		order.District,
		order.PhoneNumber,
	).Scan(&order.ID, &order.CreatedAt)
	if err != nil {
		return nil, err
	}

	for i := range order.Items {
		var price float64
		err = r.dbCon.QueryRow(`SELECT price FROM books WHERE id = $1`,
			order.Items[i].BookID,
		).Scan(&price)
		if err != nil {
			return nil, err
		}

		order.Items[i].OrderID = order.ID
		order.Items[i].UnitPrice = price
		order.Items[i].TotalPrice = price * float64(order.Items[i].Quantity)

		query = `
			INSERT INTO order_items (order_id, book_id, quantity, unit_price)
			VALUES ($1, $2, $3, $4)
			RETURNING id
		`
		err = r.dbCon.QueryRow(
			query,
			order.ID,
			order.Items[i].BookID,
			order.Items[i].Quantity,
			order.Items[i].UnitPrice,
		).Scan(&order.Items[i].ID)
		if err != nil {
			return nil, err
		}
	}

	order.Status = "Pending"
	return &order, nil
}
