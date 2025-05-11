package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	discordToken := os.Getenv("DISCORD_TOKEN")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	defer dg.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if m.Content == "enshrouded_restart" {
		gameServersCmdPath := os.Getenv("GAME_SERVERS_CMD_PATH")

		cmd := exec.Command("powershell", "-WindowStyle", "Hidden", "-Command", gameServersCmdPath+"enshrouded_restart.ps1")
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error executing enshrouded_restart: %s", err.Error())
			s.ChannelMessageSend(m.ChannelID, "Enshrouded restart error !")
			return
		}
		s.ChannelMessageSend(m.ChannelID, "Enshrouded restarted !")
	}

	if m.Content == "enshrouded_info" {
		gameServersCmdPath := os.Getenv("GAME_SERVERS_CMD_PATH")

		output, err := exec.Command("powershell", "-WindowStyle", "Hidden", "-Command", gameServersCmdPath+"enshrouded_info.ps1").Output()
		if err != nil {
			fmt.Printf("Error executing enshrouded_info: %s", err.Error())
			s.ChannelMessageSend(m.ChannelID, "Enshrouded restart error !")
			return
		}
		s.ChannelMessageSend(m.ChannelID, string(output))
	}
}
