package book

import (
	"log"

	"github.com/lib/pq"
)

func (r *bookRepo) Get() ([]*Book, error) {
	query := `SELECT id, title, author, price, description, image_path, category, brand, is_stock 
	FROM books`

	rows, err := r.dbCon.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Slice for results
	var books []*Book
	for rows.Next() {
		var b Book
		err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Author,
			&b.Price,
			&b.Description,
			pq.Array(&b.ImagePath), // 👈 image array handle করার জন্য pq.Array()
			&b.Category,
			&b.Brand,
			&b.IsStock,
		)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, &b)
	}

	// Check for any errors from iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return books, nil
}
