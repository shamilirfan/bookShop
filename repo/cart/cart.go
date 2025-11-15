package cart

import "github.com/jmoiron/sqlx"

type Cart struct {
	UserID   int `json:"user_id" db:"user_id"`
	BookID   int `json:"book_id" db:"book_id"`
	Quantity int `json:"quantity" db:"quantity"`
}

type CheckoutRequest struct {
	UserID        int    `json:"user_id" db:"user_id"`
	RoadNumber    string `json:"road_number" db:"road_number"`
	HoldingNumber string `json:"holding_number" db:"holding_number"`
	Area          string `json:"area" db:"area"`
	Thana         string `json:"thana" db:"thana"`
	District      string `json:"district" db:"district"`
	PhoneNumber   string `json:"phone_number" db:"phone_number"`
}

type CartRepo interface {
	Get() ([]*Cart, error)
	Create(newCart Cart) (*Cart, error)
	Checkout(checkout CheckoutRequest) error 
	Delete(cartID int) error
}

type cartRepo struct{ dbCon *sqlx.DB }

func NewCartRepo(dbCon *sqlx.DB) CartRepo {
	return &cartRepo{dbCon: dbCon}
}
