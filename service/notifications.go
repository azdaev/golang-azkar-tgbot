package service

import (
	"log"
	"strings"

	"github.com/azdaev/azkar-tg-bot/azkar"
	"github.com/azdaev/azkar-tg-bot/repository"
	"github.com/azdaev/azkar-tg-bot/repository/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMorningAzkarToAll(bot *tgbotapi.BotAPI, repo *repository.AzkarRepository) {
	userIDs, err := repo.GetUsersForMorningNotification()
	if err != nil {
		log.Printf("error getting users for morning notification: %v", err)
		return
	}

	log.Printf("Sending morning azkar to %d users", len(userIDs))
	for _, userID := range userIDs {
		sendAzkarToUser(bot, repo, userID, true)
	}
}

func SendEveningAzkarToAll(bot *tgbotapi.BotAPI, repo *repository.AzkarRepository) {
	userIDs, err := repo.GetUsersForEveningNotification()
	if err != nil {
		log.Printf("error getting users for evening notification: %v", err)
		return
	}

	log.Printf("Sending evening azkar to %d users", len(userIDs))
	for _, userID := range userIDs {
		sendAzkarToUser(bot, repo, userID, false)
	}
}

func sendAzkarToUser(bot *tgbotapi.BotAPI, repo *repository.AzkarRepository, userID int64, isMorning bool) {
	config, err := repo.Config(userID)
	if err != nil {
		log.Printf("error getting config for user %d: %v", userID, err)
		return
	}

	// Сбросить индекс на 0
	if isMorning {
		err = repo.SetMorningIndex(userID, 0)
	} else {
		err = repo.SetEveningIndex(userID, 0)
	}
	if err != nil {
		log.Printf("error setting index for user %d: %v", userID, err)
	}

	// Отправить азкар
	msg := tgbotapi.NewMessage(userID, azkar.Wrap(config, 0, isMorning))
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = OnlyNextKeyboard

	_, err = bot.Send(msg)
	if err != nil {
		// Проверить если бот заблокирован
		if strings.Contains(err.Error(), "Forbidden") ||
			strings.Contains(err.Error(), "bot was blocked") ||
			strings.Contains(err.Error(), "user is deactivated") ||
			strings.Contains(err.Error(), "bot can't initiate conversation") {
			err = repo.SetBlocked(userID, true)
			if err != nil {
				log.Printf("error setting blocked status for user %d: %v", userID, err)
			}
			log.Printf("user %d blocked the bot", userID)
		} else {
			log.Printf("error sending azkar to user %d: %v", userID, err)
		}
		return
	}

	// Отправить аудио если включено
	if config.Audio {
		audioPath := "media/"
		audioTitle := ""
		if isMorning {
			audioPath += "morning/0.mp3"
			audioTitle = "Утренний зикр №1"
		} else {
			audioPath += "evening/0.mp3"
			audioTitle = "Вечерний зикр №1"
		}
		audio := tgbotapi.NewAudio(userID, tgbotapi.FilePath(audioPath))
		audio.Title = audioTitle
		_, err = bot.Send(audio)
		if err != nil {
			log.Printf("error sending audio to user %d: %v", userID, err)
		}
	}
}

func ShouldShowNotificationSuggestion(repo *repository.AzkarRepository, userID int64, config *models.UserConfig) bool {
	// Если обе рассылки уже включены, не показывать
	if config.MorningNotification && config.EveningNotification {
		return false
	}

	user, err := repo.User(userID)
	if err != nil {
		log.Printf("error getting user %d: %v", userID, err)
		return false
	}

	// Показывать каждый 10-й раз
	return user.AzkarRequestCount%10 == 0 && user.AzkarRequestCount > 0
}
