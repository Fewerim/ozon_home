-- +goose UP
CREATE TABLE IF NOT EXISTS orders (
    id      BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users (id),
    item    TEXT NOT NULL,
    price   INT NOT NULL
);


-- +goose DOWN
DROP TABLE IF EXISTS orders;