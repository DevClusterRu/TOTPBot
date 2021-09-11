package main

import (
	"github.com/kardianos/osext"
	"log"
	"otpbot/app"
	"path"
	"time"
)

func main() {

	cfgfile:="local.yml"
	cfg, err := app.NewConfig(cfgfile)
	if err != nil {
		log.Println("Something wrong with config: ", err)
		folderPath, err := osext.ExecutableFolder()
		if err != nil {
			log.Fatal("Cant find env folder path", err)
		}
		cfg, err = app.NewConfig(folderPath+"/"+path.Base(cfgfile))
		if err!=nil{
			log.Fatal("Config create error", err)
		}
	}

	go func() {
		for {
			if !cfg.Tarantool.Conn.ConnectedNow() {
				cfg.Tarantool.Reconnect()
			}
			time.Sleep(5*time.Second)
		}
	}()

	go app.ADSync()
	go cfg.ListenTarantool()

	stop:=make(chan struct{})
	<-stop

}







