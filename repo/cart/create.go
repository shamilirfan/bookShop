package cart

import (
	"log"
)

func (r *cartRepo) Create(newCart Cart) (*Cart, error) {
	query := `
			INSERT INTO cart (user_id, book_id, quantity)
			VALUES ($1, $2, $3)
			ON CONFLICT (user_id, book_id)
			DO UPDATE SET quantity = cart.quantity + EXCLUDED.quantity;
		`
	_, err := r.dbCon.Exec(query, newCart.UserID, newCart.BookID, newCart.Quantity)
	if err != nil {
		log.Println("DB error: " + err.Error())
		return nil, err
	}

	return &newCart, nil
}
