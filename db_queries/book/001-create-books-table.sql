CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL UNIQUE,
    author VARCHAR(100),
    price INTEGER NOT NULL CHECK (price >= 0),
    description TEXT,
    image_path TEXT[] DEFAULT '{}',
    category VARCHAR(100) NOT NULL,
    brand VARCHAR(100),
    is_stock BOOLEAN NOT NULL DEFAULT TRUE
);
