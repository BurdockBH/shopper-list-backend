-- +goose Up
-- +goose StatementBegin
INSERT INTO items (name, PRICE, description, quantity) VALUES ("Apple", 2.99, "Red apple", 10);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO items (name, PRICE, description, quantity) VALUES ("Banana", 1.99, "Yellow banana", 10);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO items (name, PRICE, description, quantity) VALUES ("Orange", 3.99, "Orange orange", 10);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO items (name, PRICE, description, quantity) VALUES ("Pear", 4.99, "Green pear", 10);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO items (name, PRICE, description, quantity) VALUES ("Pineapple", 5.99, "Yellow pineapple", 10);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
