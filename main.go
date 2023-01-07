package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"sdushnik/cmd/bot"

	"log"
)

type Config struct {
	Env      string
	BotToken string
}

func main() {

	configPath := flag.String("config", "", "")
	flag.Parse()

	cfg := &Config{}
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	dimxxBot := bot.DimxxBot{
		bot.InitBot(cfg.BotToken),
	}

	dimxxBot.Bot.Handle("/start", dimxxBot.StartHandler)
	dimxxBot.Bot.Handle("/week", dimxxBot.SemesterHandle)
	dimxxBot.Bot.Handle("/ai", dimxxBot.AskHandler)
	dimxxBot.Bot.Handle("/rand", dimxxBot.RandomChoiceHandler)
	dimxxBot.Bot.Handle("/mtvtn", dimxxBot.RandomMotivationText)
	dimxxBot.Bot.Start()
}
