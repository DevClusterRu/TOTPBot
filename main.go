package main

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/xlzd/gotp"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	out := bytes.Buffer{}
	cmd := exec.Command("/usr/local/bin/multiotp/multiotp", " -urllink", "alexandrov.v")
	cmd.Stdout = &out
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		log.Println("==>", err)
	}
	fmt.Println(cmd.String())

	//bot_message()
}

func Generator_otp(Key string) string {
	totp := gotp.NewDefaultTOTP(Key)
	return totp.Now()

}


func bot_message() {
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
		if update.Message.Text == "give" {

			Key := Connect_db(strconv.Itoa(id))
			if Key==""{
				continue
			}

            code := Generator_otp(Key)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, code)
			//msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Для получение одноразового кода введите слово: give")
			bot.Send(msg)
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	}

}

func Connect_db(telegramId string) string {

	db, err := sql.Open("mysql", ":@tcp(192.168.0.100:3306)/totp")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	rows, err := db.Query("select security_key from users where telegram = ?", telegramId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var Key string

	for rows.Next() {

		err := rows.Scan(&Key)
		if err != nil {
			fmt.Println(err)
			continue
		}

	}
	return Key
}
