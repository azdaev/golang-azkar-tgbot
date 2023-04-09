package service

import (
	"fmt"
	"github.com/azdaev/azkar-tg-bot/azkar"
	"github.com/azdaev/azkar-tg-bot/repository"
	"github.com/azdaev/azkar-tg-bot/repository/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

var OnlyNextKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("➡️", "next"),
	),
)

var OnlyPreviousKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬅️", "previous"),
	),
)

var BothSidesKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬅️", "previous"),
		tgbotapi.NewInlineKeyboardButtonData("➡️", "next"),
	),
)

func EnsureUser(repo *repository.AzkarRepository, id int64) (err error) {
	user, err := repo.User(id)
	if err != nil {
		return
	}

	if user == nil {
		err = repo.NewUser(id)
		if err != nil {
			return
		}

		err = repo.InsertConfig(id)
		if err != nil {
			return
		}
	}

	return
}

//func PrettyConfig(config *models.ConfigInclude) string {
//	result := "Выберите, что требуется выводить:\n\n"
//
//	result += "Оригинал - "
//	if !config.Arabic {
//		result += "❌\n"
//	} else {
//		result += "✅\n"
//	}
//
//	result += "Перевод - "
//	if !config.Russian {
//		result += "❌\n"
//	} else {
//		result += "✅\n"
//	}
//
//	result += "Транскрипция - "
//	if !config.Transcription {
//		result += "❌\n"
//	} else {
//		result += "✅\n"
//	}
//
//	return result
//}

func ConfigKeyboard(config *models.ConfigInclude) *tgbotapi.InlineKeyboardMarkup {
	m := map[bool]string{
		true:  "✅",
		false: "❌",
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Оригинал "+m[config.Arabic], fmt.Sprintf("set %s %v", "arabic", !config.Arabic))),
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Перевод "+m[config.Russian], fmt.Sprintf("set %s %v", "russian", !config.Russian))),
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Транскрипция "+m[config.Transcription], fmt.Sprintf("set %s %v", "transcription", !config.Transcription))),
	)

	return &keyboard
}

func SetConfigKeyboard(message *tgbotapi.EditMessageTextConfig, config *models.ConfigInclude) {
	message.ReplyMarkup = ConfigKeyboard(config)
}

func SetDirectionKeyboard(message *tgbotapi.EditMessageTextConfig, index, length int) {
	if index == 0 {
		message.ReplyMarkup = &OnlyNextKeyboard
	} else if index < length-1 {
		message.ReplyMarkup = &BothSidesKeyboard
	} else {
		message.ReplyMarkup = &OnlyPreviousKeyboard
	}
}

func HandleDirection(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery, azkarRepository *repository.AzkarRepository) {
	sourceMessage := callbackQuery.Message
	userId := callbackQuery.From.ID
	isMorning := []rune(sourceMessage.Text)[0] == 'У'
	currentAzkarSlice := azkar.CurrentAzkarSlice(isMorning)
	currentZikrIndex := 0
	newZikrIndex := 0
	config, err := azkarRepository.Config(userId)
	if err != nil {
		log.Println(err)
	}

	if isMorning {
		currentZikrIndex, err = azkarRepository.LastMorningIndex(userId)
	} else {
		currentZikrIndex, err = azkarRepository.LastEveningIndex(userId)
	}
	if err != nil {
		return
	}

	var editedMessage tgbotapi.EditMessageTextConfig

	if callbackQuery.Data == "previous" {
		newZikrIndex = currentZikrIndex - 1
		bot.Request(tgbotapi.NewCallback(callbackQuery.ID, "Предыдущий"))
	} else if callbackQuery.Data == "next" {
		newZikrIndex = currentZikrIndex + 1
		bot.Request(tgbotapi.NewCallback(callbackQuery.ID, "Следующий"))
	}

	editedMessage = tgbotapi.NewEditMessageText(
		sourceMessage.Chat.ID,
		sourceMessage.MessageID,
		azkar.Wrap(config, newZikrIndex, isMorning),
	)

	if isMorning {
		err = azkarRepository.SetMorningIndex(userId, newZikrIndex)
	} else {
		err = azkarRepository.SetEveningIndex(userId, newZikrIndex)
	}
	if err != nil {
		log.Println(err)
	}

	SetDirectionKeyboard(&editedMessage, newZikrIndex, len(currentAzkarSlice))

	bot.Send(editedMessage)

	audioFilePath := "media/"
	audioTitle := ""

	if isMorning {
		audioFilePath += "morning/"
		audioTitle += "Утренний зикр №"
	} else {
		audioFilePath += "evening/"
		audioTitle += "Вечерний зикр №"
	}

	audioFilePath += strconv.Itoa(newZikrIndex)
	audioFilePath += ".mp3"
	audio := tgbotapi.NewAudio(sourceMessage.Chat.ID, tgbotapi.FilePath(audioFilePath))

	audio.Title = audioTitle + strconv.Itoa(newZikrIndex+1)
	//audio.Caption = azkar.Wrap(config, newZikrIndex, isMorning)
	bot.Send(audio)
}

func HandleConfigEdit(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery, azkarRepository *repository.AzkarRepository) (err error) {
	splitCallbackData := strings.Split(callbackQuery.Data, " ")
	if splitCallbackData[0] != "set" {
		return fmt.Errorf("invalid callback data")
	}

	key, value := splitCallbackData[1], splitCallbackData[2]

	sourceMessage := callbackQuery.Message
	userId := callbackQuery.From.ID

	err = azkarRepository.UpdateConfig(userId, key, value)
	if err != nil {
		return
	}

	config, err := azkarRepository.Config(userId)
	if err != nil {
		return
	}

	editedMessage := tgbotapi.NewEditMessageText(
		sourceMessage.Chat.ID,
		sourceMessage.MessageID,
		sourceMessage.Text,
	)

	SetConfigKeyboard(&editedMessage, config)
	bot.Request(tgbotapi.NewCallback(callbackQuery.ID, "Принято"))
	bot.Send(editedMessage)
	return
}
