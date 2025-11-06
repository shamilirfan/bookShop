package orders

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Order struct {
	ID            int         `json:"id" db:"id"`
	UserID        int         `json:"user_id" db:"user_id"`
	RoadNumber    string      `json:"road_number" db:"road_number"`
	HoldingNumber string      `json:"holding_number" db:"holding_number"`
	Area          string      `json:"area" db:"area"`
	Thana         string      `json:"thana" db:"thana"`
	District      string      `json:"district" db:"district"`
	PhoneNumber   string      `json:"phone_number" db:"phone_number"`
	Status        string      `json:"status" db:"status"`
	CreatedAt     time.Time   `json:"created_at" db:"created_at"`
	Items         []OrderItem `json:"items" db:"items"`
}

type OrderItem struct {
	ID         int     `json:"id" db:"id"`
	OrderID    int     `json:"order_id" db:"order_id"`
	BookID     int     `json:"book_id" db:"book_id"`
	Quantity   int     `json:"quantity" db:"quantity"`
	UnitPrice  float64 `json:"unit_price" db:"unit_price"`
	TotalPrice float64 `json:"total_price" db:"total_price"`
}

type OrderRepo interface {
	GetOrders() ([]*Order, error)
	Create(order Order) (*Order, error)
	Update(newOrder Order) (string, error)
	Cancell(newOrder Order) (*Order, error)
	Delete(bookId int) error
}

type orderRepo struct{ dbCon *sqlx.DB }

func NewOrderRepo(dbCon *sqlx.DB) OrderRepo { return &orderRepo{dbCon: dbCon} }
