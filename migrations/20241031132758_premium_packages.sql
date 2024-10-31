-- +goose Up
-- +goose StatementBegin
CREATE TABLE premium_packages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    description TEXT,
    swipe_limit BOOLEAN DEFAULT FALSE,
    verification_label BOOLEAN DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE premium_packages IF EXISTS
-- +goose StatementEnd
