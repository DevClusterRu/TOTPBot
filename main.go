package main

import (
	"fmt"
	"otpbot/app"
)

func main() {

login := app.Connect_db("123")
fmt.Println(login)
	//app.StartBot()
}







