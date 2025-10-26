package models

type User struct {
	Id                int64  `db:"id"`
	LastEveningIndex  int    `db:"last_evening_index"`
	LastMorningIndex  int    `db:"last_morning_index"`
	IsBlocked         bool   `db:"is_blocked"`
	AzkarRequestCount int    `db:"azkar_request_count"`
	CreatedAt         string `db:"created_at"`
}

type UserConfig struct {
	Arabic              bool `db:"arabic"`
	Russian             bool `db:"russian"`
	Transcription       bool `db:"transcription"`
	Audio               bool `db:"audio"`
	MorningNotification bool `db:"morning_notification"`
	EveningNotification bool `db:"evening_notification"`
}
