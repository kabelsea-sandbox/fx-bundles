package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewBotAPI(config *Config) (*tgbotapi.BotAPI, error) {
	var (
		bot *tgbotapi.BotAPI
		err error
	)

	if bot, err = tgbotapi.NewBotAPI(config.Telegram.Bot.Token); err != nil {
		return nil, err
	}

	if config.Debug {
		bot.Debug = true
	}

	return bot, nil
}
