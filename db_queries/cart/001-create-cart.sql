CREATE TABLE
    IF NOT EXISTS cart (
        id SERIAL PRIMARY KEY,
        user_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        book_id INT NOT NULL REFERENCES books (id) ON DELETE CASCADE,
        quantity INT NOT NULL CHECK (quantity > 0),
        UNIQUE (user_id, book_id) -- same product একবারই থাকবে
    );