-- +goose Up
CREATE TABLE items (
    id          BIGSERIAL PRIMARY KEY,
    title       TEXT NOT NULL,
    description TEXT NOT NULL
);

-- +goose Down
DROP TABLE items;
