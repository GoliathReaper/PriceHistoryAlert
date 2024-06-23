package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func telegramNotifications(msgTxt string) {
	const TELEGRAM_APITOKEN = "add your telegram API Token"
	const chatID = "Add your chat ID"
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_APITOKEN)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	// Now that we know we've gotten a new message, we can construct a
	// reply! We'll take the Chat ID and Text from the incoming message
	// and use it to create a new message.
	msg := tgbotapi.NewMessage(chatID, msgTxt)

	// Okay, we're sending our message off! We don't care about the message
	// we just sent, so we'll discard it.
	if _, err := bot.Send(msg); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		panic(err)
	}
}
