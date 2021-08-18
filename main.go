package main

import (
	"otpbot/app"
)

func main() {
	go app.ADSync()
	app.StartBot()
}







