-- +goose Up
-- +goose StatementBegin
-- Add new columns to users table
ALTER TABLE users ADD COLUMN is_blocked BOOLEAN DEFAULT FALSE;
ALTER TABLE users ADD COLUMN azkar_request_count INTEGER DEFAULT 0;

-- Add notification settings to configs table
ALTER TABLE configs ADD COLUMN morning_notification BOOLEAN DEFAULT FALSE;
ALTER TABLE configs ADD COLUMN evening_notification BOOLEAN DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Remove notification settings from configs table
ALTER TABLE configs DROP COLUMN evening_notification;
ALTER TABLE configs DROP COLUMN morning_notification;

-- Remove new columns from users table
ALTER TABLE users DROP COLUMN azkar_request_count;
ALTER TABLE users DROP COLUMN is_blocked;
-- +goose StatementEnd
