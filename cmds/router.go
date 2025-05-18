package cmds

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler func(s *discordgo.Session, m *discordgo.MessageCreate) error

type CommandsMap struct {
	handlers map[string]CommandHandler
}

func NewRouterCommands() *CommandsMap {
	return &CommandsMap{
		handlers: map[string]CommandHandler{
			"help":               HelpCommand,
			"ping":               PingCommand,
			"pong":               PongCommand,
			"enshrouded_info":    EnshroudedInfoCommand,
			"enshrouded_restart": EnshroudedRestartCommand,
			"palserver_info":     PalworldInfoCommand,
			"palserver_restart":  PalworldRestartCommand,
		},
	}
}

func (c *CommandsMap) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if handler, ok := c.handlers[m.Content]; ok {
		if err := handler(s, m); err != nil {
			log.Printf("Error handling command: %v", err)
		}
	}
}
