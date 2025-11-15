package cart

import (
	"log"
)

func (r *cartRepo) Get() ([]*Cart, error) {
	var cartList []*Cart

	query := `SELECT user_id, book_id, quantity FROM cart`
	err := r.dbCon.Select(&cartList, query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return cartList, nil
}
