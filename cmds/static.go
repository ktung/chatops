package cmds

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func PingCommand(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
	if err != nil {
		log.Printf("Error sending Pong message: %v", err)
	}
	return err
}

func PongCommand(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "Ping!")
	if err != nil {
		log.Printf("Error sending Ping message: %v", err)
	}
	return err
}

func NewHelpCommand(router *CommandsMap) CommandHandler {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) error {
		helpMessage := "Available commands:\n"
		for _, command := range router.GetAllCommands() {
				helpMessage += "- " + command + "\n"
		}

		_, err := s.ChannelMessageSend(m.ChannelID, helpMessage)
		if err != nil {
				log.Printf("Error sending Help message: %v", err)
		}
		return err
	}
}
