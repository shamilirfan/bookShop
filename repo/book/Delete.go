package book

import (
	"database/sql"
	"log"
	"os"
)

func (r *bookRepo) Delete(bookID int) error {
	// Delete image from server
	var imagePath string 
	query := "SELECT image_path FROM books WHERE id=$1"
	err := r.dbCon.QueryRow(query, bookID).Scan(&imagePath)

	// Error handling
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error deleting old image for book ID %d: %v", bookID, err)
		return err
	}

	// Delete image
	if imagePath != "" {
		os.Remove(imagePath)
	}

	query = `DELETE FROM books WHERE id=$1`
	result, err := r.dbCon.Exec(query, bookID)

	// Error handling
	if err != nil {
		log.Printf("Error deleting book ID %d: %v", bookID, err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
