-- +goose Up
-- +goose StatementBegin
CREATE TABLE items (
    id SERIAL primar KEY,
    name VARCHAR(255) NOT NULL,
    
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
