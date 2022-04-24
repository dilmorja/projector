package main

import (
	"os"
	"os/signal"
	"log"

	"github.com/dilmorja/projector/bot"
	"github.com/dilmorja/projector/bot/handlers"
)

func main() {
	bot.Create(nil)

	bot.Session.AddHandler(handlers.OnReady)

	err := bot.Session.Open()

	if err != nil {
		panic(err)
	}

	defer bot.Session.Close()

	alto := make(chan os.Signal, 1)
	signal.Notify(alto, os.Interrupt)
	log.Println("Running...")
	<-alto

}
