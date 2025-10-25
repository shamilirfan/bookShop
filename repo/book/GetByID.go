package book

import (
	"database/sql"
	"log"
)

func (r *bookRepo) GetByID(bookID int) (*Book, error) {
	var book Book
	query := `
		SELECT id, title, author, price, description, image_url, book_category, is_stock
		FROM books
		WHERE id = $1
	`

	err := r.dbCon.Get(&book, query, bookID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error fetching book by ID %d: %v", bookID, err)
		return nil, err
	}

	return &book, nil
}
