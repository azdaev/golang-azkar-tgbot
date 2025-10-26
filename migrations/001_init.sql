-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id INTEGER NOT NULL PRIMARY KEY UNIQUE,
    last_morning_index INTEGER DEFAULT 0,
    last_evening_index INTEGER DEFAULT 0,
    created_at DATE DEFAULT (datetime('now', 'localtime'))
);

CREATE TABLE IF NOT EXISTS configs (
    user_id INTEGER NOT NULL UNIQUE,
    arabic TEXT DEFAULT 'true',
    russian TEXT DEFAULT 'true',
    transcription TEXT DEFAULT 'true',
    audio TEXT DEFAULT 'false',
    PRIMARY KEY(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS configs;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
