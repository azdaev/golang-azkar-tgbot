-- +goose Up
-- +goose StatementBegin
-- Fix TEXT boolean columns to proper BOOLEAN type
-- SQLite doesn't support ALTER COLUMN TYPE, so we recreate the table

-- Create new configs table with correct types
CREATE TABLE configs_new (
    user_id INTEGER NOT NULL UNIQUE,
    arabic BOOLEAN DEFAULT TRUE,
    russian BOOLEAN DEFAULT TRUE,
    transcription BOOLEAN DEFAULT TRUE,
    audio BOOLEAN DEFAULT FALSE,
    morning_notification BOOLEAN DEFAULT FALSE,
    evening_notification BOOLEAN DEFAULT FALSE,
    PRIMARY KEY(user_id)
);

-- Copy data, converting TEXT 'true'/'false' to BOOLEAN
INSERT INTO configs_new (user_id, arabic, russian, transcription, audio, morning_notification, evening_notification)
SELECT
    user_id,
    CASE WHEN arabic = 'true' THEN 1 ELSE 0 END,
    CASE WHEN russian = 'true' THEN 1 ELSE 0 END,
    CASE WHEN transcription = 'true' THEN 1 ELSE 0 END,
    CASE WHEN audio = 'true' THEN 1 ELSE 0 END,
    morning_notification,
    evening_notification
FROM configs;

-- Replace old table with new one
DROP TABLE configs;
ALTER TABLE configs_new RENAME TO configs;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Revert to TEXT type for compatibility
CREATE TABLE configs_old (
    user_id INTEGER NOT NULL UNIQUE,
    arabic TEXT DEFAULT 'true',
    russian TEXT DEFAULT 'true',
    transcription TEXT DEFAULT 'true',
    audio TEXT DEFAULT 'false',
    morning_notification BOOLEAN DEFAULT FALSE,
    evening_notification BOOLEAN DEFAULT FALSE,
    PRIMARY KEY(user_id)
);

INSERT INTO configs_old (user_id, arabic, russian, transcription, audio, morning_notification, evening_notification)
SELECT
    user_id,
    CASE WHEN arabic = 1 THEN 'true' ELSE 'false' END,
    CASE WHEN russian = 1 THEN 'true' ELSE 'false' END,
    CASE WHEN transcription = 1 THEN 'true' ELSE 'false' END,
    CASE WHEN audio = 1 THEN 'true' ELSE 'false' END,
    morning_notification,
    evening_notification
FROM configs;

DROP TABLE configs;
ALTER TABLE configs_old RENAME TO configs;
-- +goose StatementEnd
