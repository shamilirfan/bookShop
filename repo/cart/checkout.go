package cart

import "log"

func (r *cartRepo) Checkout(checkout CheckoutRequest) error {
	var orderID int
	query :=
		`
	INSERT INTO orders 
	(user_id, road_number, holding_number, area, thana, district, phone_number)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id;
	`
	err := r.dbCon.QueryRow(
		query,
		checkout.UserID,
		checkout.RoadNumber,
		checkout.HoldingNumber,
		checkout.Area,
		checkout.Thana,
		checkout.District,
		checkout.PhoneNumber,
	).Scan(&orderID)

	if err != nil {
		log.Println("Failed to create order")
		return err
	}

	query =
		`
	INSERT INTO order_items 
	(order_id, book_id, quantity, unit_price)
	SELECT 
	$1, 
	c.book_id,
	c.quantity,
	b.price
	FROM cart c
	JOIN books b ON c.book_id = b.id
	WHERE c.user_id = $2;
	`
	_, err = r.dbCon.Exec(query, orderID, checkout.UserID)
	if err != nil {
		log.Println("Failed to move cart")
		return err
	}

	query = "DELETE FROM cart WHERE user_id = $1"
	_, err = r.dbCon.Exec(query, checkout.UserID)
	if err != nil {
		log.Println("Failed to clear cart")
		return err
	}

	return nil
}
