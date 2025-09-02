package telegram

import (
	"currency-converter-bot/api"
	"currency-converter-bot/storage"
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func list(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	table, err := storage.List(update.Message.From.ID)
	if err != nil {
		return err
	}

	if len(table) == 0 {
		replay(bot, "📭 У вас пока нет сохранённых конвертаций.  \nПопробуйте выполнить команду `/convert`, и результат появится здесь!\n", update.Message.Chat.ID)
		return fmt.Errorf("no data available")
	}

	for _, row := range table {
		replay(bot, fmt.Sprintf("📅 Дата: %s\n💱 Из: %s → %s\n💵 Сумма: %.2f\n📊 Результат: %.2f", row.Date, row.SourceCurrency, row.TargetCurrency, row.Amount, row.Result), update.Message.Chat.ID)
	}

	return nil
}

func ExchangeRate(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {

	args := update.Message.CommandArguments()
	if args == "" {
		replay(bot, "⚠️ Пожалуйста, укажите валюты в формате:  \n/exchangerate <FROM> <TO>  \n\nПример: `/exchangerate USD RUB`\n", update.Message.Chat.ID)
	}

	values := strings.Split(args, " ")

	source := strings.ToUpper(values[0])
	target := strings.ToUpper(values[1])

	var conversationRates map[string]float64
	conversationRates, err := api.GetInfo(source)
	if err != nil {
		return err
	}

	replay(bot, fmt.Sprintf("Актуальный курс %s к %s: %.2f", source, target, conversationRates[target]), update.Message.Chat.ID)

	return nil
}

func ConvertCurrency(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {

	args := update.Message.CommandArguments()
	if args == "" {
		replay(bot, "⚠️ Пожалуйста, используйте формат:  \n/convert <FROM> <TO> <AMOUNT>  \n\nПример: `/convert EUR USD 50`\n", update.Message.Chat.ID)
	}

	values := strings.Split(args, " ")
	source := strings.ToUpper(values[0])
	target := strings.ToUpper(values[1])
	amount, err := strconv.ParseFloat(values[2], 64)
	if err != nil {
		return err
	}

	var conversationRates map[string]float64
	conversationRates, err = api.GetInfo(source)
	if err != nil {
		return err
	}

	conversationRate := conversationRates[target]
	result := amount * conversationRate
	replay(bot, fmt.Sprintf("Актуальный курс %s к %s: %.2f\nСумма: %.2f", source, target, conversationRate, result), update.Message.Chat.ID)
	err = storage.Insert(update.Message.From.ID, source, target, amount, result)
	if err != nil {
		return err
	}
	return nil
}
