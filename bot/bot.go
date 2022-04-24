// Copyright (c) 2022, Daniel M. Jaén
// All rights reserved.

package bot

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

// A global session
var Session *discordgo.Session

// Set the global session.
// Pass a nil argument to use the environment token
// 	// token as argument
// 	token = "X.Y.Z"
// 	bot.Create(&token)
//
// 	// environment token
// 	bot.Create(nil)
func Create(token *string) {
	var (
		err        error
		readyToken string
	)

	if token != nil {
		readyToken = *token
	} else {
		readyToken = os.Getenv("PROJECTOR_TOKEN")
	}

	Session, err = discordgo.New(fmt.Sprintf("Bot %s", readyToken))

	if err != nil {
		panic(err)
	} else {
		log.Println("Session created")
	}
}
