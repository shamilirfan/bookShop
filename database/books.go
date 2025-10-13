package database

// Struct define
type Book struct {
	ID       int     `json:"id"` // It is called tag
	BookName string  `json:"bookName"`
	Author   string  `json:"author"`
	Price    float32 `json:"price"`
	IsStock  bool    `json:"isStock"`
	ImageUrl string  `json:"imageUrl"`
}

// Book list define
var bookList []Book

// BookList function
func BookList() []Book {
	return bookList
}

// Get Book by id
func GetByID(bookID int) *Book {
	// searching book
	for _, value := range bookList {
		if bookID == value.ID {
			return &value
		}
	}
	return nil
}

// Create
func Create(newBook Book) Book {
	newBook.ID = len(bookList) + 1       // Write a new book's ID
	bookList = append(bookList, newBook) // Append new book in a book list
	return newBook
}

// Update
func Update(updatedBook Book) {
	// Searching specific book
	for index, value := range bookList {
		if updatedBook.ID == value.ID {
			bookList[index] = updatedBook
		}
	}
}

// Delete
func Delete(bookID int) {
	// Store unmatched value
	var tempList []Book

	// Searching specific value
	for _, value := range bookList {
		if bookID != value.ID {
			tempList = append(tempList, value)
		}
	}

	bookList = tempList
}

// Init function
func init() {
	// List of books
	books := []Book{
		{
			ID:       1,
			BookName: "The Promise of Heaven",
			Author:   "Dr. David Jeremiah",
			Price:    100,
			IsStock:  true,
			ImageUrl: "https://m.media-amazon.com/images/I/71agofkFeiL._SY466_.jpg",
		},
		{
			ID:       2,
			BookName: "How to Test Negative for Stupid",
			Author:   "John Kennedy",
			Price:    200,
			IsStock:  false,
			ImageUrl: "https://m.media-amazon.com/images/I/71tbImx2YVL._SY466_.jpg",
		},
		{
			ID:       3,
			BookName: "The Biography Of John Neely Kennedy",
			Author:   "Marcus Parker",
			Price:    300,
			IsStock:  false,
			ImageUrl: "https://m.media-amazon.com/images/I/61jC5-3L-XL._SY466_.jpg",
		},
		{
			ID:       4,
			BookName: "Last Rites",
			Author:   "Ozzy Osbourne",
			Price:    400,
			IsStock:  true,
			ImageUrl: "https://m.media-amazon.com/images/I/81L9X2TH++L._SY466_.jpg",
		},
		{
			ID:       5,
			BookName: "One Nation Always Under God",
			Author:   "Tim Scott",
			Price:    500,
			IsStock:  true,
			ImageUrl: "https://m.media-amazon.com/images/I/711h9Ts9CjL._SY466_.jpg",
		},
	}

	// Append book list in slice/list
	bookList = append(bookList, books...)
}
