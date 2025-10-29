package book

import (
	"log"
)

func (r *bookRepo) Create(newBook Book) (*Book, error) {
	query := `
		INSERT INTO books (title, author, price, description, image_path, category, is_stock)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	row := r.dbCon.QueryRow(
		query,
		newBook.Title,
		newBook.Author,
		newBook.Price,
		newBook.Description,
		newBook.ImagePath,
		newBook.Category,
		newBook.IsStock,
	)

	err := row.Scan(&newBook.ID)
	if err != nil {
		log.Printf("Error scanning new book ID: %v", err)
		return nil, err
	}

	return &newBook, nil
}
