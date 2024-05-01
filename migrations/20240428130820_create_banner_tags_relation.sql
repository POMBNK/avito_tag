-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS banner_tags (
    banner_id INT,
    tag_id INT,
    FOREIGN KEY (banner_id) REFERENCES banners(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id),
    PRIMARY KEY (banner_id, tag_id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS banner_tags;