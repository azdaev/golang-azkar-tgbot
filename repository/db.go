package repository

import (
	"fmt"
	"github.com/azdaev/azkar-tg-bot/repository/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
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

func (repo *AzkarRepository) User(userId int64) (user *models.User, err error) {
	err = repo.db.Get(&user, "SELECT * FROM users where id = $1", userId)
	return
}

func (repo *AzkarRepository) NewUser(userId int64) (err error) {
	_, err = repo.db.Exec("INSERT INTO users(id) VALUES($1)", userId)
	return
}

func (repo *AzkarRepository) SetMorningIndex(userId int64, index int) (err error) {
	_, err = repo.db.Exec("UPDATE users SET last_morning_index = $1 WHERE id = $2", index, userId)
	return
}

func (repo *AzkarRepository) SetEveningIndex(userId int64, index int) (err error) {
	_, err = repo.db.Exec("UPDATE users SET last_evening_index = $1 WHERE id = $2", index, userId)
	return
}

func (repo *AzkarRepository) Config(userId int64) (config *models.ConfigInclude, err error) {
	config = &models.ConfigInclude{}
	err = repo.db.Get(config, "SELECT arabic, russian, transcription FROM configs WHERE user_id = $1", userId)
	return
}

func (repo *AzkarRepository) InsertConfig(userId int64) (err error) {
	_, err = repo.db.Exec("INSERT INTO configs(user_id) VALUES ($1)", userId)
	return err
}

func (repo *AzkarRepository) UpdateConfig(userId int64, key string, value string) (err error) {
	log.Println(key, value)
	query := fmt.Sprintf("UPDATE configs SET %s = '%s' WHERE user_id = %d", key, value, userId)
	log.Println(query)
	_, err = repo.db.Exec(query)
	return
}
