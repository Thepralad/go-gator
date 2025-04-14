-- +goose Up
CREATE TABLE users (
    id INT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) UNIQUE
);

-- +goose Down
-- DROP TABLE users;
