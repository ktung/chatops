package cmds

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/bwmarrin/discordgo"
)

func PalworldInfoCommand(s *discordgo.Session, m *discordgo.MessageCreate) error {
	gameServersCmdPath := os.Getenv("GAME_SERVERS_CMD_PATH")

	output, err := exec.Command("powershell", "-WindowStyle", "Hidden", "-Command", gameServersCmdPath+"palserver_info.ps1").Output()
	if err != nil {
		fmt.Printf("Error executing palserver_info: %s", err.Error())
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

	cmd := exec.Command("powershell", "-WindowStyle", "Hidden", "-Command", gameServersCmdPath+"palserver_restart.ps1")

	err := cmd.Run()
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
