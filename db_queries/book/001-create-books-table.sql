CREATE TABLE
    books (
        id SERIAL PRIMARY KEY,
        title VARCHAR(60) NOT NULL UNIQUE,
        author VARCHAR(50) NOT NULL,
        price FLOAT NOT NULL,
        description TEXT NOT NULL,
        image_path TEXT NOT NULL,
        category VARCHAR(50) NOT NULL,
        is_stock BOOLEAN NOT NULL
    );