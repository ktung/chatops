package cmds

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func NewHelpCommand(router *CommandsMap) CommandHandler {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) error {
		helpMessage := "Available commands:\n"
		for _, command := range router.GetAllCommands() {
			helpMessage += "- " + command + "\n"
		}

		_, err := s.ChannelMessageSend(m.ChannelID, helpMessage)
		if err != nil {
			log.Printf("Error sending Help message: %v", err)
			return err
		}

		return nil
	}
}

func WorkCommand(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "Back to work ! See you in 1 hour !")
	if err != nil {
		log.Printf("Error sending work message: %v", err)
		return err
	}

	go func() {
		time.Sleep(1*time.Hour)

		_, err := s.ChannelMessageSend(m.ChannelID, "Work done !")
		if err != nil {
			log.Printf("Error sending work message: %v", err)
		}
	}()

	return nil
}

func BreakCommand(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "Let's take a 15 minutes break")
	if err != nil {
		log.Printf("Error sending break message: %v", err)
		return err
	}

	go func() {
		time.Sleep(15*time.Minute)

		_, err := s.ChannelMessageSend(m.ChannelID, "Break done !")
		if err != nil {
			log.Printf("Error sending break message: %v", err)
		}
	}()

	return nil
}
