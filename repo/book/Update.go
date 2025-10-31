package book

import (
	"database/sql"
	"log"
	"os"
)

func (r *bookRepo) Update(updatedBook Book) (*Book, error) {
	// Fetch old image path
	var oldImage string
	err := r.dbCon.QueryRow("SELECT image_path FROM books WHERE id=$1", updatedBook.ID).Scan(&oldImage)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error fetching old image for book ID %d: %v", updatedBook.ID, err)
		return nil, err
	}

	// If new image uploaded, delete old image
	if updatedBook.ImagePath != "" && oldImage != "" && oldImage != updatedBook.ImagePath {
		os.Remove(oldImage)
	}

	// If no new image, keep old path
	if updatedBook.ImagePath == "" {
		updatedBook.ImagePath = oldImage
	}

	// Update DB
	query := `
		UPDATE books
		SET title=$2, author=$3, price=$4, description=$5, image_path=$6, category=$7, is_stock=$8
		WHERE id=$1
	`

	result, err := r.dbCon.Exec(
		query,
		updatedBook.ID,
		updatedBook.Title,
		updatedBook.Author,
		updatedBook.Price,
		updatedBook.Description,
		updatedBook.ImagePath,
		updatedBook.Category,
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
