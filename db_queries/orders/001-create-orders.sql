CREATE TABLE
    IF NOT EXISTS orders (
        id SERIAL PRIMARY KEY,
        user_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        road_number VARCHAR(50),
        holding_number VARCHAR(50),
        area VARCHAR(100),
        thana VARCHAR(100),
        district VARCHAR(100) NOT NULL,
        phone_number VARCHAR(20) NOT NULL,
        status VARCHAR(20) CHECK (
            status IN (
                'Pending',
                'Processing',
                'Packed',
                'Shipped',
                'Delivered',
                'Cancelled',
                'Returned'
            )
        ) DEFAULT 'Pending',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );