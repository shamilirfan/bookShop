package book

import (
	"database/sql"
	"log"
)

func (r *bookRepo) Delete(bookID int) error {
	query := `DELETE FROM books WHERE id=$1`
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
