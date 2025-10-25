package book

import "github.com/jmoiron/sqlx"

type Book struct {
	ID           int     `json:"id" db:"id"`
	Title        string  `json:"title" db:"title"`
	Author       string  `json:"author" db:"author"`
	Price        float32 `json:"price" db:"price"`
	Description  string  `json:"description" db:"description"`
	ImageUrl     string  `json:"image_url" db:"image_url"`
	BookCategory string  `json:"book_category" db:"book_category"`
	IsStock      bool    `json:"is_stock" db:"is_stock"`
}

type BookRepo interface {
	Get() ([]*Book, error)
	GetByID(bookID int) (*Book, error)
	Create(newBook Book) (*Book, error)
	Update(updatedBook Book) (*Book, error)
	Delete(bookID int) error
}

type bookRepo struct{ dbCon *sqlx.DB }

func NewBookRepo(dbCon *sqlx.DB) BookRepo { return &bookRepo{dbCon: dbCon} }
