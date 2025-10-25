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

	err := r.dbCon.Select(&books, query)
	if err != nil {
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

	err := r.dbCon.Get(&book, query, bookID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error fetching book by ID %d: %v", bookID, err)
		return nil, err
	}

	return &book, nil
}

func (r *bookRepo) Create(newBook Book) (*Book, error) {
	query := `
		INSERT INTO books (title, author, price, description, image_url, book_category, is_stock)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	row := r.dbCon.QueryRow(
		query,
		newBook.Title,
		newBook.Author,
		newBook.Price,
		newBook.Description,
		newBook.ImageUrl,
		newBook.BookCategory,
		newBook.IsStock,
	)

	err := row.Scan(&newBook.ID)
	if err != nil {
		log.Printf("Error scanning new book ID: %v", err)
		return nil, err
	}

	return &newBook, nil
}

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
