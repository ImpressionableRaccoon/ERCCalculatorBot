package main

import (
	"errors"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ImpressionableRaccoon/ERCCalculatorBot/configs"
	"github.com/ImpressionableRaccoon/ERCCalculatorBot/decoder"
)

const (
	helloMessage = `
Привет!
Этот бот генерирует ключ разблокировки магнитол Toyota
Для работы просто оправьте сюда ERC`
	lengthErrorMessage = `
Неправильная длина ERC!
Должно быть 16 символов`
	formatErrorMessage = `
Неправильный формат ERC!
Допустимые символы: 0123456789ABCDEF`
)

var ErrSendingMessage = errors.New("error while sending message")

func main() {
	cfg := configs.LoadEnvVariables()

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Text == "/start" {
				sendMessage(bot, update.Message.Chat.ID, 0, helloMessage)
				continue
			}

			key, err := decoder.Decode(update.Message.Text)
			if errors.Is(err, decoder.ErrInvalidNumberOfCharacters) {
				sendMessage(bot, update.Message.Chat.ID, update.Message.MessageID, lengthErrorMessage)
				continue
			}
			if errors.Is(err, decoder.ErrWrongERCFormat) {
				sendMessage(bot, update.Message.Chat.ID, update.Message.MessageID, formatErrorMessage)
				continue
			}
			if err != nil {
				sendMessage(bot, update.Message.Chat.ID, update.Message.MessageID, err.Error())
				continue
			}

			sendMessage(bot, update.Message.Chat.ID, update.Message.MessageID, key)
		}
	}
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, messageID int, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ReplyToMessageID = messageID
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("%s, %s", ErrSendingMessage, err)
	}
}
