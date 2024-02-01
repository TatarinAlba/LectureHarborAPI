-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE ,
    email VARCHAR(255) NOT NULL UNIQUE ,
    hashed_password VARCHAR(255) NOT NULL UNIQUE,
    registration_timestamp TIME DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
