-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_preferences (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    gender_preference VARCHAR(10),
    age_range_min INT CHECK (age_range_min >= 18),
    age_range_max INT CHECK (age_range_max >= age_range_min)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_preferences IF EXISTS
-- +goose StatementEnd
