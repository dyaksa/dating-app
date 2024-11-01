-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE swipes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    swiper_id UUID REFERENCES users(id) ON DELETE CASCADE,
    target_profile_id UUID REFERENCES users(id) ON DELETE CASCADE,
    swipe_type BOOLEAN NOT NULL, -- TRUE for 'like' and FALSE for 'dislike'
    last_swipe_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE swipes IF EXISTS
-- +goose StatementEnd
