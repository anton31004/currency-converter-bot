package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot, nil
}

func StartBot(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		} else if update.Message.IsCommand() {
			err := handleMessage(update, bot)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func handleMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	chatID := update.Message.Chat.ID
	switch update.Message.Command() {
	case "start":
		replay(bot, "👋 Добро пожаловать!  \nЯ — ваш помощник для конвертации валют и просмотра актуальных курсов.  \n\n📌 Доступные команды:  \n/start — начать работу  \n/help — справка  \n/list — история конвертаций  \n/exchangerate — узнать курс валют  \n/convert — конвертировать сумму\n", chatID)
	case "help":
		replay(bot, "ℹ️ Справка по командам:\n\n/exchangerate <FROM> <TO>  \n→ Покажу актуальный курс.  \nПример: `/exchangerate USD EUR`\n\n/convert <FROM> <TO> <AMOUNT>  \n→ Переведу указанную сумму.  \nПример: `/convert USD RUB 100`\n\n/list  \n→ Отображу историю ваших конвертаций.\n", chatID)
	case "list":
		err := list(bot, update)
		if err != nil {
			return err
		}
	case "ExchangeRate":
		err := ExchangeRate(bot, update)
		if err != nil {
			return err
		}
	case "convert":
		err := ConvertCurrency(bot, update)
		if err != nil {
			return err
		}
	default:
		replay(bot, "🤔 Я не знаю такую команду.  \nНапишите `/help`, чтобы посмотреть доступные варианты.\n", update.Message.Chat.ID)
	}
	return nil
}

func replay(bot *tgbotapi.BotAPI, text string, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)
}
