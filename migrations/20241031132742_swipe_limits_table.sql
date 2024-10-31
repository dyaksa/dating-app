-- +goose Up
-- +goose StatementBegin
CREATE TABLE swipe_limits (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    daily_swipes INT DEFAULT 0,
    last_swipe_date DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE swipe_limits IF EXISTS
-- +goose StatementEnd
