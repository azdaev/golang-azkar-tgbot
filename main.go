package main

import (
	"log"
	"os"
	"strings"

	"github.com/azdaev/azkar-tg-bot/azkar"
	"github.com/azdaev/azkar-tg-bot/repository"
	"github.com/azdaev/azkar-tg-bot/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sqlx.Connect("sqlite3", "repository/azkar")
	if err != nil {
		log.Fatalln(err)
	}
	azkarRepository := repository.NewAzkarRepository(db)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	menu := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     "/morning",
			Description: "Утренние азкары",
		},
		tgbotapi.BotCommand{
			Command:     "/evening",
			Description: "Вечерние азкары",
		},
		tgbotapi.BotCommand{
			Command:     "/settings",
			Description: "Настройки вывода",
		},
	)

	_, _ = bot.Request(menu)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			m := update.Message

			err = service.EnsureUser(azkarRepository, m.From.ID) // save user to db if not exists
			if err != nil {
				log.Printf("error ensure user: %s\n", err)
			}

			config, err := azkarRepository.Config(m.From.ID) // config: what to print out
			if err != nil {
				log.Printf("error get config: %s\n", err)
			}
			if config == nil {
				err := azkarRepository.InsertConfig(m.From.ID)
				if err != nil {
					log.Printf("error create config: %s\n", err)
				}

				config, _ = azkarRepository.Config(m.From.ID)
			}

			var response tgbotapi.MessageConfig

			switch command := m.Command(); command {
			case "start":
				messageText := "السلام عليكم ورحمة الله وبركاته \n\n"
				messageText += "Выберите время для азкаров:"
				response := tgbotapi.NewMessage(m.Chat.ID, messageText)
				response.ReplyMarkup = service.MorningEveningKeyboard
				bot.Send(response)
				continue

			case "morning": // TODO: export to another function
				response = tgbotapi.NewMessage(m.Chat.ID, azkar.Wrap(config, 0, true))
				response.ParseMode = "HTML"
				err := azkarRepository.SetMorningIndex(m.From.ID, 0)
				if err != nil {
					log.Printf("error set morning index: %s\n", err)
				}

				response.ReplyMarkup = service.OnlyNextKeyboard
				bot.Send(response)

				if !config.Audio {
					continue
				}

				audio := tgbotapi.NewAudio(m.Chat.ID, tgbotapi.FilePath("media/morning/0.mp3"))
				audio.Title = "Утренний зикр №1"
				bot.Send(audio)

			case "evening": // TODO: export to another function
				response = tgbotapi.NewMessage(m.Chat.ID, azkar.Wrap(config, 0, false))
				response.ParseMode = "HTML"
				err := azkarRepository.SetEveningIndex(m.From.ID, 0)
				if err != nil {
					log.Printf("error set evening index: %s\n", err)
				}

				response.ReplyMarkup = service.OnlyNextKeyboard
				bot.Send(response)

				if !config.Audio {
					continue
				}

				audio := tgbotapi.NewAudio(m.Chat.ID, tgbotapi.FilePath("media/evening/0.mp3"))
				audio.Title = "Вечерний зикр №1"
				bot.Send(audio)

			case "settings":
				response = tgbotapi.NewMessage(m.Chat.ID, "Выберите что требуется выводить")
				response.ReplyMarkup = service.ConfigKeyboard(config)
				bot.Send(response)
			}

		} else if update.CallbackQuery != nil {
			switch {
			case update.CallbackQuery.Data == "show_morning" || update.CallbackQuery.Data == "show_evening":
				service.HandleMorningEvening(bot, update.CallbackQuery, azkarRepository)
			case update.CallbackQuery.Data == "previous" || update.CallbackQuery.Data == "next":
				service.HandleDirection(bot, update.CallbackQuery, azkarRepository)
			case strings.HasPrefix(update.CallbackQuery.Data, "set"):
				err := service.HandleConfigEdit(bot, update.CallbackQuery, azkarRepository)
				if err != nil {
					log.Printf("error handle config edit: %s\n", err)
				}
			}
		}
	}
}
