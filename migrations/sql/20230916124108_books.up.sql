CREATE TABLE IF NOT EXISTS books (
    id uuid,
    tittle VARCHAR(150) NOT NULL,
    description VARCHAR(225),
    price int NOT NULL DEFAULT 0,
    image VARCHAR(255),
    categories VARCHAR(225),
    keywords VARCHAR(225),
    stock int NOT NULL DEFAULT 0,
    publisher VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);