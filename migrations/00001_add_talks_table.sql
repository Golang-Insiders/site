-- +goose Up
-- +goose StatementBegin
CREATE TABLE talks (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    twitter_username TEXT NOT NULL,
    title TEXT NOT NULL,
    summary TEXT NOT NULL,
    timezone TEXT NOT NULL
);

-- TODO: add indexes

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS talks;

-- +goose StatementEnd
