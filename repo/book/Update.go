package book

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
)

func (r *bookRepo) Update(updatedBook Book) (*Book, error) {
	query := `
		UPDATE books
		SET title = $2, author = $3, price = $4, description = $5, image_path = $6, 
		category = $7, brand = $8, is_stock = $9
		WHERE id = $1
	`
	result, err := r.dbCon.Exec(
		query,
		updatedBook.ID,
		updatedBook.Title,
		updatedBook.Author,
		updatedBook.Price,
		updatedBook.Description,
		pq.Array(updatedBook.ImagePath),
		updatedBook.Category,
		updatedBook.Brand,
		updatedBook.IsStock,
	)

	if err != nil {
		log.Printf("Error updating book ID %d: %v", updatedBook.ID, err)
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return &updatedBook, nil
}
