package book

import (
	"log"

	"github.com/lib/pq"
)

func (r *bookRepo) GetByID(bookID int) (*Book, error) {
	var book Book
	query := `SELECT id, title, author, price, description, image_path, category, brand, is_stock
		FROM books
		WHERE id = $1
	`
	rows, err := r.dbCon.Query(query, bookID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Slice for results
	for rows.Next() {
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Price,
			&book.Description,
			pq.Array(&book.ImagePath), // 👈 image array handle করার জন্য pq.Array()
			&book.Category,
			&book.Brand,
			&book.IsStock,
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Check for any errors from iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &book, nil
}
