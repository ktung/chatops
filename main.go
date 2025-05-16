package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"chatops/internal/bot"
	"chatops/internal/gameserver"
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
	if m.Author.ID == s.State.User.ID {
		return
	}

	router := bot.NewCommandRouter()
	serverManager := gameserver.NewServerManager(os.Getenv("GAME_SERVERS_CMD_PATH"))

	// Register basic commands
	router.Register("ping", func(s *discordgo.Session, m *discordgo.MessageCreate) error {
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
		return err
	})

	router.Register("pong", func(s *discordgo.Session, m *discordgo.MessageCreate) error {
		_, err := s.ChannelMessageSend(m.ChannelID, "Ping!")
		return err
	})

	// Register game server commands
	router.Register("enshrouded_restart", createServerHandler(s, m, serverManager, "enshrouded", "restart"))
	router.Register("enshrouded_info", createServerHandler(s, m, serverManager, "enshrouded", "info"))
	router.Register("palserver_restart", createServerHandler(s, m, serverManager, "palserver", "restart"))
	router.Register("palserver_info", createServerHandler(s, m, serverManager, "palserver", "info"))

	router.Handle(s, m)
}

func createServerHandler(s *discordgo.Session, m *discordgo.MessageCreate, manager *gameserver.ServerManager, serverType, action string) bot.CommandHandler {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) error {
		output, err := manager.ExecuteCommand(serverType, action)
		if err != nil {
			log.Printf("Error executing %s_%s: %s", serverType, action, err)
			_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s %s error!", serverType, action))
			return err
		}
		_, err = s.ChannelMessageSend(m.ChannelID, output)
		return err
	}
}
