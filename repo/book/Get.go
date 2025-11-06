package book

import "log"

func (r *bookRepo) Get() ([]*Book, error) {
	var books []*Book
	query := `
	SELECT 
		id, 
		title, 
		author, 
		price, 
		description, 
		image_path, 
		category, 
		is_stock
	FROM books
	`

	err := r.dbCon.Select(&books, query)
	if err != nil {
		log.Printf("Error fetching books: %v", err)
		return nil, err
	}

	return books, nil
}
