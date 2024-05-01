-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS banners (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    feature_id INT UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL,
    content JSONB,
    is_active BOOLEAN DEFAULT true,
    FOREIGN KEY (feature_id) REFERENCES features(id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS banners;