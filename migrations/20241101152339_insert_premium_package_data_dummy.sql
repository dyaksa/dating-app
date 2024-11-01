-- +goose Up
-- +goose StatementBegin
INSERT INTO premium_packages (name, price, description, swipe_limit, verification_label) VALUES ('Bronze', 10000, 'bronze package', false, true)
INSERT INTO premium_packages (name, price, description, swipe_limit, verification_label) VALUES ('Gold', 20000, 'gold package', true, true)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
