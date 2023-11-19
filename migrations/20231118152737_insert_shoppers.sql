-- +goose Up
-- +goose StatementBegin
INSERT INTO shopper (first_name, last_name) VALUES ('Edo', 'Cicak');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO shopper (first_name, last_name) VALUES ('John', 'Doe');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO shopper (first_name, last_name) VALUES ('Another', 'Userino');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO shopper (first_name, last_name) VALUES ('Jane', 'Doette');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO shopper (first_name, last_name) VALUES ('Doerino', 'Edin');
-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
