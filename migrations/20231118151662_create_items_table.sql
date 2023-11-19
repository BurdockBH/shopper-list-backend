-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS items (
                                         id INT PRIMARY KEY AUTO_INCREMENT NOT NULL ,
                                         name VARCHAR(45) NULL,
                                         price DOUBLE NULL,
                                         description VARCHAR(45) NULL,
                                         quantity INT NULL
                                         );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
