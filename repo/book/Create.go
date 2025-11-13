package book

import (
	"log"

	"github.com/lib/pq"
)

func (r *bookRepo) Create(newBook Book) (*Book, error) {
	var bookID int
	query := `
		INSERT INTO books 
		(title, author, price, description, category, brand, is_stock)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	row := r.dbCon.QueryRow(
		query,
		newBook.Title,
		newBook.Author,
		newBook.Price,
		newBook.Description,
		newBook.Category,
		newBook.Brand,
		newBook.IsStock,
	)

	err := row.Scan(&bookID)
	if err != nil {
		log.Printf("Error scanning new book ID: %v", err)
		return nil, err
	}

	query = `UPDATE books SET image_path=$1 WHERE id=$2`
	_, err = r.dbCon.Exec(query, pq.Array(newBook.ImagePath), bookID)
	if err != nil {
		log.Printf("Failed to update images %v", err)
		return nil, err
	}

	// Assign ID to the struct before returning
	newBook.ID = bookID
	return &newBook, nil
}
