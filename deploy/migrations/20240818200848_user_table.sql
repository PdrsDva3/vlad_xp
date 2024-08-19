-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    username VARCHAR,
    surname VARCHAR,
    email VARCHAR UNIQUE,
    phone VARCHAR UNIQUE,
    hashed_pwd VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE
-- +goose StatementEnd
