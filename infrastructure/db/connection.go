package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return "user=postgres password=1234 host=localhost port=5432 dbname=bookShop sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionString()
	dbCon, err := sqlx.Connect("postgres", dbSource)

	if err != nil {
		fmt.Println("❌ Database connection failed:", err)
		return nil, err
	}

	fmt.Println("✅ Connected to Database successfully!")
	return dbCon, nil
}
