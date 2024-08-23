-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR        DEFAULT 'name',
    surname    VARCHAR        DEFAULT 'surname',
    email      VARCHAR UNIQUE,
    phone      VARCHAR DEFAULT '+70000000000',
    hashed_pwd VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE
-- +goose StatementEnd
