package main

import (
	"chatops/cmds"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var router *cmds.CommandsMap

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	discordToken := os.Getenv("DISCORD_TOKEN")

	discordSession, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Fatal("Error creating Discord session,", err)
	}

	// listent to GuildMessages only, see https://discord.com/developers/docs/events/gateway#gateway-intents
	discordSession.Identify.Intents = discordgo.IntentsGuildMessages

	router = cmds.NewRouterCommands()
	discordSession.AddHandler(messageCreateHandler)

	err = discordSession.Open()
	if err != nil {
		log.Fatal("Error opening connection,", err)
	}
	defer discordSession.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	router.Handle(s, m)
}
