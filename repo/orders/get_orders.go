package orders

import "time"

func (r *orderRepo) Get() (interface{}, error) {
	type Item struct {
		BookID      int     `json:"book_id" db:"book_id"`
		Title       string  `json:"title" db:"title"`
		Author      string  `json:"author" db:"author"`
		Description string  `json:"description" db:"description"`
		ImagePath   string  `json:"image_path" db:"image_path"`
		Category    string  `json:"category" db:"category"`
		IsStock     bool    `json:"is_stock" db:"is_stock"`
		Quantity    int     `json:"quantity" db:"quantity"`
		UnitPrice   float64 `json:"unit_price" db:"unit_price"`
		TotalPrice  float64 `json:"total_price" db:"total_price"`
	}

	type Order struct {
		UserID        int       `json:"user_id" db:"user_id"`
		UserName      string    `json:"user_name" db:"user_name"`
		Email         string    `json:"email" db:"email"`
		PhoneNumber   string    `json:"phone_number" db:"phone_number"`
		RoadNumber    string    `json:"road_number" db:"road_number"`
		HoldingNumber string    `json:"holding_number" db:"holding_number"`
		Area          string    `json:"area" db:"area"`
		Thana         string    `json:"thana" db:"thana"`
		District      string    `json:"district" db:"district"`
		CreatedAt     time.Time `json:"created_at" db:"created_at"`
		Items         []Item    `json:"items"`
	}

	// Marge 2 struct
	type Row struct {
		Order
		Item
	}

	var rows []Row
	query := `
	SELECT
	    u.id AS user_id,
	    u.user_name,
	    u.email,
	    o.phone_number,
	    o.road_number,
	    o.holding_number,
	    o.area,
	    o.thana,
	    o.district,
	    b.id AS book_id,
	    b.title,
	    b.author,
	    b.description,
	    b.image_path,
	    b.category,
	    b.is_stock,
	    oi.quantity,
	    oi.unit_price,
	    oi.total_price,
		o.created_at
	FROM
	    users u
	    JOIN orders o ON o.user_id = u.id
	    JOIN order_items oi ON oi.order_id = o.id
	    JOIN books b ON b.id = oi.book_id
	ORDER BY u.id, o.id;
	`

	err := r.dbCon.Select(&rows, query)
	if err != nil {
		return nil, err
	}

	// Group items by user
	resultMap := make(map[int]*Order)
	for _, row := range rows {
		if _, exists := resultMap[row.UserID]; !exists {
			resultMap[row.UserID] = &Order{
				UserID:        row.UserID,
				UserName:      row.UserName,
				Email:         row.Email,
				PhoneNumber:   row.PhoneNumber,
				RoadNumber:    row.RoadNumber,
				HoldingNumber: row.HoldingNumber,
				Area:          row.Area,
				Thana:         row.Thana,
				District:      row.District,
				CreatedAt:     row.CreatedAt,
				Items:         []Item{},
			}
		}

		resultMap[row.UserID].Items = append(resultMap[row.UserID].Items, Item{
			BookID:      row.BookID,
			Title:       row.Title,
			Author:      row.Author,
			Description: row.Description,
			ImagePath:   row.ImagePath,
			Category:    row.Category,
			IsStock:     row.IsStock,
			Quantity:    row.Quantity,
			UnitPrice:   row.UnitPrice,
			TotalPrice:  row.TotalPrice,
		})
	}

	// Convert map to slice
	var result []Order
	for _, order := range resultMap {
		result = append(result, *order)
	}

	return result, nil
}
