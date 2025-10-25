package repo

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID           int     `json:"id" db:"id"`
	Title        string  `json:"title" db:"title"`
	Author       string  `json:"author" db:"author"`
	Price        float32 `json:"price" db:"price"`
	Description  string  `json:"description" db:"description"`
	ImageUrl     string  `json:"imageUrl" db:"image_url"`
	BookCategory string  `json:"bookCategory" db:"book_category"`
	IsStock      bool    `json:"isStock" db:"is_stock"`
}

// BookRepo interface
type BookRepo interface {
	List() ([]*Book, error)
	GetByID(bookID int) (*Book, error)
	Create(newBook Book) (*Book, error)
	Update(updatedBook Book) (*Book, error)
	Delete(bookID int) error
}

// bookRepo implementation
type bookRepo struct {
	dbCon *sqlx.DB
}

func NewBookRepo(dbCon *sqlx.DB) BookRepo {
	return &bookRepo{dbCon: dbCon}
}

// List all books
func (r *bookRepo) List() ([]*Book, error) {
	var books []*Book
	query := `
		SELECT id, title, author, price, description, image_url, book_category, is_stock
		FROM books
	`

	if err := r.dbCon.Select(&books, query); err != nil {
		log.Printf("Error fetching books: %v", err)
		return nil, err
	}

	return books, nil
}

// GetByID fetches a single book by ID
func (r *bookRepo) GetByID(bookID int) (*Book, error) {
	var book Book
	query := `
		SELECT id, title, author, price, description, image_url, book_category, is_stock
		FROM books
		WHERE id = $1
	`

	if err := r.dbCon.Get(&book, query, bookID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error fetching book by ID %d: %v", bookID, err)
		return nil, err
	}

	return &book, nil
}

// Create inserts a new book and returns it with the assigned ID
func (r *bookRepo) Create(newBook Book) (*Book, error) {
	query := `
		INSERT INTO books (title, author, price, description, image_url, book_category, is_stock)
		VALUES (:title, :author, :price, :description, :image_url, :book_category, :is_stock)
		RETURNING id
	`

	rows, err := r.dbCon.NamedQuery(query, newBook)
	if err != nil {
		log.Printf("Error creating book: %v", err)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&newBook.ID); err != nil {
			log.Printf("Error scanning new book ID: %v", err)
			return nil, err
		}
	}

	return &newBook, nil
}

// Update modifies an existing book
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

// Delete removes a book by ID
func (r *bookRepo) Delete(bookID int) error {
	query := `DELETE FROM books WHERE id=$1`

	result, err := r.dbCon.Exec(query, bookID)
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

// func generateInitialBook(r *bookRepo) {
// 	// List of books
// 	books := []*Book{
// 		{
// 			ID:           1,
// 			Title:        "The Promise of Heaven",
// 			Author:       "Dr. David Jeremiah",
// 			Price:        100,
// 			Description:  "Description",
// 			ImageUrl:     "https://m.media-amazon.com/images/I/71agofkFeiL._SY466_.jpg",
// 			BookCategory: "History",
// 			IsStock:      true,
// 		},
// 		{
// 			ID:           2,
// 			Title:        "How to Test Negative for Stupid",
// 			Author:       "John Kennedy",
// 			Price:        200,
// 			Description:  "Description",
// 			ImageUrl:     "https://m.media-amazon.com/images/I/71tbImx2YVL._SY466_.jpg",
// 			BookCategory: "Chomic",
// 			IsStock:      false,
// 		},
// 		{
// 			ID:           3,
// 			Title:        "The Biography Of John Neely Kennedy",
// 			Author:       "Marcus Parker",
// 			Price:        300,
// 			Description:  "Description",
// 			ImageUrl:     "https://m.media-amazon.com/images/I/61jC5-3L-XL._SY466_.jpg",
// 			BookCategory: "Novel",
// 			IsStock:      false,
// 		},
// 		{
// 			ID:           4,
// 			Title:        "Last Rites",
// 			Author:       "Ozzy Osbourne",
// 			Price:        400,
// 			Description:  "Description",
// 			ImageUrl:     "https://m.media-amazon.com/images/I/81L9X2TH++L._SY466_.jpg",
// 			BookCategory: "Novel",
// 			IsStock:      true,
// 		},
// 		{
// 			ID:           5,
// 			Title:        "One Nation Always Under God",
// 			Author:       "Tim Scott",
// 			Price:        500,
// 			Description:  "Description",
// 			ImageUrl:     "https://m.media-amazon.com/images/I/711h9Ts9CjL._SY466_.jpg",
// 			BookCategory: "History",
// 			IsStock:      true,
// 		},
// 	}

// 	// Append book list in slice/list
// 	r.bookList = append(r.bookList, books...)
// }
