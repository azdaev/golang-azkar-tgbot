package models

type User struct {
	Id               int64  `db:"id"`
	LastEveningIndex int    `db:"last_evening_index"`
	LastMorningIndex int    `db:"last_morning_index"`
	CreatedAt        string `db:"created_at"`
}

type ConfigInclude struct {
	Arabic        bool `db:"arabic"`
	Russian       bool `db:"russian"`
	Transcription bool `db:"transcription"`
}
