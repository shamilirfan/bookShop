package cart

import "fmt"

func (r *cartRepo) Delete(cartID int) error {
	query := `DELETE FROM cart WHERE id = $1`
	_, err := r.dbCon.Exec(query, cartID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
