package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"path"
)

type Config struct {
	TARANTOOL string `yaml:"TARANTOOL"`
	TELEGRAM  string `yaml:"TELEGRAM"`
	Button string `yaml:"button"`
	Greet string `yaml:"greet"`
	Answer string `yaml:"answer"`
	CantFind string `yaml:"cant_find"`
	Tarantool *TTool
	Metrics   *Metrics
	Bot *tgbotapi.BotAPI
}

func NewConfig(fileName string) (config *Config, err error) {
	log.Printf("reading config from '%s'", fileName)
	if ext := path.Ext(fileName); ext != ".yaml" && ext != ".yml" {
		err = fmt.Errorf("invalid file '%s' extenstion, expected 'yaml' or 'yml'", ext)
		return
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		err = fmt.Errorf("can't read file '%s'", fileName)
		return
	}

	config = new(Config)
	if err = yaml.Unmarshal(file, config); err != nil {
		err = fmt.Errorf("file %s yaml unmarshal error: %v", fileName, err)
	}

	config.Metrics = &Metrics{}

	config.Bot, err = tgbotapi.NewBotAPI(config.TELEGRAM)
	if err != nil {
		log.Panic(err)
	}

	config.Tarantool = NewTarantool(config.TARANTOOL)
	return config, err
}
