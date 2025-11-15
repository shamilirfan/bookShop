CREATE TABLE
    IF NOT EXISTS order_items (
        id SERIAL PRIMARY KEY,
        order_id INT NOT NULL REFERENCES orders (id) ON DELETE CASCADE,
        book_id INT NOT NULL REFERENCES books (id) ON DELETE CASCADE,
        quantity INT NOT NULL CHECK (quantity > 0),
        unit_price NUMERIC(10, 2) NOT NULL CHECK (unit_price >= 0),
        total_price NUMERIC(10, 2) GENERATED ALWAYS AS (quantity * unit_price) STORED
    );