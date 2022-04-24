package handlers

import (
	"log"

	. "github.com/bwmarrin/discordgo"
)

// Prints the text if the bot is ready
func OnReady(s *Session, r *Ready) {
	log.Printf("%v#%v is ready!",
		s.State.User.Username,
		s.State.User.Discriminator)
}
