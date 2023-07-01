-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks
(
    id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    summary      TEXT,
    is_completed BOOL DEFAULT false,
    created_at   DATETIME,
    updated_at   DATETIME,
    deleted_at   DATETIME NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;
-- +goose StatementEnd
