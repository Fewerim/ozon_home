-- +goose UP
CREATE TABLE IF NOT EXISTS users (
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age  INT NOT NULL,
    city TEXT NOT NULL
);


-- +goose DOWN
DROP TABLE IF EXISTS users;