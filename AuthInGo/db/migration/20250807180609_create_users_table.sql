-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users (
    id SERIAL NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin 

SELECT 'down SQL query';
DROP TABLE Users;
-- +goose StatementEnd
