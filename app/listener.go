package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/xlzd/gotp"
	"strconv"
	"strings"
	"time"
)

func Generator_otp(Key string) (string, int64) {
	totp := gotp.NewDefaultTOTP(Key)
	code, exp:=totp.NowWithExpiration()
	return code, exp
	//"ФИО: %s\nЛогин: %s\nПароль: %s\nдо окончания действия данного пароля осталось %d сек."
	//return fmt.Sprintf("пароль: %s \nдо окончания действия данного пароля осталось %d секунд",code, (exp-time.Now().Unix()))
}

func (c *Config) ListenTarantool() {

	var numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(c.Button),
		),
	)

	for {

		resp, err := c.Tarantool.Conn.Call("take", []interface{}{"telegramMsgs"})
		if err != nil {
			fmt.Println("TAKE error: ", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if len(resp.Data) == 0 {
			time.Sleep(5 * time.Second)
			continue
		}



		id := resp.Data[0].([]interface{})[0].(uint64)
		uid := resp.Data[0].([]interface{})[2].(uint64)
		chatId := resp.Data[0].([]interface{})[3].(uint64)

		defer func() {
			_, err := c.Tarantool.Conn.Call("ack", []interface{}{"telegramMsgs", id})
			if err != nil {
				fmt.Println("ACK error: ", err)
				return
			}
		}()

		login := Connect_db(strconv.Itoa(int(uid)))
		if strings.TrimSpace(login) == "" {
			msg := tgbotapi.NewMessage(int64(chatId), c.CantFind)
			msg.ReplyMarkup = numericKeyboard
			c.Bot.Send(msg)
		}

		Key := GetUserToken(login)
		if Key == "" {
			continue
		}
		code, exp := Generator_otp(Key)
		//"ФИО: %s\nЛогин: %s\nПароль: %s\nдо окончания действия данного пароля осталось %d сек."
		message:=fmt.Sprintf(c.Answer, "FIO", "LOGIN", code, (exp-time.Now().Unix()))

		msg := tgbotapi.NewMessage(int64(chatId), message)
		msg.ReplyMarkup = numericKeyboard
		c.Bot.Send(msg)
	}

}