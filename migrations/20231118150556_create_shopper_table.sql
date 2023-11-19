-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS shopper (
id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
first_name VARCHAR(45) NULL,
last_name VARCHAR(45) NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shopper;
-- +goose StatementEnd
