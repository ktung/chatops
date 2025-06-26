package cmds

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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
	minutes := 60 // Default to 60 minutes
	messageFields := strings.Fields(m.Content)
	if len(messageFields) == 0 {
		return nil
	}

	args := messageFields[1:]
	if len(args) > 0 {
		if n, err := strconv.Atoi(args[0]); err == nil && n > 0 {
			minutes = n
		}
	}

	_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Back to work ! See you in %d minutes !", minutes))
	if err != nil {
		log.Printf("Error sending work message: %v", err)
		return err
	}

	go func() {
		time.Sleep(time.Duration(minutes) * time.Minute)

		_, err := s.ChannelMessageSend(m.ChannelID, "Work done !")
		if err != nil {
			log.Printf("Error sending work message: %v", err)
		}
	}()

	return nil
}

func BreakCommand(s *discordgo.Session, m *discordgo.MessageCreate) error {
	minutes := 15 // Default to 15 minutes
	messageFields := strings.Fields(m.Content)
	if len(messageFields) == 0 {
		return nil
	}

	args := messageFields[1:]
	if len(args) > 0 {
		if n, err := strconv.Atoi(args[0]); err == nil && n > 0 {
			minutes = n
		}
	}

	_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Let's take a %d minutes break", minutes))
	if err != nil {
		log.Printf("Error sending break message: %v", err)
		return err
	}

	go func() {
		time.Sleep(time.Duration(minutes)*time.Minute)

		_, err := s.ChannelMessageSend(m.ChannelID, "Break done !")
		if err != nil {
			log.Printf("Error sending break message: %v", err)
		}
	}()

	return nil
}
