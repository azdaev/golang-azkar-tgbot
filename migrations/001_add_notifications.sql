-- Add new columns to users table
ALTER TABLE users ADD COLUMN is_blocked BOOLEAN DEFAULT FALSE;
ALTER TABLE users ADD COLUMN azkar_request_count INTEGER DEFAULT 0;

-- Add notification settings to configs table
ALTER TABLE configs ADD COLUMN morning_notification BOOLEAN DEFAULT FALSE;
ALTER TABLE configs ADD COLUMN evening_notification BOOLEAN DEFAULT FALSE;
