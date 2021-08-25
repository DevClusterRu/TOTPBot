package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/xlzd/gotp"
	"log"
	"strconv"
	"strings"
	"time"
)

func Generator_otp(Key string) string {
	totp := gotp.NewDefaultTOTP(Key)
	code, exp:=totp.NowWithExpiration()
	return fmt.Sprintf("%s Осталось %d",code, (exp-time.Now().Unix()))
}

func StartBot() {
	var numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Получить код доступа"),
		),
	)

	bot, err := tgbotapi.NewBotAPI("1938796209:AAG9APOYakqvg0PjygEfmyUKEW-LANZY5gc")
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		id := update.Message.From.ID

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		if update.Message.Text == "Получить код доступа" {

			login := Connect_db(strconv.Itoa(id))
			if strings.TrimSpace(login)==""{
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Такой пользователь не найден")
				msg.ReplyMarkup = numericKeyboard
				bot.Send(msg)
			}
			Key := GetUserToken(login)
			if Key == "" {
				continue
			}
			code := Generator_otp(Key)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, code)
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ReplyMarkup = numericKeyboard
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, `Для получение одноразового кода нажмите кнопку "Получить код доступа"` )
			msg.ReplyMarkup = numericKeyboard
			bot.Send(msg)
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	}

}

