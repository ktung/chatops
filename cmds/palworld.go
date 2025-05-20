package cmds

import (
	"chatops/internal"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func PalworldInfoCommand(s *discordgo.Session, m *discordgo.MessageCreate) error {
	gameServersCmdPath := os.Getenv("GAME_SERVERS_CMD_PATH")
	cmd := gameServersCmdPath + "palserver_info.ps1"

	output, err := internal.ExecutePowerShellCommand(cmd)
	if err != nil {
		fmt.Printf("Error executing %s: %s", cmd, err.Error())
		_, err := s.ChannelMessageSend(m.ChannelID, "Palserver info error !")
		if err != nil {
			fmt.Printf("Error sending message: %s", err.Error())
		}
		return err
	}

	s.ChannelMessageSend(m.ChannelID, string(output))
	return nil
}

func PalworldRestartCommand(s *discordgo.Session, m *discordgo.MessageCreate) error {
	gameServersCmdPath := os.Getenv("GAME_SERVERS_CMD_PATH")
	cmd := gameServersCmdPath + "palserver_restart.ps1"

	err := internal.RunPowerShellCommand(cmd)
	if err != nil {
		fmt.Printf("Error executing palserver_restart: %s", err.Error())
		_, err := s.ChannelMessageSend(m.ChannelID, "Palserver restart error !")
		if err != nil {
			fmt.Printf("Error sending message: %s", err.Error())
		}
		return err
	}

	s.ChannelMessageSend(m.ChannelID, "Palserver restarted !")
	return nil
}
