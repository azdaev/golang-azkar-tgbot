package repository

import (
	"github.com/azdaev/azkar-tg-bot/repository/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type AzkarRepository struct {
	db *sqlx.DB
}

func NewAzkarRepository(db *sqlx.DB) *AzkarRepository {
	return &AzkarRepository{db: db}
}

func (repo *AzkarRepository) LastMorningIndex(userId int64) (index int, err error) {
	err = repo.db.Get(&index, "SELECT last_morning_index FROM users WHERE id=$1", userId)
	return
}

func (repo *AzkarRepository) LastEveningIndex(userId int64) (index int, err error) {
	err = repo.db.Get(&index, "SELECT last_evening_index FROM users WHERE id=$1", userId)
	return
}

func (repo *AzkarRepository) User(userId int64) (*models.User, error) {
	user := &models.User{}
	err := repo.db.Get(user, "SELECT * FROM users where id = $1", userId)
	return user, err
}

func (repo *AzkarRepository) NewUser(userId int64) error {
	_, err := repo.db.Exec("INSERT INTO users(id) VALUES($1)", userId)
	return err
}

func (repo *AzkarRepository) SetMorningIndex(userId int64, index int) (err error) {
	_, err = repo.db.Exec("UPDATE users SET last_morning_index = $1 WHERE id = $2", index, userId)
	return
}

func (repo *AzkarRepository) SetEveningIndex(userId int64, index int) (err error) {
	_, err = repo.db.Exec("UPDATE users SET last_evening_index = $1 WHERE id = $2", index, userId)
	return
}

func (repo *AzkarRepository) Config(userId int64) (*models.ConfigInclude, error) {
	config := &models.ConfigInclude{}
	err := repo.db.Get(config, "SELECT arabic, russian, transcription, audio FROM configs WHERE user_id=$1", userId)
	return config, err
}

func (repo *AzkarRepository) InsertConfig(userId int64) error {
	_, err := repo.db.Exec("INSERT INTO configs(user_id) VALUES($1)", userId)
	return err
}

func (repo *AzkarRepository) UpdateConfig(userId int64, key string, value string) error {
	_, err := repo.db.Exec("UPDATE configs SET "+key+" = $1 WHERE user_id = $2", value, userId)
	return err
}
