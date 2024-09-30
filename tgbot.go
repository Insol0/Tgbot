package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7140252531:AAHmIWCqnkmtgwv_WWnhAhFw8dilIk7sM1U")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите опцию:")
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Получить отзывы"),
				tgbotapi.NewKeyboardButton("Оставить сообщение"),
			),
		)

		bot.Send(msg)
	}
}

// func Messages() {
// 	response, err := http.Get("eyJhbGciOiJFUzI1NiIsImtpZCI6IjIwMjQwNzE1djEiLCJ0eXAiOiJKV1QifQ.eyJlbnQiOjEsImV4cCI6MTczNjk5NTIzMCwiaWQiOiJiYmI4YTQyOS0yYjRiLTQ4ZTUtOTRlNi02YzFmNDBkZGE5ZDAiLCJpaWQiOjUzMjk4ODk0LCJvaWQiOjExNzI0NzQsInMiOjEyOCwic2lkIjoiMzFhMzNhZjMtN2JkNy00MzQ1LWFmYmItOWJlZDgwYTJiMzE1IiwidCI6ZmFsc2UsInVpZCI6NTMyOTg4OTR9.tn9yaSThziQ-AxluPhy-hGZhrlpnINCCQoaf2tKV3I-_FcRobHb2V0dZrBe0sEepHBbxSgSl2fTIEfSLa7oTew")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Print(response)
// 	defer response.Body.Close()
// }
