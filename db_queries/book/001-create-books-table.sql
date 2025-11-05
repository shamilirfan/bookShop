CREATE TABLE
    IF NOT EXISTS books (
        id SERIAL PRIMARY KEY,
        title VARCHAR(100) NOT NULL UNIQUE,
        author VARCHAR(50),
        price FLOAT NOT NULL,
        description TEXT,
        image_path TEXT NOT NULL,
        category VARCHAR(50) NOT NULL,
        is_stock BOOLEAN NOT NULL
    );