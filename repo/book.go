package repo

// Struct define
type Book struct {
	ID           int     `json:"id"` // It is called tag
	Title        string  `json:"title"`
	Author       string  `json:"author"`
	Price        float32 `json:"price"`
	Description  string  `json:"description"`
	ImageUrl     string  `json:"imageUrl"`
	BookCatagory string  `json:"bookCatagory"`
	IsStock      bool    `json:"isStock"`
}

// interface - It carries only function signature
type BookRepo interface {
	GetByID(bookID int) (*Book, error)
	Create(newBook Book) (*Book, error)
	Update(updatedBook Book) (*Book, error)
	Delete(bookID int) error
	List() ([]*Book, error)
}

type bookRepo struct {
	bookList []*Book
}

func NewBookRepo() BookRepo {
	repo := &bookRepo{}

	generateInitialBook(repo)
	return repo
}

func (r *bookRepo) GetByID(bookID int) (*Book, error) {
	for _, value := range r.bookList {
		if bookID == value.ID {
			return value, nil
		}
	}
	return nil, nil
}

func (r *bookRepo) Create(newBook Book) (*Book, error) {
	newBook.ID = len(r.bookList) + 1          // Write a new book's ID
	r.bookList = append(r.bookList, &newBook) // Append new book in a book list
	return &newBook, nil
}

func (r *bookRepo) Update(updatedBook Book) (*Book, error) {
	for index, value := range r.bookList {
		if updatedBook.ID == value.ID {
			r.bookList[index] = &updatedBook
		}
	}
	return nil, nil
}

func (r *bookRepo) Delete(bookID int) error {
	// Store unmatched value
	var tempList []*Book

	// Searching specific value
	for _, value := range r.bookList {
		if bookID != value.ID {
			tempList = append(tempList, value)
		}
	}

	r.bookList = tempList

	return nil
}

func (r *bookRepo) List() ([]*Book, error) {
	return r.bookList, nil
}

func generateInitialBook(r *bookRepo) {
	// List of books
	books := []*Book{
		{
			ID:           1,
			Title:        "The Promise of Heaven",
			Author:       "Dr. David Jeremiah",
			Price:        100,
			Description:  "Description",
			ImageUrl:     "https://m.media-amazon.com/images/I/71agofkFeiL._SY466_.jpg",
			BookCatagory: "History",
			IsStock:      true,
		},
		{
			ID:           2,
			Title:        "How to Test Negative for Stupid",
			Author:       "John Kennedy",
			Price:        200,
			Description:  "Description",
			ImageUrl:     "https://m.media-amazon.com/images/I/71tbImx2YVL._SY466_.jpg",
			BookCatagory: "Chomic",
			IsStock:      false,
		},
		{
			ID:           3,
			Title:        "The Biography Of John Neely Kennedy",
			Author:       "Marcus Parker",
			Price:        300,
			Description:  "Description",
			ImageUrl:     "https://m.media-amazon.com/images/I/61jC5-3L-XL._SY466_.jpg",
			BookCatagory: "Novel",
			IsStock:      false,
		},
		{
			ID:           4,
			Title:        "Last Rites",
			Author:       "Ozzy Osbourne",
			Price:        400,
			Description:  "Description",
			ImageUrl:     "https://m.media-amazon.com/images/I/81L9X2TH++L._SY466_.jpg",
			BookCatagory: "Novel",
			IsStock:      true,
		},
		{
			ID:           5,
			Title:        "One Nation Always Under God",
			Author:       "Tim Scott",
			Price:        500,
			Description:  "Description",
			ImageUrl:     "https://m.media-amazon.com/images/I/711h9Ts9CjL._SY466_.jpg",
			BookCatagory: "History",
			IsStock:      true,
		},
	}

	// Append book list in slice/list
	r.bookList = append(r.bookList, books...)
}
