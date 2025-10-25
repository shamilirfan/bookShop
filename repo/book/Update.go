package book

import (
	"database/sql"
	"log"
)

func (r *bookRepo) Update(updatedBook Book) (*Book, error) {
	query := `
		UPDATE books
		SET title=$1, author=$2, price=$3, description=$4, image_url=$5, book_category=$6, is_stock=$7
		WHERE id=$8
	`

	result, err := r.dbCon.Exec(
		query,
		updatedBook.Title,
		updatedBook.Author,
		updatedBook.Price,
		updatedBook.Description,
		updatedBook.ImageUrl,
		updatedBook.BookCategory,
		updatedBook.IsStock,
		updatedBook.ID,
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
